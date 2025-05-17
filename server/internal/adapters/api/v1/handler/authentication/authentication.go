package authentication

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type AuthHandler struct {
	authService ports.AuthService 
	userService ports.UserService
}

func NewAuthHandler(auth ports.AuthService, user ports.UserService) *AuthHandler {
	return &AuthHandler{
		authService: auth,
		userService: user,
	}
}

func (s *AuthHandler) Routes(parrentGrp *huma.Group) {
	authGrp := huma.NewGroup(parrentGrp, "/auth")

	s.Login(authGrp)
	s.Logout(authGrp)
	s.Refresh(authGrp)
	s.Register(authGrp)
	s.RegisterFinish(authGrp)
	s.ResetPassword(authGrp)
	s.UpdatePassword(authGrp)
	s.Verify(authGrp)
	s.EmailTaken(authGrp)
}
