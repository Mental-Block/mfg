package services

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/server/internal"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type AuthService struct {
	authUserStore 		ports.AuthUserStore
	userStore			ports.UserStore
	authStore       	ports.AuthStore
	tokenStore      	ports.TokenStore
	passwordService 	ports.PasswordService
	emailService    	ports.SMTPService
	refreshTokenService ports.TokenService
	authTokenService  	ports.TokenService
	emailTokenService 	ports.TokenService
}

func NewAuthService(
	authUser ports.AuthUserStore,
	user ports.UserStore,
	auth ports.AuthStore,
	tokenStore ports.TokenStore,
	password ports.PasswordService,
	smtp ports.SMTPService,
	refreshToken ports.TokenService,
	authToken ports.TokenService,
	emailToken ports.TokenService,
) *AuthService {
	return &AuthService{
		authUserStore: 	 authUser,
		userStore: 	 	 user,
		authStore:       auth,
		tokenStore:      tokenStore,
		emailService:    smtp,
		passwordService: password,
		refreshTokenService: refreshToken,
		authTokenService: authToken,
		emailTokenService: emailToken,
	}
}

func (s *AuthService) newRefreshToken(authId domain.Id, version int) (*string, error) {
	return s.refreshTokenService.Create(jwt.MapClaims{
		"id": authId,
		"version": version,
		"exp": time.Now().Add(domain.RefreshTokenDuration).Unix(),
	})
}

func (s *AuthService) newAuthToken(userId domain.Id, roles []string) (*string, error) {
	return s.authTokenService.Create(jwt.MapClaims{
		"id": userId,
		"roles": roles,
		"exp":   time.Now().Add(domain.AuthTokenDuration).Unix(),
	})
}

func (s *AuthService) newPasswordResetToken(email domain.Email) (*string, error) {
	return s.emailTokenService.Create(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(domain.PasswordResetTokenDuration).Unix(),
	})
}

func (s *AuthService) newEmailVerificationToken(email domain.Email) (*string, error) {
	return s.emailTokenService.Create(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(domain.EmailVerificationToken).Unix(),
	})
}

func (s *AuthService) Permission(ctx context.Context, token string) (*string, *string, *domain.UserAuth, error)  {
	claims, err := s.refreshTokenService.Parse(token)

	if err != nil {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}
	
	authId, ok := claims["id"].(float64)

	if (!ok) {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	version, ok := claims["version"].(float64)

	if (!ok) {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	validVersion := int(version)
	validAuthId := domain.NewId(int(authId))

	currentVersion, err := s.authStore.SelectVersion(ctx, validAuthId)

	if err != nil {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	if (*currentVersion != validVersion) {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	authUser, err := s.authUserStore.Select(ctx, validAuthId)

	if err != nil {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserRoleStore.Error())
	}
	
	authToken, err := s.newAuthToken(authUser.Id, authUser.Roles) 

	if (err != nil) {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	refreshToken, err := s.newRefreshToken(validAuthId, validVersion)
	
	if (err != nil) {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	return refreshToken, authToken, authUser, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (*string, *string, *domain.UserAuth, error) {
	
	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, nil, nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()))
	}
	
	validPassword, err := domain.NewPassword(password)

	if err != nil {
		return nil, nil, nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid password supplied: %s", err.Error()))
	}

	auth, err := s.authStore.Select(ctx, validEmail)

	if (err != nil) {
		if (err.Error() == domain.ErrAuthNotFound.Error()) {
			return nil, nil, nil, internal.NewErrorf(internal.ErrorCodeNotAuthorized, domain.ErrIncorrectPassword.Error())
		}

		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())	
	}

	authUser, err := s.authUserStore.Select(ctx, auth.Id)
	
	if err != nil {
			return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserRoleStore.Error())	
	}

	if isVerified, err := s.passwordService.ComparePasswordAndHash(string(validPassword), string(auth.Password)); !isVerified || err != nil {
		return nil, nil, nil, internal.NewErrorf(internal.ErrorCodeNotAuthorized, domain.ErrIncorrectPassword.Error())
	}

	refreshToken, err := s.newRefreshToken(auth.Id, auth.Version)
	
	if err != nil {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	authtoken, err := s.newAuthToken(authUser.Id, authUser.Roles)

	if err != nil {
		return nil, nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	return refreshToken, authtoken, authUser, nil
}

func (s *AuthService) Register(ctx context.Context, email string, username string, password string) error {

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()) )
	}
	
	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid username supplied: %s", err.Error()) )
	}

	validPassword, err := domain.NewPassword(password)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid password supplied: %s", err.Error()))
	}

	auth, err := s.authStore.Select(ctx, validEmail)

	if err != nil && err.Error() != domain.ErrAuthNotFound.Error() {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	if auth != nil && auth.Email == validEmail {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, ErrEmailAlreadyInUse.Error())
	}

	err = s.emailService.DNSLookUp(string(validEmail))
	
	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrEmailService.Error())
	}

	token, err := s.newEmailVerificationToken(validEmail)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8080/confirm-account/%s", *token)

	err = s.emailService.Send(
		[]string{"aaron.tibben@gmail.com"},
		"Subject:Do Not Reply - Finish Registration\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		s.emailService.VerificationTemplate(validationEndPoint),
	)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrEmailService.Error())
	}

	hashedPassword, err := s.passwordService.CreateHash(string(validPassword))

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	err = s.authStore.InsertCache(
		ctx, 
		validEmail, 
		domain.Password(hashedPassword), 
		validUsername,
		*token,
	)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrTokenStore.Error())
	}

	return nil
}

