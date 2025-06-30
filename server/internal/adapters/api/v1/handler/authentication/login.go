package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal/adapters/api/v1/handler/user"
)

type LoginRequest struct {
	Body struct {
		Email    string `example:"bob@gmail.com" doc:"unique email to each account"`
		Password string `example:"MyNewPassword123!" doc:"account login password"`
	}
}

type LoginResponse struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
	Body *user.User
}

func (s *AuthHandler) Login(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "login-account",
		Summary:       "Login account",
		Description:   "validates user credentials, returning a refresh token",
		Path:          "/login/",
		Method:        http.MethodPost,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {

		// startFlowReq := &domain.StartFlowReq{
		// 	Reason: "login",
		// 	Strategy: "password",
		// 	Email:  req.Body.Email,
		// 	Payload: map[string]string {
		// 		"password": req.Body.Password,
		// 	},
		// 	CallBackUrl: "",
		// 	ReturnToURL: "",
		// }

		// startflowResp, err := s.authService.StartFlow(ctx, startFlowReq)

		// if (err != nil) {
		// 	return nil, util.HumaError(err)
		// }

		// finishFlowReq := &domain.FinishFlowReq{
		// 	State: startflowResp.State,
		// 	Strategy: startflowResp.Flow.Strategy,
		// 	Reason: startflowResp.Flow.Reason,
		// }

		// _, err := s.authService.FinishFlow(ctx, finishFlowReq)

		// if (err != nil) {
		// 	return nil, util.HumaError(err)
		// }

			


		//s.authService.BuildToken(ctx, )


		// cookie := http.Cookie{
		// 	Secure: false,	
		// 	HttpOnly: false,
		// 	SameSite: http.SameSiteLaxMode,
		// 	Path: "/",
		// }

		// if s.enviroment == env.Production || s.enviroment == env.Test {
		// 	cookie.Domain = s.host
		// 	cookie.HttpOnly = true
		// 	cookie.Secure = true
		// 	cookie.SameSite = http.SameSiteStrictMode
		// }

		// refreshCookie := cookie
		// authCookie := cookie

		// refreshCookie.Name = domain.RefreshTokenName;
		// refreshCookie.Value = refreshToken;
		// refreshCookie.Expires = time.Now().Add(domain.RefreshTokenDuration)

		// authCookie.Name = domain.AuthTokenName;
		// authCookie.Value =  authToken;
		// authCookie.Expires = time.Now().Add(domain.PermissionTokenDuration)

		// resp := &LoginResponse{}

		// resp.SetCookie = []*http.Cookie{
		// 	&refreshCookie,	
		// 	&authCookie,
		// }

		// resp.Body = &user.User{
		// 	Id: loginUser.Id,
		// 	Username: loginUser.Username,
		// 	Title: loginUser.Title,
		// 	Active: loginUser.Active,
		// 	Avatar: loginUser.Avatar,
		// 	Metadata: loginUser.Metadata,
		// }

		return nil, nil
	})
}
