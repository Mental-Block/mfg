package domain

import "time"


var (
	RefreshTokenName = "mfg-refresh-token"
	AuthTokenName = "mfg-authourization"
)

var (
	PasswordResetTokenDuration = time.Minute * 5
	RefreshTokenDuration = time.Hour * 24 * 30
	AuthTokenDuration = time.Minute * 15
	EmailVerificationToken = time.Minute * 15
)

type AuthUser struct {
	Username Username
	Email    Email
	Password Password
	Token    string
}


