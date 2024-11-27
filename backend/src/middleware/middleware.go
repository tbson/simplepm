package middleware

import (
	"src/common/ctype"
	"src/module/account/repo/iam"
	"src/module/account/repo/user"
	"src/util/cookieutil"
	"src/util/dbutil"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/numberutil"
	"src/util/ssoutil"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func BlankMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
			localizer := localeutil.Get()
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: localeutil.Unauthorized,
			})
			// check access_token cookie
			iamRepo := iam.New(ssoutil.Client())
			userRepo := user.New(dbutil.Db())

			accessToken := cookieutil.GetValue(c, "access_token")
			realm := cookieutil.GetValue(c, "realm")

			if accessToken == "" || realm == "" {
				return c.JSON(401, errutil.New("", []string{msg}))
			}

			userInfo, err := iamRepo.ValidateToken(accessToken, realm)
			if err != nil {
				return c.JSON(401, err)
			}

			email := userInfo.Email
			// preload roles and pems which pem
			user, err := userRepo.Retrieve(ctype.QueryOptions{
				Filters:  ctype.Dict{"email": email},
				Preloads: []string{"Roles.Pems"},
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

			if !isRbac {
				c.Set("UserID", user.ID)
				c.Set("Admin", user.Admin)
				c.Set("TenantID", tenantID)
				return next(c)
			}

			for _, role := range user.Roles {
				for _, pem := range role.Pems {
					if pem.Module == module && pem.Action == action {
						c.Set("UserID", user.ID)
						c.Set("Admin", user.Admin)
						c.Set("TenantID", tenantID)
						return next(c)
					}
				}
			}

			// return next(c)
			return c.JSON(401, errutil.New("", []string{msg}))
		}
	}
}
