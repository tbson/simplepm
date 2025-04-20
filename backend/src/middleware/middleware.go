package middleware

import (
	"context"
	"src/common/ctype"
	"src/module/account/domain/srv/authtoken"
	"src/module/account/repo/tenant"
	"src/module/account/repo/user"
	"src/module/account/schema"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/numberutil"
	"strings"

	"src/common/setting"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func BlankMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		localizer := localeutil.Get()
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MissingTenantID,
		})

		requestDomain := c.Request().Host
		domainParts := strings.Split(requestDomain, ".")
		if len(domainParts) < 2 {
			return c.JSON(400, errutil.New("", []string{msg}))
		}
		tenantUID := domainParts[0]
		tenantRepo := tenant.New(dbutil.Db(nil))
		tenant, err := tenantRepo.Retrieve(ctype.QueryOpts{
			Filters: ctype.Dict{"Uid": tenantUID},
		})
		if err != nil {
			return c.JSON(400, errutil.New("", []string{msg}))
		}
		c.Set("TenantID", tenant.ID)
		c.Set("TenantUid", tenant.Uid)
		return next(c)
	}
}

func LangMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accept := c.Request().Header.Get("Accept-Language")
		localizer := localeutil.Init(accept)
		c.Set("localizer", localizer)
		c.Set("lang", accept)
		return next(c)
	}
}

func AuthMiddleware(module string, action string, isRbac bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenSettings := setting.AUTH_TOKEN_SETTINGS()

			localizer := localeutil.Get()
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.Unauthorized,
			})
			// check access_token cookie
			userRepo := user.New(dbutil.Db(nil))
			authTokenSrv := authtoken.New(
				tokenSettings.AccessTokenSecret,
				tokenSettings.RefreshTokenSecret,
				tokenSettings.AccessTokenLifetime,
				tokenSettings.RefreshTokenLifetime,
			)

			accessToken := cookieutil.GetValue(c, "access_token")

			if accessToken == "" {
				return c.JSON(401, errutil.New("", []string{msg}))
			}

			userID, err := authTokenSrv.VerifyAccessToken(accessToken)
			if err != nil {
				return c.JSON(401, err)
			}

			// preload roles and pems which pem
			user, err := userRepo.Retrieve(ctype.QueryOpts{
				Filters:  ctype.Dict{"id": userID},
				Preloads: []string{"Roles.Pems", "Tenant"},
			})
			if err != nil {
				return c.JSON(401, errutil.New("", []string{msg}))
			}

			if user.LockedAt != nil {
				msg := localizer.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: localeutil.LockedAccount,
				})
				return c.JSON(401, errutil.New("", []string{msg}))
			}

			// check cross tenant query
			// only admin can do it
			var tenantID uint = user.TenantID
			if user.Admin == true && user.TenantTmpID != nil {
				tenantID = *user.TenantTmpID
			}
			specificTenantIDStr := c.QueryParam("tenant_id")
			if specificTenantIDStr != "" {
				specificTenantID := numberutil.StrToUint(specificTenantIDStr, 0)
				if user.Admin == false && specificTenantID != tenantID {
					return c.JSON(401, errutil.New("", []string{msg}))
				}
				tenantID = specificTenantID
			}

			var setContext = func(c echo.Context, user *schema.User, tenantID uint) {
				c.Set("User", user)
				c.Set("UserID", user.ID)
				c.Set("Admin", user.Admin)
				c.Set("TenantID", tenantID)
				c.Set("TenantUid", user.Tenant.Uid)

				ctx := c.Request().Context()
				ctx = context.WithValue(ctx, "UserID", user.ID)
				ctx = context.WithValue(ctx, "TenantID", tenantID)
				c.SetRequest(c.Request().WithContext(ctx))
			}

			if !isRbac {
				setContext(c, user, tenantID)
				return next(c)
			}

			for _, role := range user.Roles {
				for _, pem := range role.Pems {
					if pem.Module == module && pem.Action == action {
						setContext(c, user, tenantID)
						return next(c)
					}
				}
			}

			// return next(c)
			return c.JSON(401, errutil.New("", []string{msg}))
		}
	}
}