func (s *AuthService) Verify(ctx context.Context, token string) (*string, error) {
	claims, err := s.emailTokenService.Parse(token)

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	email, ok := claims["email"].(string)

	if !ok {
		return nil, internal.NewErrorf(internal.ErrorCodeUnknown, "could not convert to string")
	}

	return &email, nil
}

func (s *AuthService) RegisterFinish(ctx context.Context, token string) error {
	email, err := s.Verify(ctx, token)

	if (err != nil) {
		return internal.NewErrorf(internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	validEmail, err := domain.NewEmail(*email)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()) )
	}

	user, err := s.authStore.SelectCache(ctx, validEmail)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	// prevent multiple active tokens, token must match current token
	if token != user.Token {
		return internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}

	_, err = s.userStore.Insert(ctx, user.Username, user.Email, user.Password, false)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrUserStore.Error())
	}
	
	err = s.authStore.DeleteCache(ctx, validEmail)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	return nil 
}

func (s *AuthService) UpdatePassword(ctx context.Context, token string, password string) error {
	email, err := s.Verify(ctx, token)

	if (err != nil) {
		return err
	}

	validEmail, err := domain.NewEmail(*email)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()) )
	}

	newPassword, err := domain.NewPassword(password)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid password supplied: %s", err.Error()))
	}

	curToken, err := s.tokenStore.Select(ctx, string(validEmail))

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrTokenStore.Error())
	}
	
	// prevent multiple active tokens, token must match current token
	if (*curToken != token) {
		return internal.WrapErrorf(err, internal.ErrorCodeNotAuthorized, ErrInvalidToken.Error())
	}
	
	hashedpas, err := s.passwordService.CreateHash(string(newPassword))

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	// updating password will also make all current refresh tokens invalid with this user invalid incase user got hacked
	err = s.authStore.UpdatePassword(ctx, validEmail, domain.Password(hashedpas))

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrAuthStore.Error())
	}

	err = s.tokenStore.Delete(ctx, *email)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrTokenStore.Error())
	}

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, email string) error {
	
	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, fmt.Sprintf("invalid email supplied: %s", err.Error()) )
	}

	token, err := s.newPasswordResetToken(validEmail)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrSigningToken.Error())
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8080/confirm-account-reset/%s", *token)

	err = s.emailService.Send(
		[]string{"aaron.tibben@gmail.com"},
		"Subject:Do Not Reply - Reset Password\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		s.emailService.RestPasswordTemplate(validationEndPoint),
	)

	if (err != nil) {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrEmailService.Error())
	}

	err = s.tokenStore.Insert(ctx, string(validEmail), *token, domain.PasswordResetTokenDuration)

	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, ErrTokenStore.Error())
	}

	return nil
}
