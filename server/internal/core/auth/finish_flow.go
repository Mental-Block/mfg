package auth

import (
	"context"
	"crypto/subtle"
	"fmt"

	"github.com/server/internal/adapters/store/postgres/tx"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/internal/core/auth/session"
	"github.com/server/internal/core/auth/strategy"
	user "github.com/server/internal/core/user/domain"

	"github.com/server/pkg/utils"
)

type FinishFlowFn func(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error)

var finishFlowMap = map[domain.Strategy]map[domain.Reason]FinishFlowFn{
	 strategy.PasswordStrategy: {
		strategy.ForgotPasswordReason: finishResetPasswordFlow,
		strategy.LoginReason: finishLoginPasswordFlow,
		strategy.RegisterReason: finishRegisterPasswordFlow,
	},
	strategy.OIDCStrategy: {
		strategy.LoginReason: finishOIDCLoginFlow,
		strategy.RegisterReason: finishRegisterOIDCFlow,
	},
	strategy.LinkStrategy: {
		strategy.LoginReason: finshLoginMailLinkFlow,
	},
	strategy.OTPStrategy: {
		strategy.LoginReason: finishLoginOTPFlow,
		strategy.RegisterReason: finishRegisterOTPFlow,
	},
}

func finishLoginPasswordFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error){
	
	flowId, err := utils.ConvertStringToUUID(input.State)

	if err != nil {
		return nil, err
	}
	
	key := s.flowStore.GenerateKey(flowId.String())
	flow, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, err
	}

	authId, err := utils.ConvertStringToUUID(flow.Metadata["authId"].(string))

	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no authId found")
	}

	userModel, err := s.authStore.SelectUser(ctx, authId)

	if (err != nil) {
		return nil, err
	}

	santizedUser, err := userModel.Transform()

	if (err != nil) {
		return nil, err
	}

	version, ok := flow.Metadata["version"].(int)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no version found")
	}  

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: authId,
		Email: flow.Email,
		Version: version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}


	return &domain.FinishFlowResp{
		User: santizedUser.Transform(),
		Flow: &domain.Flow{
			Id: authId,
		},
	}, nil
}

func finishRegisterPasswordFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	
	if len(input.Code) == 0 {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(input.State)

	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
	flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()
	
	if !flow.IsValid(s.Now()) {
		return nil, ErrFlowInvalid
	}

	if subtle.ConstantTimeCompare([]byte(flow.Nonce), []byte(input.Code)) == 0 {
		// avoid brute forcing otp
		attemptInt := 0
		if attempts, ok := flow.Metadata[strategy.OtpAttemptKey]; ok {
			attemptInt, _ = attempts.(int)
		}

		if attemptInt < strategy.MaxOTPAttempt {
			flow.Metadata[strategy.OtpAttemptKey] = attemptInt + 1
			
			if err = s.flowStore.Insert(ctx, key, flow, flow.ExpiresAt.UTC().Sub(flow.ExpiresAt)); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		} else {
			if err = s.flowStore.Delete(ctx, key); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		}

		return nil, ErrInvalidMailOTP
	}

	u, ok := flow.Metadata["username"].(string)

	if (!ok) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "missing username")
	}

	p, ok :=  flow.Metadata["password"].(string)

	if (!ok) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "missing password")
	}

	authId := utils.NewUUID()
	version := 0

	user := user.SanitizedUser{
		Id:	utils.NewUUID(),
		Username: user.Username(u),
		Active: true,
	}
	
	cmd := tx.TXAuthCmd{ 
		User:  user,
		Auth: domain.SanitizedAuth{
			Id: authId,
			Email: flow.Email,
			Password: domain.Password(p),
			PasswordActive: true,
			Version: version,
		},
	}
	
	err = s.transactService.UserCreationTransaction(ctx, cmd)

	if (err != nil) {
		return nil, err
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: authId,
		Email: flow.Email,
		Version: version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}


	return &domain.FinishFlowResp{
		User: user.Transform(),
		Flow: &flow,
	}, nil
}

func finishResetPasswordFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error){
	password, err := domain.Password(input.StateConfig["password"].(string)).NewPassword()

	if err != nil {
		return nil,  utils.NewErrorf(utils.ErrorCodeInvalidArgument, ErrInvalidPasswordFormat, err)
	}
	
	if len(input.Code) == 0 {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(input.State)

	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
	flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()
	
	if !flow.IsValid(s.Now()) {
		return nil, ErrFlowInvalid
	}

	if subtle.ConstantTimeCompare([]byte(flow.Nonce), []byte(input.Code)) == 0 {
		// avoid brute forcing otp
		attemptInt := 0
		if attempts, ok := flow.Metadata[strategy.OtpAttemptKey]; ok {
			attemptInt, _ = attempts.(int)
		}

		if attemptInt < strategy.MaxOTPAttempt {
			flow.Metadata[strategy.OtpAttemptKey] = attemptInt + 1
			
			if err = s.flowStore.Insert(ctx, key, flow, flow.ExpiresAt.UTC().Sub(flow.ExpiresAt)); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		} else {
			if err = s.flowStore.Delete(ctx, key); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		}

		return nil, ErrInvalidMailOTP
	}

	hash, err := strategy.NewPassword(s.cfg.Password.Params).CreateHash(password.String())

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	auth, err := s.authStore.Select(ctx, domain.Email(flow.Email))

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = s.sessionService.RemoveAllUserSessions(ctx, auth.Id.String())

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	
	err = s.authStore.Update(ctx, domain.SanitizedAuth{
		Id: auth.Id,
		Email: auth.Email,
		Password: hash,
		Version: auth.Version + 1,
	})

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	return &domain.FinishFlowResp{
		Flow: &flow,
	}, nil
}

func finishLoginOTPFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	if len(input.Code) == 0 {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(input.State)

	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
	flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()
	
	if !flow.IsValid(s.Now()) {
		return nil, ErrFlowInvalid
	}

	if subtle.ConstantTimeCompare([]byte(flow.Nonce), []byte(input.Code)) == 0 {
		// avoid brute forcing otp
		attemptInt := 0
		if attempts, ok := flow.Metadata[strategy.OtpAttemptKey]; ok {
			attemptInt, _ = attempts.(int)
		}

		if attemptInt < strategy.MaxOTPAttempt {
			flow.Metadata[strategy.OtpAttemptKey] = attemptInt + 1
			
			if err = s.flowStore.Insert(ctx, key, flow, flow.ExpiresAt.UTC().Sub(flow.ExpiresAt)); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		} else {
			if err = s.flowStore.Delete(ctx, key); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		}

		return nil, ErrInvalidMailOTP
	}

	auth, err := s.authStore.Select(ctx, flow.Email)

	if (err != nil) {
		return nil, err
	}

	usrModel, err := s.authStore.SelectUser(ctx, auth.Id)
		
	if (err != nil) {
		return nil, err
	}

	usr, err := usrModel.Transform()

	if (err != nil) {
		return nil, err
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: auth.Id,
		Email: flow.Email,
		Version: auth.Version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}

	return &domain.FinishFlowResp{
		User: usr.Transform(),
		Flow: &flow,
	}, nil
}

func finishRegisterOTPFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	if len(input.Code) == 0 {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(input.State)

	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
	flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()

	if !flow.IsValid(s.Now()) {
		return nil, ErrFlowInvalid
	}

	if subtle.ConstantTimeCompare([]byte(flow.Nonce), []byte(input.Code)) == 0 {
		// avoid brute forcing otp
		attemptInt := 0
		if attempts, ok := flow.Metadata[strategy.OtpAttemptKey]; ok {
			attemptInt, _ = attempts.(int)
		}

		if attemptInt < strategy.MaxOTPAttempt {
			flow.Metadata[strategy.OtpAttemptKey] = attemptInt + 1
			
			if err = s.flowStore.Insert(ctx, key, flow, flow.ExpiresAt.UTC().Sub(flow.ExpiresAt)); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		} else {
			if err = s.flowStore.Delete(ctx, key); err != nil {
				return nil, fmt.Errorf("failed to process flow code missmatch")
			}
		}

		return nil, ErrInvalidMailOTP
	}

	username, ok :=  flow.Metadata["username"].(string)

	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, user.ErrInvalidUsernameFormat)
	}

	authId := utils.NewUUID()
	version := 0

	user := user.SanitizedUser{
		Id:	utils.NewUUID(),
		Username: user.Username(username),
	}
	
	cmd := tx.TXAuthCmd{ 
		User:  user,
		Auth: domain.SanitizedAuth{
			Id: authId,
			Email: flow.Email,
			OTPActive: true,
			Version: version,
		},
	}
	
	err = s.transactService.UserCreationTransaction(ctx, cmd)

	if err != nil {
		return nil, err
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: authId,
		Email: flow.Email,
		Version: version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}

	return &domain.FinishFlowResp{
		User: user.Transform(),
		Flow: &flow,
	}, nil
}

func finshLoginMailLinkFlow(s AuthService, ctx context.Context, flow *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	return finishLoginOTPFlow(s, ctx, flow)
}

func finishOIDCLoginFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	// flow id is added in state params
	if len(input.State) == 0 {
		return nil, ErrInvalidOIDCState
	}

	// flow id is added in state params
	if len(input.Code) == 0 {
		return nil, ErrMissingOIDCCode
	}

	// check for oidc flow via fetching oauth state, method parameter will not be set for oauth
	flowIDFromState, err := strategy.ExtractFlowFromOIDCState(input.State)
	
	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(flowIDFromState)
	
	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
	flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()

	oidcConfig, ok := s.cfg.OIDC[flow.Strategy]

	if !ok {
		return nil, ErrStrategyNotApplicable
	}

	cb, ok := flow.Metadata["callback_url"].(string)

	if !ok {
		return nil, fmt.Errorf("callback url not configured")
	}

	idp, err := strategy.NewRelyingPartyOIDC(
		oidcConfig.ClientId,
		oidcConfig.ClientSecret,
		cb).
		Init(ctx, oidcConfig.IssuerUrl)
	
	if err != nil {
		return nil, err
	}

	oAuthToken, err := idp.Token(ctx, input.Code, flow.Nonce)
	
	if err != nil {
		return nil, err
	}
	
	oAuthProfile, err := idp.GetUser(ctx, oAuthToken)
	
	if err != nil {
		return nil, err
	}

	auth, err := s.authStore.Select(ctx, domain.Email(oAuthProfile.Email))

	if (err != nil) {
		return nil, err
	}

	usrModel, err := s.authStore.SelectUser(ctx, auth.Id)

	if (err != nil) {
		return nil, err
	}

	usr, err := usrModel.Transform()

	if (err != nil) {
		return nil, err
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: auth.Id,
		Email: flow.Email,
		Version: auth.Version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}


	return &domain.FinishFlowResp{
		User: usr.Transform(),
		Flow: &flow,
	}, nil
}

func finishRegisterOIDCFlow(s AuthService, ctx context.Context, input *domain.FinishFlowReq) (*domain.FinishFlowResp, error) {
	// flow id is added in state params
	if len(input.State) == 0 {
		return nil, ErrInvalidOIDCState
	}

	// flow id is added in state params
	if len(input.Code) == 0 {
		return nil, ErrMissingOIDCCode
	}

	// check for oidc flow via fetching oauth state, method parameter will not be set for oauth
	flowIDFromState, err := strategy.ExtractFlowFromOIDCState(input.State)
	
	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	flowId, err := utils.ConvertStringToUUID(flowIDFromState)
	
	if err != nil {
		return nil, ErrStrategyNotApplicable
	}

	key := s.flowStore.GenerateKey(flowId.String())
		flowModel, err := s.flowStore.Select(ctx, key)

	if err != nil {
		return nil, fmt.Errorf("invalid state for mail otp: %w", err)
	}

	flow := flowModel.Transform()

	oidcConfig, ok := s.cfg.OIDC[flow.Strategy]

	if !ok {
		return nil, ErrStrategyNotApplicable
	}

	cb, ok := flow.Metadata["callback_url"].(string)

	if !ok {
		return nil, fmt.Errorf("callback url not configured")
	}

	idp, err := strategy.NewRelyingPartyOIDC(
		oidcConfig.ClientId,
		oidcConfig.ClientSecret,
		cb).
		Init(ctx, oidcConfig.IssuerUrl)
	
	if err != nil {
		return nil, err
	}

	authToken, err := idp.Token(ctx, input.Code, flow.Nonce)
	
	if err != nil {
		return nil, err
	}

	oAuthProfile, err := idp.GetUser(ctx, authToken)
	
	if err != nil {
		return nil, err
	}

	authId := utils.NewUUID()
	version := 0

	user := user.SanitizedUser{
		Id:	utils.NewUUID(),
		Username: user.Username(oAuthProfile.Name),
	}
	
	cmd := tx.TXAuthCmd{ 
		User:  user,
		Auth: domain.SanitizedAuth{
			Id: authId,
			Email: domain.Email(oAuthProfile.Email),
			OIDCActive: true,
			Version: version,
		},
	}
	
	err = s.transactService.UserCreationTransaction(ctx, cmd)

	if err != nil {
		return nil, err
	}

	if err = s.flowStore.Delete(ctx, key); err != nil {
		return nil, err
	}

	_, err = s.sessionService.New(ctx, domain.SanitizedAuth{
		Id: authId,
		Email: flow.Email,
		Version: version,
	})

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, session.ErrCouldntCreateSession.Error())
	}

	return &domain.FinishFlowResp{
		User: user.Transform(),
		Flow: &flow,
	}, nil
}
