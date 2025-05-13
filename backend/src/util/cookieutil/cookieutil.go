package cookieutil

import (
	"net/http"
	"src/common/setting"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func newCookie(name string, value string, path string, expiredMins int) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = setting.DOMAIN()
	cookie.Path = path
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode

	if expiredMins > 0 {
		expiration := time.Now().Add(time.Duration(expiredMins) * time.Minute)
		cookie.Expires = expiration
	} else if expiredMins == 0 {
		cookie.MaxAge = -1
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		cookie.Value = ""
	}
	return cookie
}

func SetAccessTokenCookie(value string) *http.Cookie {
	ttl := setting.COOKIE_LIFE_TIME_MINS()
	if value == "" {
		ttl = 0
	}
	return newCookie("access_token", value, "/api/v1/", ttl)
}

func SetRefreshTokenCookie(value string) *http.Cookie {
	ttl := setting.COOKIE_LIFE_TIME_MINS()
	if value == "" {
		ttl = 0
	}
	return newCookie("refresh_token", value, "/api/v1/account/auth/refresh-token", ttl)
}

func SetSessionIDCookie(value string) *http.Cookie {
	ttl := setting.COOKIE_LIFE_TIME_MINS()
	if value == "" {
		ttl = 0
	}
	return newCookie("session_id", value, "/api/v1/socket/jwt/subscription/", ttl)
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
