package authentication

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/config/env"
	"github.com/server/internal/core/auth"
)

type AuthHandler struct {	
	host 	   string
	enviroment env.ENVIROMENT
	authService auth.AuthService 
}

func NewAuthHandler(
	host 	   string, 
	enviroment env.ENVIROMENT,
	auth auth.AuthService, 
	) *AuthHandler {
	return &AuthHandler{
		enviroment: enviroment,
		authService: auth,
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
