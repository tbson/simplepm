package setting

import (
	"fmt"
	"src/util/fwutil"
)

type authTokenSetting struct {
	AccessTokenLifetime  int
	AccessTokenSecret    string
	RefreshTokenLifetime int
	RefreshTokenSecret   string
}

var debug bool = fwutil.BoolEnv("DEBUG", true)

func DEBUG() bool {
	return debug
}

var baseUrl string = fwutil.StrEnv("BASE_URL", "")

func BASE_URL() string {
	return baseUrl
}

var domain string = fwutil.StrEnv("DOMAIN", "")

func DOMAIN() string {
	return domain
}

var appName string = fwutil.StrEnv("APP_NAME", "")

func APP_NAME() string {
	return appName
}

func COOKIE_LIFE_TIME_MINS() int {
	return fwutil.IntEnv("COOKIE_LIFE_TIME_MINS", 60)
}

var accessTokenLifetime int = fwutil.IntEnv("ACCESS_TOKEN_LIFETIME", 15)

func ACCESS_TOKEN_LIFETIME() int {
	return accessTokenLifetime
}

var accessTokenSecret string = fwutil.StrEnv("ACCESS_TOKEN_SECRET", "")

func ACCESS_TOKEN_SECRET() string {
	return accessTokenSecret
}

var refreshTokenLifetime int = fwutil.IntEnv("REFRESH_TOKEN_LIFETIME", 1440)

func REFRESH_TOKEN_LIFETIME() int {
	return refreshTokenLifetime
}

var refreshTokenSecret string = fwutil.StrEnv("REFRESH_TOKEN_SECRET", "")

func REFRESH_TOKEN_SECRET() string {
	return refreshTokenSecret
}

func AUTH_TOKEN_SETTINGS() authTokenSetting {
	return authTokenSetting{
		AccessTokenLifetime:  ACCESS_TOKEN_LIFETIME(),
		AccessTokenSecret:    ACCESS_TOKEN_SECRET(),
		RefreshTokenLifetime: REFRESH_TOKEN_LIFETIME(),
		RefreshTokenSecret:   REFRESH_TOKEN_SECRET(),
	}
}

var dbHost string = fwutil.StrEnv("DB_HOST", "")

func DB_HOST() string {
	return dbHost
}

var dbPort string = fwutil.StrEnv("DB_PORT", "5432")

func DB_PORT() string {
	return dbPort
}

var dbUser string = fwutil.StrEnv("DB_USER", "")

func DB_USER() string {
	return dbUser
}

var dbName string = fwutil.StrEnv("DB_NAME", "")

func DB_NAME() string {
	return dbName
}

var dbPassword string = fwutil.StrEnv("DB_PASSWORD", "")

func DB_PASSWORD() string {
	return dbPassword
}

var nosqlHost string = fwutil.StrEnv("NOSQL_HOST", "")

func NOSQL_HOST() string {
	return nosqlHost
}

var nosqlPort string = fwutil.StrEnv("NOSQL_PORT", "9042")

func NOSQL_PORT() string {
	return nosqlPort
}

var emailFrom string = fwutil.StrEnv("EMAIL_FROM", "")

func EMAIL_FROM() string {
	return emailFrom
}

func DEFAULT_EMAIL_FROM() string {
	return fmt.Sprintf("%s <%s>", APP_NAME(), EMAIL_FROM())
}

var emailHost string = fwutil.StrEnv("EMAIL_HOST", "")

func EMAIL_HOST() string {
	return emailHost
}

var emailPort int = fwutil.IntEnv("EMAIL_PORT", 587)

func EMAIL_PORT() int {
	return emailPort
}

var emailHostUser string = fwutil.StrEnv("EMAIL_HOST_USER", "")

func EMAIL_HOST_USER() string {
	return emailHostUser
}

var emailHostPassword string = fwutil.StrEnv("EMAIL_HOST_PASSWORD", "")

func EMAIL_HOST_PASSWORD() string {
	return emailHostPassword
}

var emailUseTls bool = fwutil.BoolEnv("EMAIL_USE_TLS", true)

func EMAIL_USE_TLS() bool {
	return emailUseTls
}

const DEFAULT_LANG = "en"

var timeZone string = fwutil.StrEnv("TIME_ZONE", "")

func TIME_ZONE() string {
	return timeZone
}

var defaultAdminEmail string = fwutil.StrEnv("DEFAULT_ADMIN_EMAIL", "admin@local.dev")

func DEFAULT_ADMIN_EMAIL() string {
	return defaultAdminEmail
}

var adminTeantUid string = fwutil.StrEnv("ADMIN_TEANT_UID", "admin")

func ADMIN_TEANT_UID() string {
	return adminTeantUid
}

var adminTeantTitle string = fwutil.StrEnv("ADMIN_TEANT_TITLE", "admin")

func ADMIN_TEANT_TITLE() string {
	return adminTeantTitle
}

var testTeantUid string = fwutil.StrEnv("TEST_TEANT_UID", "")

func TEST_TEANT_UID() string {
	return testTeantUid
}

var tetsTeantTitle string = fwutil.StrEnv("TETS_TEANT_TITLE", "")

func TETS_TEANT_TITLE() string {
	return tetsTeantTitle
}

var testUserEmailAdmin string = fwutil.StrEnv("TEST_USER_EMAIL_ADMIN", "")

func TEST_USER_EMAIL_ADMIN() string {
	return testUserEmailAdmin
}

var testUserEmailStaff string = fwutil.StrEnv("TEST_USER_EMAIL_STAFF", "")

func TEST_USER_EMAIL_STAFF() string {
	return testUserEmailStaff
}

var testUserEmailOwner string = fwutil.StrEnv("TEST_USER_EMAIL_OWNER", "")

