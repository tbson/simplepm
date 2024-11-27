package infra

import (
	"net/http"

	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/user"
	"src/module/account/usecase/auth/app"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/ssoutil"

	"github.com/labstack/echo/v4"
)

func getService() app.Service {
	client := dbutil.Db()
	userRepo := user.New(client)
	iamRepo := iam.New(ssoutil.Client())
	authRepo := New(client)
	return app.New(userRepo, iamRepo, authRepo)
}

func CheckAuthUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")

	srv := getService()

	_, err := srv.GetAuthUrl(tenantUid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func GetAuthUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")

	srv := getService()

	url, err := srv.GetAuthUrl(tenantUid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GetLogoutUrl(c echo.Context) error {
	tenantUid := c.Param("tenantUid")

	srv := getService()

	url, err := srv.GetLogoutUrl(tenantUid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	if code == "" || state == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	srv := getService()

	result, err := srv.HandleCallback(c.Request().Context(), state, code)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/auth-error")
	}
	return CallbackPres(c, result)
}

func PostLogout(c echo.Context) error {
	data := map[string]interface{}{}
	return c.Render(http.StatusOK, "post-logout.html", data)
}

func RefreshToken(c echo.Context) error {
	refreshToken := cookieutil.GetValue(c, "refresh_token")
	realm := cookieutil.GetValue(c, "realm")

	srv := getService()

	result, err := srv.RefreshToken(c.Request().Context(), realm, refreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return RefreshTokenPres(c, result)
}

func RefreshTokenCheck(c echo.Context) error {
	result := ctype.Dict{}
	return c.JSON(http.StatusOK, result)
}
