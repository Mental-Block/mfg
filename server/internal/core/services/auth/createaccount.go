package auth

import (
	"context"
	"fmt"

	"github.com/server/internal/core/domain/auth"
	"github.com/server/internal/core/domain/user"
	"github.com/server/internal/core/domain/userAuth"
	"github.com/server/internal/core/services/email"
)

type CreateAccountInput struct {
	Username string
	Password string
	Email    string
}

type CreateAccountOutput = user.UserEntity

func (s *Service) CreateAccount(ctx context.Context, account CreateAccountInput) (*CreateAccountOutput, error) {

	newUsername, err := user.NewUsername(account.Username)

	if err != nil {
		return nil, fmt.Errorf("invalid username supplied: %w", err)
	}

	newEmail, err := auth.NewEmail(account.Email)

	if err != nil {
		return nil, fmt.Errorf("invalid email supplied: %w", err)
	}

	newPassword, err := auth.NewPassword(account.Password)

	if err != nil {
		return nil, fmt.Errorf("invalid password supplied: %w", err)
	}

	isSent, err := email.SendEmail(
		"aaron.tibben@gmail.com",
		"Subject:Do Not Reply - Finish Registration\n",
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
		email.VerificationEmailTemplate("my api endpoint"),
	)

	if err != nil && !isSent {
		return nil, fmt.Errorf("failed to send email:%w", err)
	}

	// var once sync.Once
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	once.Do(func() {

	// 		time.Sleep(time.Minute * 5)
	// 	})
	// }()

	// wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	newUser, err := s.userStore.InsertUser(ctx,
		userAuth.New(newUsername, newEmail, newPassword))

	if err != nil {
		return nil, fmt.Errorf("failed to add a user: %w", err)
	}

	return newUser, nil
}
