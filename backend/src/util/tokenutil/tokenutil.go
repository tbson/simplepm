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
	userID uint,
	clientSecret string,
	expSeconds int,
) (string, error) {
	localizer := localeutil.Get()
	token, err := jwt.NewBuilder().
		Subject(numberutil.UintToStr(userID)).
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
