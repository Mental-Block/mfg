package services

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type AuthService struct {
	userStore       	ports.UserStore
	authStore       	ports.AuthStore
	tokenStore      	ports.TokenStore
	passwordService 	ports.PasswordService
	emailService    	ports.SMTPService
	authTokenService  	ports.TokenService
	emailTokenService 	ports.TokenService
}

func NewAuthService(
	user ports.UserStore,
	auth ports.AuthStore,
	tokenStore ports.TokenStore,
	password ports.PasswordService,
	smtp ports.SMTPService,
	authToken ports.TokenService,
	emailToken ports.TokenService,
) *AuthService {
	return &AuthService{
		userStore:       user,
		authStore:       auth,
		tokenStore:      tokenStore,
		emailService:    smtp,
		authTokenService:    authToken,
		passwordService: password,
		emailTokenService: emailToken,
	}
}

func (s *AuthService) NewRefreshToken(id domain.Id) (*string, error) {
	return s.authTokenService.Create(jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(domain.RefreshTokenDuration).Unix(),
	})
}

func (s *AuthService) NewPasswordResetToken(email domain.Email) (*string, error) {
	return s.emailTokenService.Create(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(domain.PasswordResetTokenDuration).Unix(),
	})
}

func (s *AuthService) NewEmailVerificationToken(email domain.Email) (*string, error) {
	return s.emailTokenService.Create(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(domain.AuthTokenDuration).Unix(),
	})
}

func (s *AuthService) NewAuthToken(id domain.Id, roles []string) (*string, error) {
	return s.authTokenService.Create(jwt.MapClaims{
		"id":    id,
		"roles": roles,
		"exp":   time.Now().Add(domain.AuthTokenDuration).Unix(),
	})
}

func (s *AuthService) Permission() {
	
	
	//s.userStore.

	//s.authStore.SelectUser(ctx, )


}

func (s *AuthService) LoginOAuth(ctx context.Context) {

}

func (s *AuthService) Login(ctx context.Context, email string, password string) (*string, error) {
	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}
	
	validPassword, err := domain.NewPassword(password)

	if err != nil {
		return nil, fmt.Errorf("invalid password supplied: %w", err)
	}

	user, err := s.userStore.Select(ctx, validEmail)

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrUserStore
	}

	if isVerified := s.passwordService.Verify(validPassword, user.Password); !isVerified {
		return nil, ErrIncorrectPassword
	}

	token, err := s.NewRefreshToken(user.Id)
	
	if err != nil {
		slog.Error(err.Error())
		return nil, ErrSigningToken
	}

	key := strconv.Itoa(int(user.Id))

	err = s.tokenStore.Insert(ctx, key, *token, domain.RefreshTokenDuration) // insert refresh token

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrTokenStore;
	}

	return token, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	claims, err := s.authTokenService.Parse(token)

	if err != nil {
		return ErrInvalidToken
	}

	id := claims["id"].(int)

	key := strconv.Itoa(id)
	
	err = s.tokenStore.Remove(ctx, key) // remove refresh token 

	if err != nil {
		slog.Error(err.Error())
		return ErrTokenStore;
	}

	return nil
}

func (s *AuthService) RegisterOAuth(ctx context.Context, email string) {

}

