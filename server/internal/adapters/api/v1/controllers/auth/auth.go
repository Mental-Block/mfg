package auth

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/core/ports"
)

type ServiceInject struct {
	authService ports.AuthService
	parrentGrp  *huma.Group
}

func NewServiceInject(authService ports.AuthService, grp *huma.Group) *ServiceInject {
	return &ServiceInject{
		authService: authService,
		parrentGrp:  grp,
	}
}

func (s *ServiceInject) Routes() {
	authGrp := huma.NewGroup(s.parrentGrp, "/auth")

	s.AuthToken(authGrp)
	s.Login(authGrp)
	s.Logout(authGrp)
	//s.RefreshToken(authGrp)
	s.Register(authGrp)
	s.ResetPassword(authGrp)
	s.UpdatePassword(authGrp)
	s.VerifyEmail(authGrp)

}
