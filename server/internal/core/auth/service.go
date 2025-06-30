package auth

import (
	"context"
	"slices"
	"time"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/token"
	"github.com/server/pkg/utils"

	"github.com/server/internal/adapters/bootstrap/schema"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/internal/core/auth/session"
	"github.com/server/internal/core/auth/strategy"
	authToken "github.com/server/internal/core/auth/token"
)

/*
 High level overview of AuthService should not be directly imported.
 Copy interface and use dependancy injection over direct import.
*/
type IAuthService interface {
	SupportedStrategies() []string
	SanitizeReturnToURL(url string) string
	SanitizeCallbackURL(url string) string
	StartFlow(ctx context.Context, input *domain.StartFlowReq) (*domain.StartFlowResp, error)
	FinishFlow(ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error)
	IsEmailTaken(ctx context.Context, email string) (*bool, error) 
	GetPrincipal(ctx context.Context, assertions ...domain.ClientAssertion) (domain.Principal, error)
	PublicKeySet(ctx context.Context) token.JWkSet 
	BuildToken(ctx context.Context, principal domain.Principal, metadata map[string]string) ([]byte, error) 
}

type AuthService struct {
	cfg						Config
	authStore 				IAuthStore
	emailService 			IEmailService
	flowStore      			IFlowStore
	transactService			ITXAuthService
	tokenService			ITokenService
	sessionService			ISessionService
	serviceUserService 		IServiceUserService
	Now 					func() time.Time
}

func NewAuthService(
	cfg					Config,
	auth				IAuthStore,
	flow    			IFlowStore,
	transact     		ITXAuthService,
	smtp 				IEmailService,
	token 				ITokenService,
	serviceUser		 	IServiceUserService,
	session				ISessionService,
) AuthService {
	return AuthService{
		cfg:					cfg,
		authStore:  			auth, 		
		flowStore:      		flow,
		transactService: 		transact,		
		emailService:    		smtp,
		tokenService: 			token,	
		sessionService:			session,	
		serviceUserService: 	serviceUser,
		Now: func () time.Time {
			return time.Now().UTC()
		},
	}
}

// SanitizeReturnToURL allows only redirect to white listed domains from config
// to avoid https://cheatsheetseries.owasp.org/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.html
func (s AuthService) SanitizeReturnToURL(url string) string {
	if len(url) == 0 {
		return ""
	}

	if len(s.cfg.AuthorizedRedirectURLs) == 0 {
		return ""
	}

	if slices.Contains(s.cfg.AuthorizedRedirectURLs, url) {
		return url
	}

	return ""
}

// SanitizeCallbackURL allows only callback host to white listed domains from config
func (s AuthService) SanitizeCallbackURL(url string) string {
	if len(s.cfg.CallbackURLs) == 0 {
		return ""
	}

	if len(url) == 0 {
		return s.cfg.CallbackURLs[0]
	}

	if slices.Contains(s.cfg.CallbackURLs, url) {
		return url
	}

	return ""
}