func TEST_USER_EMAIL_OWNER() string {
	return testUserEmailOwner
}

var testUserEmailManager string = fwutil.StrEnv("TEST_USER_EMAIL_MANAGER", "")

func TEST_USER_EMAIL_MANAGER() string {
	return testUserEmailManager
}

var testUserEmailUser string = fwutil.StrEnv("TEST_USER_EMAIL_USER", "")

func TEST_USER_EMAIL_USER() string {
	return testUserEmailUser
}

var testUserPassword string = fwutil.StrEnv("TEST_USER_PASSWORD", "")

func TEST_USER_PASSWORD() string {
	return testUserPassword
}

var s3AccountId string = fwutil.StrEnv("S3_ACCOUNT_ID", "")

func S3_ACCOUNT_ID() string {
	return s3AccountId
}

var s3AccessKeyId string = fwutil.StrEnv("S3_ACCESS_KEY_ID", "")

func S3_ACCESS_KEY_ID() string {
	return s3AccessKeyId
}

var s3SecretAccessKey string = fwutil.StrEnv("S3_SECRET_ACCESS_KEY", "")

func S3_SECRET_ACCESS_KEY() string {
	return s3SecretAccessKey
}

var s3BucketName string = fwutil.StrEnv("S3_BUCKET_NAME", "")

func S3_BUCKET_NAME() string {
	return s3BucketName
}

var s3Region string = fwutil.StrEnv("S3_REGION", "")

func S3_REGION() string {
	return s3Region
}

var s3EndpointUrl string = fwutil.StrEnv("S3_ENDPOINT_URL", "")

func S3_ENDPOINT_URL() string {
	return s3EndpointUrl
}

var sentryDsn string = fwutil.StrEnv("SENTRY_DSN", "")

func SENTRY_DSN() string {
	return sentryDsn
}

func FE_REDIRECT_URI() string {
	return fmt.Sprintf("%s%s", BASE_URL(), "/login")
}

var centrifugoClientSecret string = fwutil.StrEnv("CENTRIFUGO_CLIENT_SECRET", "")

func CENTRIFUGO_CLIENT_SECRET() string {
	return centrifugoClientSecret
}

var centrifugoApiKey string = fwutil.StrEnv("CENTRIFUGO_API_KEY", "")

func CENTRIFUGO_API_KEY() string {
	return centrifugoApiKey
}

var centrifugoApiEndpoint string = fwutil.StrEnv("CENTRIFUGO_API_ENDPOINT", "")

func CENTRIFUGO_API_ENDPOINT() string {
	return centrifugoApiEndpoint
}

var centrifugoJwtLifeSpan int = fwutil.IntEnv("CENTRIFUGO_JWT_LIFE_SPAN", 1200)

func CENTRIFUGO_JWT_LIFE_SPAN() int {
	return centrifugoJwtLifeSpan
}

var rabbitmqHost string = fwutil.StrEnv("RABBITMQ_HOST", "localhost")

func RABBITMQ_HOST() string {
	return rabbitmqHost
}

var rabbitmqPort int = fwutil.IntEnv("RABBITMQ_PORT", 9092)

func RABBITMQ_PORT() int {
	return rabbitmqPort
}

var rabbitmqUser string = fwutil.StrEnv("RABBITMQ_USER", "guest")

func RABBITMQ_USER() string {
	return rabbitmqUser
}

var rabbitmqPassword string = fwutil.StrEnv("RABBITMQ_PASSWORD", "guest")

func RABBITMQ_PASSWORD() string {
	return rabbitmqPassword
}

var githubAppPublicLink string = fwutil.StrEnv("GITHUB_APP_PUBLIC_LINK", "")

func GITHUB_APP_PUBLIC_LINK() string {
	return githubAppPublicLink
}

var githubClientId string = fwutil.StrEnv("GITHUB_CLIENT_ID", "")

func GITHUB_CLIENT_ID() string {
	return githubClientId
}

var githubClientSecret string = fwutil.StrEnv("GITHUB_CLIENT_SECRET", "")

func GITHUB_CLIENT_SECRET() string {
	return githubClientSecret
}

var githubPrivateKey string = fwutil.StrEnv("GITHUB_PRIVATE_KEY", "")

func GITHUB_PRIVATE_KEY() string {
	return githubPrivateKey
}

func GITHUB_PRIVATE_KEY_PATH() string {
	return "config/" + GITHUB_PRIVATE_KEY()
}

var queueBackend string = fwutil.StrEnv("QUEUE_BACKEND", "rabbitmq")

func QUEUE_BACKEND() string {
	return queueBackend
}

var msgPageSize int = fwutil.IntEnv("MSG_PAGE_SIZE", 25)

func MSG_PAGE_SIZE() int {
	return msgPageSize
}

var maxResetPwdPeriodDays int = fwutil.IntEnv("MAX_RESET_PWD_PERIOD_DAYS", 90)

func MAX_RESET_PWD_PERIOD_DAYS() int {
	return maxResetPwdPeriodDays
}

var otpLength int = fwutil.IntEnv("OTP_LENGTH", 6)

func OTP_LENGTH() int {
	return otpLength
}

var otpLifeMinutes int = fwutil.IntEnv("OTP_LIFE_MINUTES", 10)

func OTP_LIFE_MINUTES() int {
	return otpLifeMinutes
}

var pwdFailedAttemps int = fwutil.IntEnv("MAX_PWD_FAILED_ATTEMPTS", 5)

func MAX_PWD_FAILED_ATTEMPTS() int {
	return pwdFailedAttemps
}

var lastPwdsCheck int = fwutil.IntEnv("LAST_PWDS_CHECK", 5)

func LAST_PWDS_CHECK() int {
	return lastPwdsCheck
}
