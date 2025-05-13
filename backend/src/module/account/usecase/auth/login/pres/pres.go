package pres

import (
	"src/common/ctype"
	"src/module/account/usecase/auth/login/repo"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/presutil"

	"src/module/account/domain/model"

	"src/util/tokenutil"

	"github.com/labstack/echo/v4"
)

func Login(
	c echo.Context,
	loginResult model.AuthResult,
	next string,
	clientType string,
) error {
	if clientType == model.CLIENT_TYPE_WEB {
		return cookiePres(c, loginResult, next)
	}
	return jsonPres(c, loginResult, next)
}

func cookiePres(c echo.Context, loginResult model.AuthResult, next string) error {
	resp := presutil.New(c)
	authRepo := repo.New(dbutil.Db(nil))

	accessToken := loginResult.TokenPair.AccessToken
	refreshToken := loginResult.TokenPair.RefreshToken
	userInfo := loginResult.UserInfo
	sessionID := tokenutil.GenerateUUID()

	accessTokenCookie := cookieutil.NewAccessTokenCookie(accessToken)
	refreshTokenCookie := cookieutil.NewRefreshTokenCookie(refreshToken)
	sessionIDCookie := cookieutil.NewSessionIDCookie(sessionID)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(refreshTokenCookie)
	c.SetCookie(sessionIDCookie)

	pemModulesActionsMap, err := authRepo.GetPemModulesActionsMap(userInfo.ID)

	if err != nil {
		return resp.Err(err)
	}

	authInfo := ctype.Dict{
		"userInfo":             userInfo,
		"pemModulesActionsMap": pemModulesActionsMap,
	}

	data := map[string]interface{}{
		"auth": authInfo,
		"next": next,
	}

	// return c.Render(http.StatusOK, "post-login.html", data)
	return resp.Ok(data)
}

func jsonPres(c echo.Context, loginResult model.AuthResult, next string) error {
	resp := presutil.New(c)
	authRepo := repo.New(dbutil.Db(nil))

	accessToken := loginResult.TokenPair.AccessToken
	refreshToken := loginResult.TokenPair.RefreshToken
	userInfo := loginResult.UserInfo
	sessionID := tokenutil.GenerateUUID()

	pemModulesActionsMap, err := authRepo.GetPemModulesActionsMap(userInfo.ID)

	if err != nil {
		return resp.Err(err)
	}

	authInfo := ctype.Dict{
		"userInfo":             userInfo,
		"pemModulesActionsMap": pemModulesActionsMap,
	}

	data := map[string]interface{}{
		"session_id":    sessionID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"auth":          authInfo,
		"next":          next,
	}

	// return c.Render(http.StatusOK, "post-login.html", data)
	return resp.Ok(data)
}
