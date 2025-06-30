package auth

import "errors"

	
var (
	ErrInvalidPasswordFormat = "invalid password supplied: %s"
	ErrInvalidEmailFormat = "invalid email supplied: %s"	
	
	ErrAuthService       	 = errors.New("auth service error")
  	ErrEmailAlreadyInUse 	 = errors.New("email address is already registered")
	ErrInvalidToken      	 = errors.New("token not valid")
	ErrIncorrectPassword 	 = errors.New("incorrect email or password")
	ErrStrategyNotApplicable = errors.New("strategy not applicable")
	ErrUnsupportedMethod     = errors.New("unsupported authentication method")
	ErrInvalidMailOTP        = errors.New("invalid mail otp")
	ErrMissingOIDCCode       = errors.New("OIDC code is missing")
	ErrInvalidOIDCState      = errors.New("invalid auth state")
	ErrFlowInvalid           = errors.New("invalid flow or expired")
)