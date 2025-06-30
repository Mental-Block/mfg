package authentication

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type RefreshRequest struct{
	RefreshCookie http.Cookie `cookie:"mfg-refresh-token"`
	AuthCookie 	  http.Cookie `cookie:"mfg-authorization-token"`
}

type RefreshResponse struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
}

func (s *AuthHandler) Refresh(api huma.API) {
	huma.Register(api, huma.Operation{
		Tags:          []string{"authentication"},
		OperationID:   "refresh-token",
		Summary:       "Refresh token",
		Description:   "Gets a new auth token, removes the old auth token, refresh token",
		Path:          "/refresh/",
		Method:        http.MethodGet,
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {

		// refreshToken, authToken, err := s.authService.VerifyRefreshToken(ctx, req.RefreshCookie.Value, req.AuthCookie.Value)

		// if err != nil  {
		// 	return nil, util.HumaError(err)
		// }		

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
		// authCookie.Value = authToken;
		// authCookie.Expires = time.Now().Add(domain.PermissionTokenDuration)

		// resp := &RefreshResponse{}

		// resp.SetCookie = []*http.Cookie{
		// 	&refreshCookie,	
		// 	&authCookie,
		// }

		return nil, nil
	})
}