func (s AuthService) findTemplate(reason domain.Reason, strat domain.Strategy) (strategy.MailTemplateConfig, error) {
	switch strat {
		case strategy.LinkStrategy: 
			if !s.cfg.Link.MailTemplates[reason].Enabled {
				return strategy.DefaultMajicLinkTemplate(reason.String(), s.cfg.OTP.MailTemplates[reason]), nil
			} else {
				return s.cfg.Link.MailTemplates[reason], nil
			} 
		case strategy.OTPStrategy:
			if !s.cfg.OTP.MailTemplates[reason].Enabled {
				return strategy.DefaultOTPTemplate(reason.String(), s.cfg.OTP.MailTemplates[reason]), nil
			} else {
				return s.cfg.OTP.MailTemplates[reason], nil
			} 
		case strategy.OIDCStrategy:
			return strategy.MailTemplateConfig{}, nil	
		case strategy.PasswordStrategy:
			if !s.cfg.Password.MailTemplates[reason].Enabled {
				return strategy.DefaultMajicLinkTemplate(reason.String(), s.cfg.Password.MailTemplates[reason]), nil
			} else {
				return s.cfg.Password.MailTemplates[reason], nil
			}
	}

	return strategy.MailTemplateConfig{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no valid strategy or reason provided")
}


func (s AuthService) StartAuthFlow(ctx context.Context, input *domain.StartFlowReq) (*domain.StartFlowResp, error) {
	reason := domain.Reason(input.Reason)
	flowMethod := domain.Strategy(input.Strategy) 
	
	// support strategy by the developers
	if !utils.Contains(s.SupportedStrategies(), flowMethod) {
		return nil, ErrUnsupportedMethod
	}

	// valid email format
	validEmail, err := domain.Email(input.Email).NewEmail()

	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidEmailFormat, err)
	}

	// account is already registered
	if (reason == strategy.RegisterReason) {
		auth, err := s.authStore.Select(ctx, validEmail)

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrAuthService.Error())
		}
		
		if (auth != nil) {
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrEmailAlreadyInUse.Error())
		}
	}

	// check to see if user has disabled this strategy
	if (reason == strategy.LoginReason) {
		strategies, err := s.authStore.Select(ctx, validEmail)

		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrAuthService.Error())
		}
		
		var strat bool
		switch flowMethod {
			case strategy.PasswordStrategy:
				strat = strategies.PasswordActive
			case strategy.LinkStrategy: 
				strat = strategies.MajicActive
			case strategy.OTPStrategy: 
				strat = strategies.OTPActive
			case strategy.OIDCStrategy: 
				strat = strategies.OIDCActive
		}

		if (!strat) {
			return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrStrategyNotApplicable.Error())
		}
	}

	var requestStrategy domain.Strategy = flowMethod
	if _, ok := s.cfg.OIDC[flowMethod]; ok {
		requestStrategy = strategy.OIDCStrategy
	}

    startSelectedStrategyFn, ok := startFlowMap[requestStrategy][reason]
    
	if !ok {
        return nil, ErrUnsupportedMethod
    }	

	flow := &domain.Flow{
			Id:        utils.NewUUID(),
			Reason:    reason,	
			Strategy:  flowMethod,
			FinishURL: input.ReturnToURL,
			CreatedAt: s.Now(),
			ExpiresAt: s.Now(),
			Email:     validEmail,
			Metadata: metadata.Metadata{
				"payload": input.Payload,
				"callback_url": input.CallBackUrl,
		},
	}

	return startSelectedStrategyFn(s, ctx, flow)
}

func (s AuthService) FinishAuthFlow(ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	
	flowMethod := domain.Strategy(input.Strategy) 
	reason := domain.Reason(input.Reason)
	
	if !utils.Contains(s.SupportedStrategies(), flowMethod) {
		return nil, ErrUnsupportedMethod
	}

	if _, ok := s.cfg.OIDC[flowMethod]; ok {
		flowMethod = strategy.OIDCStrategy
	}

   	finishFn, ok := finishFlowMap[flowMethod][reason]
    
	if !ok {
        return nil, ErrUnsupportedMethod
    }

	return finishFn(s, ctx, input)
}

func (s AuthService) SupportedStrategies()[]domain.Strategy {
	var strategies []domain.Strategy

	for provider := range s.cfg.OIDC {
		if s.cfg.OIDC[provider].Enabled {
			strategies = append(strategies, provider)
		}
	}

	if s.emailService != nil && s.cfg.Link.Enabled {
		strategies = append(strategies, strategy.LinkStrategy)
	}

	if s.emailService != nil && s.cfg.OTP.Enabled {
		strategies = append(strategies, strategy.OTPStrategy)
	}

	if s.cfg.Password.Enabled {
		strategies = append(strategies, strategy.PasswordStrategy)
	}

	return strategies
} 

func (s AuthService) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	
	validEmail, err := domain.Email(email).NewEmail()

	if err != nil {
		return false, utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidEmailFormat, err)
	}

	model, err := s.authStore.Select(ctx, validEmail)
	
	if (err != nil) {
		return false, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "%s", ErrAuthService)	
	}

	if model != nil {
		return false, nil
	}

	return true, nil
}

// BuildToken creates an access token for the given subjectID
func (s AuthService) BuildToken(ctx context.Context, principal domain.Principal, metadata map[string]string) ([]byte, error) {
	metadata[authToken.SubTypeClaimsKey] = principal.Type.String()
	
	if principal.Type == schema.UserPrincipal && s.cfg.Token.Claims.AddUserEmailClaim {
		metadata[authToken.SubEmailClaimsKey] = principal.Email.String()
	}

	return s.tokenService.Build(principal.Id, metadata)
}

