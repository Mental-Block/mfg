package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/server/internal/core/domain"
	"github.com/server/internal/core/ports"
)

type AuthService struct {
	userStore       ports.UserStore
	authStore       ports.UserAuthStore
	passwordService ports.PasswordService
	emailService    ports.SMTPService
	jwtService      ports.TokenService
}

func NewAuthService(
	user ports.UserStore,
	auth ports.UserAuthStore,
	smtp ports.SMTPService,
	jwt ports.TokenService,
	password ports.PasswordService,
) *AuthService {
	return &AuthService{
		userStore:       user,
		authStore:       auth,
		emailService:    smtp,
		jwtService:      jwt,
		passwordService: password,
	}
}

func (s *AuthService) NewRefreshToken() (*string, error) {
	return s.jwtService.CreateToken(jwt.MapClaims{
		"exp": time.Now().Add((time.Hour * 24 * 30)).Unix(),
	})
}

func (s *AuthService) newPasswordResetToken(email domain.Email) (*string, error) {
	return s.jwtService.CreateToken(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add((time.Minute * 5)).Unix(),
	})
}

func (s *AuthService) newVerificationToken(email domain.Email) (*string, error) {
	return s.jwtService.CreateToken(jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add((time.Minute * 15)).Unix(),
	})
}

func (s *AuthService) NewAuthToken(id domain.Id, roles []string) (*string, error) {
	return s.jwtService.CreateToken(jwt.MapClaims{
		"id":    id,
		"roles": roles,
		"exp":   time.Now().Add((time.Hour * 24 * 30)).Unix(),
	})
}

func (s *AuthService) Login(ctx context.Context, email string, password string, oauth bool) (*string, error) {
	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	if oauth {
		if err := s.loginOAuth(ctx, validEmail); err != nil {
			return nil, err
		}
	} else {
		validPassword, err := domain.NewPassword(password)

		if err != nil {
			return nil, fmt.Errorf("invalid password supplied: %w", err)
		}

		if err := s.loginDefault(ctx, validEmail, validPassword); err != nil {
			return nil, err
		}
	}

	token, err := s.NewRefreshToken()

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrSigningToken
	}

	return token, nil
}

func (s *AuthService) loginDefault(ctx context.Context, email domain.Email, password domain.Password) error {
	user, err := s.authStore.Select(ctx, email)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	if isVerified := s.passwordService.VerifyPassword(password, user.Password); !isVerified || !user.Verified {
		return ErrIncorrectPassword
	}

	return nil
}

func (s *AuthService) loginOAuth(ctx context.Context, email domain.Email) error {
	return nil
}

func (s *AuthService) Register(ctx context.Context, email string, username string, password string, oauth bool) (*string, error) {

	validEmail, err := domain.NewEmail(email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	user, err := s.authStore.Select(ctx, validEmail)

	if err != nil && err != domain.ErrDataNotFound {
		slog.Error(err.Error())
		return nil, ErrUserAuthStore
	}

	if user != nil && user.Email == validEmail {
		return nil, fmt.Errorf("email address is already registered")
	}

	validUsername, err := domain.NewUsername(username)

	if err != nil {
		return nil, fmt.Errorf("invalid username supplied: %w", err)
	}

	if oauth {
		if err := s.registerOAuth(ctx, validEmail, validUsername); err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	} else {
		validPassword, err := domain.NewPassword(password)

		if err != nil {
			return nil, fmt.Errorf("invalid password supplied: %w", err)
		}

		if err := s.registerDefualt(ctx, validEmail, validUsername, validPassword); err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	token, err := s.NewRefreshToken()

	if err != nil {
		slog.Error(err.Error())
		return nil, ErrSigningToken
	}

	return token, nil
}

func (s *AuthService) registerOAuth(ctx context.Context, email domain.Email, username domain.Username) error {
	_, err := s.userStore.Insert(ctx, email, domain.Password(""), username, true)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	return nil
}

func (s *AuthService) registerDefualt(ctx context.Context, email domain.Email, username domain.Username, password domain.Password) error {
	hashedPassword := s.passwordService.HashPassword(password)

	_, err := s.userStore.Insert(ctx, email, domain.Password(hashedPassword), username, false)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	token, err := s.newVerificationToken(email)

	if err != nil {
		slog.Error(err.Error())
		return ErrSigningToken
	}

	err = s.authStore.UpdateVerifiedToken(ctx, email, *token)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserAuthStore
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8084/api/v1/auth/verify/%v", token)

	err = s.emailService.Send(
		[]string{"aaron.tibben@gmail.com"},
		"Subject:Do Not Reply - Finish Registration\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		s.emailService.VerificationTemplate(validationEndPoint),
	)

	if err != nil {
		slog.Error(err.Error())
		return ErrEmailService
	}

	return nil
}

func (s *AuthService) VerifyUser(ctx context.Context, token string) error {
	claims, err := s.jwtService.ParseToken(token)

	if err != nil {
		return ErrInvalidToken
	}

	email := claims["email"].(domain.Email)

	err = s.authStore.UpdateVerified(ctx, email)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	return nil
}

func (s *AuthService) UpdatePassword(ctx context.Context, token string, password string) error {
	claims, err := s.jwtService.ParseToken(token)

	if err != nil {
		return ErrInvalidToken
	}

	email := claims["email"].(domain.Email)

	newPassword, err := domain.NewPassword(password)

	if err != nil {
		return fmt.Errorf("invalid password supplied: %w", err)
	}

	hashedpas := s.passwordService.HashPassword(newPassword)

	err = s.authStore.UpdatePassword(ctx, email, hashedpas)

	if err != nil {
		return ErrUserStore
	}

	return nil
}

func (s *AuthService) ResetPassword(ctx context.Context, email domain.Email) error {
	token, err := s.newPasswordResetToken(email)

	if err != nil {
		slog.Error(err.Error())
		return ErrSigningToken
	}

	err = s.authStore.UpdateResetPasswordToken(ctx, email, *token)

	if err != nil {
		slog.Error(err.Error())
		return ErrUserStore
	}

	validationEndPoint := fmt.Sprintf("http://localhost:8084/api/v1/auth/reset/%v", *token)

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

	return nil
}
