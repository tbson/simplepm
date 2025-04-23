package cookie

import (
	"net/http"
	"src/common/ctype"
	"src/module/account/usecase/auth/login/repo"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/errutil"

	"src/module/account/domain/model"

	"src/util/tokenutil"

	"github.com/labstack/echo/v4"
)

func LoginPres(c echo.Context, loginResult model.LoginResult, next string) error {
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
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
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
	return c.JSON(http.StatusOK, data)
}

func RefreshTokenPres(c echo.Context, accessToken string, refreshToken string) error {
	accessTokenCookie := cookieutil.NewAccessTokenCookie(accessToken)
	refreshTokenCookie := cookieutil.NewRefreshTokenCookie(refreshToken)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(refreshTokenCookie)

	return c.JSON(http.StatusOK, ctype.Dict{})
}
