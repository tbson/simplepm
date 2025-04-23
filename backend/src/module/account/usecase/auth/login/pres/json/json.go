package json

import (
	"net/http"
	"src/common/ctype"
	"src/module/account/usecase/auth/login/repo"
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

	pemModulesActionsMap, err := authRepo.GetPemModulesActionsMap(userInfo.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.(*errutil.CustomError).Localize())
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
	return c.JSON(http.StatusOK, data)
}

func RefreshTokenPres(c echo.Context, accessToken string, refreshToken string) error {
	data := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(http.StatusOK, data)
}
