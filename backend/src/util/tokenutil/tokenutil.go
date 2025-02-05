package tokenutil

import (
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/numberutil"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func GenerateSimpleJWT(
	clientID string,
	userID uint,
	clientSecret string,
	expSeconds int,
) (string, error) {
	localizer := localeutil.Get()
	userIDStr := numberutil.UintToStr(userID)
	token, err := jwt.NewBuilder().
		Subject(userIDStr).
		Claim("client", clientID).
		Expiration(time.Now().Add(time.Duration(expSeconds) * time.Second)).
		Build()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToCreateToken,
		})
		return "", errutil.New("", []string{msg})
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(clientSecret)))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToSignToken,
		})
		return "", errutil.New("", []string{msg})
	}

	return string(signed), nil
}

func GenerateSubscriptionJWT(
	clientID string,
	clientSecret string,
	channel string,
) (string, error) {
	// Create a new token with the desired claims.
	localizer := localeutil.Get()
	token, err := jwt.NewBuilder().
		Claim("client", clientID).
		Claim("channel", channel).
		Build()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToCreateToken,
		})
		return "", errutil.New("", []string{msg})
	}

	// Sign the token using HS256 with "secret" as the key.
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(clientSecret)))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToSignToken,
		})
		return "", errutil.New("", []string{msg})
	}

	return string(signed), nil
}
