package auth

import (
	"context"
	"time"

	"github.com/server/internal/core/auth/domain"
	"github.com/server/internal/core/auth/strategy"
	user "github.com/server/internal/core/user/domain"

	"github.com/server/pkg/utils"
)

type StartFlowFn func(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error)

var startFlowMap = map[domain.Strategy]map[domain.Reason]StartFlowFn {
    strategy.PasswordStrategy: {
		strategy.ForgotPasswordReason: startResetPasswordFlow,
		strategy.LoginReason: startLoginPasswordFlow,
		strategy.RegisterReason: startRegisterPasswordFlow,
	},
	strategy.OIDCStrategy: {
		strategy.LoginReason: startOIDCLoginFlow,
		strategy.RegisterReason: startOIDCRegisterFlow,
	},
	strategy.LinkStrategy: {
		strategy.LoginReason: startLoginMailLinkFlow,
	},
	strategy.OTPStrategy: {
		strategy.LoginReason: startLoginOTPFlow,
		strategy.RegisterReason: startRegisterOTPFlow,
	},
}


func startLoginOTPFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	
	mailTemplate, err := s.findTemplate(flow.Reason, flow.Strategy)
	
	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to set mail template")
	}
	
	mailOtpStrat := strategy.NewMailOTP(
		s.emailService, 
		mailTemplate.Subject, 
		mailTemplate.Body,
	)
	
	nonce, err := mailOtpStrat.SendMail(flow.Id.String(), flow.Email.String())

	if err != nil {
		return nil, err
	}

	flow.Nonce = nonce

	duration := strategy.AccountLoginEmailDuration
	if (mailTemplate.Validity != 0) {
		duration = mailTemplate.Validity
	} 

	flow.ExpiresAt = flow.ExpiresAt.Add(duration)

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow:  flow,
	}, nil
}

func startRegisterOTPFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	payload, ok := flow.Metadata["payload"].(map[string]string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no payload provided")
	}

	username, err := user.NewUsername(payload["username"])

	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, user.ErrInvalidUsernameFormat, err)
	}

	mailTemplate, err := s.findTemplate(flow.Reason, flow.Strategy)
	
	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to set mail template")
	}

	mailOtpStrat := strategy.NewMailOTP(
		s.emailService, 
		mailTemplate.Subject, 
		mailTemplate.Body,
	)
	
	nonce, err := mailOtpStrat.SendMail(flow.Id.String(), flow.Email.String())

	if err != nil {
		return nil, err
	}

	flow.Nonce = nonce
	flow.Metadata["username"] = username

	duration := strategy.AccountLoginEmailDuration
	if (mailTemplate.Validity != 0) {
		duration = mailTemplate.Validity
	} 

	flow.ExpiresAt = flow.ExpiresAt.Add(duration)

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow:  flow,
	}, nil
}

func startLoginMailLinkFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	cb, ok := flow.Metadata["callback_url"].(string)
	
	if !ok || len(cb) == 0 {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "callback url not configured")
	}

	mailTemplate, err := s.findTemplate(flow.Reason, flow.Strategy)
	
	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to set mail template")
	}
	
	mailLinkStrat := strategy.NewMailLink(
		s.emailService, 
		cb, 
		flow.FinishURL, 
		mailTemplate.Subject, 
		mailTemplate.Body,
	)

	nonce, err := mailLinkStrat.SendMail(flow.Id.String(), flow.Email.String())
	
	if err != nil {
		return nil, err
	}

	flow.Nonce = nonce

	duration := strategy.AccountLoginEmailDuration
	if (mailTemplate.Validity != 0) {
		duration = mailTemplate.Validity
	} 

	flow.ExpiresAt = flow.ExpiresAt.Add(duration)

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}

func startResetPasswordFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	cb, ok := flow.Metadata["callback_url"].(string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no call back provided")
	}

	mailTemplate, err := s.findTemplate(flow.Reason, flow.Strategy)
	
	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to set mail template")
	}
	
	mailLinkStrat := strategy.NewMailLink(
		s.emailService, 
		cb, 
		"",
		mailTemplate.Subject, 
		mailTemplate.Body,
	)

	nonce, err := mailLinkStrat.SendMail(flow.Id.String(), flow.Email.String())

	flow.Nonce = nonce

	duration := strategy.AccountForgotPasswordEmailDuration
	if (mailTemplate.Validity != 0) {
		duration = mailTemplate.Validity
	} 

	flow.ExpiresAt = flow.ExpiresAt.Add(duration)

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}

func startLoginPasswordFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	payload, ok := flow.Metadata["payload"].(map[string]string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no payload provided")
	}

	if _, ok := flow.Metadata["callback_url"].(string); !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no call back provided")
	}

	auth, err := s.authStore.Select(ctx, flow.Email)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, ErrAuthService.Error())
	}

	if _, err := strategy.IsPasswordValid(payload["password"], auth.Password); err != nil {
		return nil, err
	}

	flow.Metadata["authId"] = auth.Id
	flow.Metadata["version"] = auth.Version

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, 2 * time.Minute); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}

func startRegisterPasswordFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	payload, ok := flow.Metadata["payload"].(map[string]string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no payload provided")
	}

	cb, ok := flow.Metadata["callback_url"].(string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no call back provided")
	}

	if err := domain.Email(flow.Email).DNSLookUp(); err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "dns look up failed")
	}

	username, err := user.NewUsername(payload["username"])

	if err != nil {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, user.ErrInvalidUsernameFormat, err)
	}

	password, err := domain.Password(payload["password"]).NewPassword()

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "invalid password format")
	}

	hash, err := strategy.NewPassword(s.cfg.Password.Params).CreateHash(password.String())

	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeInvalidArgument, "invalid password format")
	}

	mailTemplate, err := s.findTemplate(flow.Reason, flow.Strategy)
	
	if (err != nil) {
		return nil, utils.NewErrorf(utils.ErrorCodeUnknown, "failed to set mail template")
	}

	mailLinkStrat := strategy.NewMailLink(s.emailService, 
		cb, 
		"", 
		mailTemplate.Subject, 
		mailTemplate.Body,
	)

	nonce, err := mailLinkStrat.SendMail(flow.Id.String(), flow.Email.String())
 
	if (err != nil) {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "failed to send email")
	}

	flow.Metadata["username"] = username
	flow.Metadata["password"] = hash
	flow.Nonce = nonce
	
	duration := strategy.AccountVerificationEmailDuration
	if (mailTemplate.Validity != 0) {
		duration = mailTemplate.Validity
	} 

	flow.ExpiresAt = flow.ExpiresAt.Add(duration)

	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}

func startOIDCLoginFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	cb, ok := flow.Metadata["callback_url"].(string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no call back provided")
	}

	idp, err := strategy.NewRelyingPartyOIDC(
			s.cfg.OIDC[flow.Strategy].ClientId,
			s.cfg.OIDC[flow.Strategy].ClientSecret,
			cb,
	).Init(ctx, s.cfg.OIDC[flow.Strategy].IssuerUrl)
		
	if err != nil {
		return nil, err
	}

	oidcState, err := strategy.EmbedFlowInOIDCState(flow.Id.String())

	if err != nil {
		return nil, err
	}
	
	endpoint, nonce, err := idp.AuthURL(oidcState)
		
	if err != nil {
		return nil, err
	}

	flow.StartURL = endpoint
	flow.Nonce = nonce
	duration := strategy.AccountLoginEmailDuration
	flow.ExpiresAt = flow.CreatedAt.Add(duration)
	
	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}

func startOIDCRegisterFlow(s AuthService, ctx context.Context, flow *domain.Flow) (*domain.StartFlowResp, error) {
	cb, ok := flow.Metadata["callback_url"].(string)
	
	if !ok {
		return nil, utils.NewErrorf(utils.ErrorCodeInvalidArgument, "no call back provided")
	}

	idp, err := strategy.NewRelyingPartyOIDC(
			s.cfg.OIDC[flow.Strategy].ClientId,
			s.cfg.OIDC[flow.Strategy].ClientSecret,
			cb,
	).Init(ctx, s.cfg.OIDC[flow.Strategy].IssuerUrl)
		
	if err != nil {
		return nil, err
	}

	oidcState, err := strategy.EmbedFlowInOIDCState(flow.Id.String())

	if err != nil {
		return nil, err
	}
	
	endpoint, nonce, err := idp.AuthURL(oidcState)
		
	if err != nil {
		return nil, err
	}

	flow.StartURL = endpoint
	flow.Nonce = nonce

	duration := strategy.AccountLoginEmailDuration
	
	flow.ExpiresAt = flow.CreatedAt.Add(duration)
	
	key := s.flowStore.GenerateKey(flow.Id.String())
	if err = s.flowStore.Insert(ctx, key, *flow, duration); err != nil {
		return nil, err
	}

	return &domain.StartFlowResp{
		Flow: flow,
	}, nil
}