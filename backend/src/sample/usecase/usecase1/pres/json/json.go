package json

import (
	"net/http"

	"src/module/account/domain/model"

	"github.com/labstack/echo/v4"
)

func RefreshTokenPres(c echo.Context, tokenPair model.TokenPair) error {
	data := map[string]interface{}{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	}

	return c.JSON(http.StatusOK, data)
}
