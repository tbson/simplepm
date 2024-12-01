package infra

import (
	"encoding/json"
	"net/http"
	"src/common/authtype"
	"src/common/ctype"
	"src/util/cookieutil"
	"src/util/dbutil"

	"github.com/labstack/echo/v4"
)

func CallbackPres(c echo.Context, r authtype.AuthCallbackResult) error {
	authRepo := New(dbutil.Db())

	accessTokenCookie := cookieutil.NewAccessTokenCookie(r.AccessToken)
	realmCookie := cookieutil.NewRealmCookie(r.Realm)
	refreshTokenCookie := cookieutil.NewRefreshTokenCookie(r.RefreshToken)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(realmCookie)
	c.SetCookie(refreshTokenCookie)

	userInfo := r.UserInfo

	pemModulesActionsMap, err := authRepo.GetPemModulesActionsMap(userInfo.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	auth := ctype.Dict{
		"userInfo":             userInfo,
		"pemModulesActionsMap": pemModulesActionsMap,
	}
	authJson, _ := json.Marshal(auth)

	data := map[string]interface{}{
		"auth": string(authJson),
	}

	return c.Render(http.StatusOK, "post-login.html", data)
}

func RefreshTokenPres(c echo.Context, r authtype.SsoCallbackResult) error {
	accessTokenCookie := cookieutil.NewAccessTokenCookie(r.AccessToken)
	realmCookie := cookieutil.NewRealmCookie(r.Realm)
	refreshTokenCookie := cookieutil.NewRefreshTokenCookie(r.RefreshToken)
	c.SetCookie(accessTokenCookie)
	c.SetCookie(realmCookie)
	c.SetCookie(refreshTokenCookie)

	return c.JSON(http.StatusOK, ctype.Dict{})
}