// PublicKeySet returns the public keys to verify the access token
func (s AuthService) PublicKeySet(ctx context.Context) token.JWkSet {
	return s.tokenService.GetPublicKeySet()
}

func (s AuthService) GetPrincipal(ctx context.Context, tkn []byte, assertions ...domain.ClientAssertion) (*domain.Principal, error) {
	if len(assertions) == 0 {
		assertions = domain.APIAssertions
	}

	if slices.Contains[[]domain.ClientAssertion](assertions, domain.SessionClientAssertion) {

		sess, err := s.sessionService.VerifyStrict(ctx, tkn)

		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, "can't get principal: %v",  session.ErrSessionHasExpired)
		}

		userModel, err := s.authStore.SelectUser(ctx, sess.AuthId)

		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to get user")
		}

		user, err := userModel.Transform()

		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "error transforming to user")
		}

		return &domain.Principal{
			Id: user.Id.String(),
			Type: schema.UserPrincipal,
			User: user.Transform(),
		}, nil
	}

	if slices.Contains[[]domain.ClientAssertion](assertions, domain.AccessTokenClientAssertion) {

		insecureTkn, err := token.ParseInsecure(tkn)

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
		}
		
		var genClaim string
		err = insecureTkn.Get(authToken.GeneratedClaimKey, &genClaim)

		if err != nil || genClaim != authToken.GeneratedClaimKey {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
		}

		id, claims, err := s.tokenService.Parse(ctx, tkn)

		authId := utils.UUID(id)

		if err != nil || !utils.IsValidUUID(id) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
		}

		if claims[authToken.SubTypeClaimsKey] == schema.ServiceUserPrincipal {
			serviceUser, err := s.serviceUserService.Get(ctx, authId.String())

			if err != nil {
				return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to get service user") 
			}

			return &domain.Principal{
				Id:          serviceUser.Id,
				Type:        schema.ServiceUserPrincipal,
				ServiceUser: serviceUser,
			}, nil
		}

		userModel, err := s.authStore.SelectUser(ctx, authId)

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to get user")
		}

		user, err := userModel.Transform()

		if (err != nil) {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "error transforming to user")
		}

		return &domain.Principal{
			Id:   user.Id.String(),
			Type: schema.UserPrincipal,
			User: user.Transform(),
		}, nil
	}

	if slices.Contains[[]domain.ClientAssertion](assertions, domain.JWTGrantClientAssertion) {
		serviceUser, err := s.serviceUserService.GetByJWT(ctx, tkn)

		if err != nil {
			return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to get service user") 
		}

		return &domain.Principal{
			Id:          serviceUser.Id,
			Type:        schema.ServiceUserPrincipal,
			ServiceUser: serviceUser,
		}, nil
	}
	
	// unsecure 
	//check for client secret
	// if slices.Contains[[]domain.ClientAssertion](assertions, domain.ClientCredentialsClientAssertion) || 
	//    slices.Contains[[]domain.ClientAssertion](assertions, domain.OpaqueTokenClientAssertion) {
		
		
	// 	userSecretRaw, secretOK := GetSecretFromContext(ctx)

	// 	if secretOK {
	// 		// verify client secret
	// 		userSecret, err := base64.StdEncoding.DecodeString(userSecretRaw)
	// 		if err != nil {
	// 			s.log.Debug("failed to decode user secret", "err", err)
	// 			return Principal{}, errors.ErrUnauthenticated
	// 		}
	// 		userSecretParts := strings.Split(string(userSecret), ":")
	// 		if len(userSecretParts) != 2 {
	// 			s.log.Debug("failed to parse user secret", "err", err)
	// 			return Principal{}, errors.ErrUnauthenticated
	// 		}
	// 		clientID, clientSecret := userSecretParts[0], userSecretParts[1]

	// 		// extract user from secret if it's a service user
	// 		serviceUser, err := s.serviceUserService.GetBySecret(ctx, clientID, clientSecret)

	// 		if err == nil {
	// 			return Principal{
	// 				ID:          serviceUser.ID,
	// 				Type:        schema.ServiceUserPrincipal,
	// 				ServiceUser: &serviceUser,
	// 			}, nil
	// 		}
	// 		if err != nil {
	// 			s.log.Debug("failed to parse as user token ", "err", err)
	// 			return Principal{}, errors.ErrUnauthenticated
	// 		}
	// 	}
	// }

	return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no valid assertions found.")
}











