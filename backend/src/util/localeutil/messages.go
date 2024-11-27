package localeutil

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	CannotReadRequestBody = &i18n.Message{
		ID:    "CannotReadRequestBody",
		Other: "can not read request body",
	}

	InvalidJSONPayload = &i18n.Message{
		ID:    "InvalidJSONPayload",
		Other: "invalid JSON payload",
	}

	FieldRequired = &i18n.Message{
		ID:    "FieldRequired",
		Other: "this field is required",
	}

	BadJson = &i18n.Message{
		ID:    "BadJson",
		Other: "request body contains badly-formed JSON (at position {{.Offset}})",
	}

	InvalidFieldValue = &i18n.Message{
		ID:    "InvalidFieldValue",
		Other: "invalid value for field '{{.FieldName}}'. Expected type {{.Type}} but got {{.Value}}",
	}

	InvalidValue = &i18n.Message{
		ID:    "InvalidValue",
		Other: "invalid value",
	}

	NotRecognizedField = &i18n.Message{
		ID:    "NotRecognizedField",
		Other: "this field is not recognized",
	}

	FailToDecodeJSON = &i18n.Message{
		ID:    "FailToDecodeJSON",
		Other: "failed to decode JSON",
	}

	EmptyRequestBody = &i18n.Message{
		ID:    "EmptyRequestBody",
		Other: "request body must not be empty",
	}

	MustBeOneOf = &i18n.Message{
		ID:    "MustBeOneOf",
		Other: "must be one of: {{.Values}}",
	}

	GormDuplicateKey = &i18n.Message{
		ID:    "GormDuplicateKey",
		Other: "value already exists",
	}

	InvalidState = &i18n.Message{
		ID:    "InvalidState",
		Other: "invalid state",
	}

	AuthorizationCodeNotFound = &i18n.Message{
		ID:    "AuthorizationCodeNotFound",
		Other: "authorization code not found",
	}

	CannotExchangeAuthorizationCode = &i18n.Message{
		ID:    "CannotExchangeAuthorizationCode",
		Other: "can not exchange authorization code for tokens",
	}

	FailedToFetchJWKS = &i18n.Message{
		ID:    "FailedToFetchJWKS",
		Other: "failed to fetch JWKS",
	}

	FailedToParseToken = &i18n.Message{
		ID:    "FailedToParseToken",
		Other: "failed to parse token",
	}

	NoKidFieldInJWTTokenHeader = &i18n.Message{
		ID:    "NoKidFieldInJWTTokenHeader",
		Other: "no 'kid' field in JWT token header",
	}

	UnableToFindKeyWithKid = &i18n.Message{
		ID:    "UnableToFindKeyWithKid",
		Other: "unable to find key with kid",
	}

	FailedToCreateRawKey = &i18n.Message{
		ID:    "FailedToCreateRawKey",
		Other: "failed to create raw key",
	}

	ExpectedRSAKey = &i18n.Message{
		ID:    "ExpectedRSAKey",
		Other: "expected RSA public key but got something else",
	}

	FailedToVerifyToken = &i18n.Message{
		ID:    "FailedToVerifyToken",
		Other: "failed to verify token",
	}

	TokenHasExpired = &i18n.Message{
		ID:    "TokenHasExpired",
		Other: "token has expired",
	}

	NoRealmFound = &i18n.Message{
		ID:    "NoRealmFound",
		Other: "no realm found",
	}

	RefreshTokenNotFound = &i18n.Message{
		ID:    "RefreshTokenNotFound",
		Other: "fefresh token not found",
	}

	CannotExchangeRefreshToken = &i18n.Message{
		ID:    "CannotExchangeRefreshToken",
		Other: "can not exchange refresh token for tokens",
	}

	Unauthorized = &i18n.Message{
		ID:    "Unauthorized",
		Other: "unauthorized",
	}

	CannotGetUserInfo = &i18n.Message{
		ID:    "CannotGetUserInfo",
		Other: "cannot get user info",
	}

	CannotLoginAdmin = &i18n.Message{
		ID:    "cannotLoginAdmin",
		Other: "Cannot login admin",
	}

	CannotUpdateIAMUser = &i18n.Message{
		ID:    "CannotUpdateIAMUser",
		Other: "cannot update IAM user",
	}

	CannotSetPassword = &i18n.Message{
		ID:    "CannotSetPassword",
		Other: "cannot set password",
	}

	PasswordsNotMatch = &i18n.Message{
		ID:    "PasswordsNotMatch",
		Other: "passwords do not match",
	}

	FailedToOpenFile = &i18n.Message{
		ID:    "FailedToOpenFile",
		Other: "failed to open file '{{.Filename}}'",
	}

	FailedToUploadFileToS3 = &i18n.Message{
		ID:    "FailedToUploadFileToS3",
		Other: "failed to upload file to S3",
	}

	MissingTenantID = &i18n.Message{
		ID:    "MissingTenantID",
		Other: "missing tenant ID",
	}

	SubClaimNotFound = &i18n.Message{
		ID:    "SubClaimNotFound",
		Other: "sub claim not found",
	}

	LockedAccount = &i18n.Message{
		ID:    "LockedAccount",
		Other: "account is locked",
	}

	CannotLogout = &i18n.Message{
		ID:    "CannotLogout",
		Other: "cannot logout",
	}

	CannotCreateIAMUser = &i18n.Message{
		ID:    "CannotCreateIAMUser",
		Other: "cannot create IAM user",
	}

	CannotSendVerifyEmail = &i18n.Message{
		ID:    "CannotSendVerifyEmail",
		Other: "cannot send verify email",
	}
)
