package cookieutil

import (
	"net/http"
	"src/common/setting"
	"strings"

	"github.com/labstack/echo/v4"
)

func newCookie(name string, value string, path string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = setting.DOMAIN         // Set the Domain attribute
	cookie.Path = path                     // Set the Path attribute
	cookie.Secure = true                   // Set the Secure attribute
	cookie.HttpOnly = true                 // Prevents JavaScript access (optional)
	cookie.SameSite = http.SameSiteLaxMode // Set the SameSite attribute
	return cookie
}

func NewAccessTokenCookie(value string) *http.Cookie {
	return newCookie("access_token", value, "/api/v1/")
}

func NewRealmCookie(value string) *http.Cookie {
	return newCookie("realm", value, "/api/v1/")
}

func NewRefreshTokenCookie(value string) *http.Cookie {
	return newCookie("refresh_token", value, "/api/v1/account/auth/sso/refresh-token")
}

func GetValue(c echo.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err == nil {
		return cookie.Value
	}

	if name == "access_token" {
		name = "Authorization"
	}

	header := c.Request().Header.Get(name)
	if header != "" {
		if name == "Authorization" {
			return strings.Split(header, " ")[1]
		}
		return header
	}

	return ""
}
