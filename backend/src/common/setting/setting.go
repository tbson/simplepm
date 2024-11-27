package setting

import (
	"fmt"
	"src/util/frameworkutil"
)

var DEBUG bool = frameworkutil.GetEnv("DEBUG", "true") == "true"
var BASE_URL string = frameworkutil.GetEnv("BASE_URL", "")
var DOMAIN string = frameworkutil.GetEnv("DOMAIN", "")
var DB_HOST string = frameworkutil.GetEnv("DB_HOST", "")
var DB_PORT string = frameworkutil.GetEnv("DB_PORT", "5432")
var DB_USER string = frameworkutil.GetEnv("DB_USER", "")
var DB_NAME string = frameworkutil.GetEnv("DB_NAME", "")
var DB_PASSWORD string = frameworkutil.GetEnv("DB_PASSWORD", "")

const DEFAULT_LANG = "en"

var TIME_ZONE string = frameworkutil.GetEnv("TIME_ZONE", "")
var DEFAULT_ADMIN_EMAIL string = frameworkutil.GetEnv("DEFAULT_ADMIN_EMAIL", "admin@local.dev")
var ADMIN_TEANT_UID string = frameworkutil.GetEnv("ADMIN_TEANT_UID", "admin")
var ADMIN_TEANT_TITLE string = frameworkutil.GetEnv("ADMIN_TEANT_TITLE", "admin")

var S3_ACCOUNT_ID string = frameworkutil.GetEnv("S3_ACCOUNT_ID", "")
var S3_ACCESS_KEY_ID string = frameworkutil.GetEnv("S3_ACCESS_KEY_ID", "")
var S3_SECRET_ACCESS_KEY string = frameworkutil.GetEnv("S3_SECRET_ACCESS_KEY", "")
var S3_BUCKET_NAME string = frameworkutil.GetEnv("S3_BUCKET_NAME", "")
var S3_REGION string = frameworkutil.GetEnv("S3_REGION", "")
var S3_ENDPOINT_URL string = frameworkutil.GetEnv("S3_ENDPOINT_URL", "")

var SENTRY_DSN string = frameworkutil.GetEnv("SENTRY_DSN", "")

var KEYCLOAK_ADMIN string = frameworkutil.GetEnv("KEYCLOAK_ADMIN", "")
var KEYCLOAK_ADMIN_PASSWORD string = frameworkutil.GetEnv("KEYCLOAK_ADMIN_PASSWORD", "")
var KEYCLOAK_URL string = frameworkutil.GetEnv("KEYCLOAK_URL", "")
var KEYCLOAK_DEFAULT_REALM string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_REALM", "")
var KEYCLOAK_DEFAULT_CLIENT_ID string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_ID", "")
var KEYCLOAK_DEFAULT_CLIENT_SECRET string = frameworkutil.GetEnv("KEYCLOAK_DEFAULT_CLIENT_SECRET", "")
var KEYCLOAK_REDIRECT_URI string = fmt.Sprintf(
	"%s%s",
	BASE_URL,
	frameworkutil.GetEnv("KEYCLOAK_REDIRECT_URI", ""),
)
var KEYCLOAK_POST_LOGOUT_URI string = fmt.Sprintf(
	"%s%s",
	BASE_URL,
	frameworkutil.GetEnv("KEYCLOAK_POST_LOGOUT_URI", ""),
)

var FE_REDIRECT_URI string = fmt.Sprintf(
	"%s%s",
	BASE_URL,
	"/login",
)