func (s *AuthService) Register(ctx context.Context, email string, username string, password string) (error) {

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return fmt.Errorf("invalid email supplied: %w", err)
	}
	
	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return fmt.Errorf("invalid username supplied: %w", err)
	}

	validPassword, err := domain.NewPassword(password)

	if err != nil {
		return fmt.Errorf("invalid password supplied: %w", err)
	}

	user, err := s.userStore.Select(ctx, validEmail)

	if err != nil && err != domain.ErrDataNotFound {
		slog.Error(err.Error())
		return  ErrUserAuthStore
	}

	if user != nil && user.Email == validEmail {
		return  fmt.Errorf("email address is already registered")
	}

	err = s.emailService.DNSLookUp(string(validEmail))
	
	if (err != nil) {
		return err
	}

	token, err := s.NewEmailVerificationToken(validEmail)

	if err != nil {
		slog.Error(err.Error())
		return ErrSigningToken
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8080/confirm-account/%s", *token)

	err = s.emailService.Send(
		[]string{"aaron.tibben@gmail.com"},
		"Subject:Do Not Reply - Finish Registration\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		s.emailService.VerificationTemplate(validationEndPoint),
	)

	if err != nil {
		slog.Error(err.Error())
		return  ErrEmailService
	}

	hashedPassword := s.passwordService.Hash(validPassword)

	err = s.authStore.InsertUser(
		ctx, 
		validEmail, 
		hashedPassword, 
		validUsername,
		*token,
	)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	return nil
}

func (s *AuthService) Verify(ctx context.Context, token string) (string, error) {
	claims, err := s.emailTokenService.Parse(token)

	if err != nil {
		return "", ErrInvalidToken
	}

	email := claims["email"].(string)

	return email, nil
}

func (s *AuthService) FinishRegister(ctx context.Context, token string) error {
	email, err := s.Verify(ctx, token)

	if (err != nil) {
		return err
	}

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return fmt.Errorf("invalid email supplied: %w", err)
	}

	user, err := s.authStore.SelectUser(context.Background(), validEmail)

	if (err != nil) {
		slog.Error(err.Error())
		return ErrUserStore
	}

	// prevent multiple active tokens, token must match current token
	if (user.Token != token) {
		return ErrInvalidToken
	}

	_, err = s.userStore.Insert(ctx, user.Email, user.Password, user.Username, false)
    
	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	err = s.authStore.RemoveUser(ctx, validEmail)

	if err != nil {
		return err
	}

	return nil 
}

func (s *AuthService) UpdatePassword(ctx context.Context, token string, password string) error {
	email, err := s.Verify(ctx, token)

	if (err != nil) {
		slog.Error(token)
		return err
	}

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return fmt.Errorf("invalid email supplied: %w", err)
	}

	newPassword, err := domain.NewPassword(password)

	if err != nil {
		return fmt.Errorf("invalid password supplied: %w", err)
	}

	user, err := s.userStore.Select(ctx, validEmail)

	if (err != nil) {
		slog.Error(err.Error());
		return err
	}

	curToken, err := s.tokenStore.Select(ctx, string(validEmail))

	if (err != nil) {
		slog.Error(err.Error())
		return err
	}
	
	// prevent multiple active tokens, token must match current token
	if (*curToken != token) {
		return  ErrInvalidToken
	}

	hashedpas := s.passwordService.Hash(newPassword)

	err = s.authStore.UpdatePassword(ctx, validEmail, hashedpas)

	if err != nil {
		return ErrUserStore
	}

	err = s.tokenStore.Remove(ctx, email)

	if err != nil {
		slog.Error(err.Error())
		return ErrTokenStore
	}

	key := strconv.Itoa(int(user.Id))
	
	err = s.tokenStore.Remove(ctx, key)
	
	if err != nil {
		slog.Error(err.Error())
		return ErrTokenStore;
	}

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, email string) error {
	
	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return fmt.Errorf("invalid email supplied: %w", err)
	}

	token, err := s.NewPasswordResetToken(validEmail)

	if err != nil {
		slog.Error(err.Error())
		return ErrSigningToken
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8080/confirm-account-reset/%s", *token)

	err = s.emailService.Send(
		[]string{"aaron.tibben@gmail.com"},
		"Subject:Do Not Reply - Reset Password\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		s.emailService.RestPasswordTemplate(validationEndPoint),
	)

	if err != nil {
		slog.Error(err.Error())
		return ErrEmailService
	}

	err = s.tokenStore.Insert(ctx, string(validEmail), *token, domain.PasswordResetTokenDuration)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
