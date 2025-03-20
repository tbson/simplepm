package tokenutil

import (
	"context"
	"src/common/ctype"
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

func GenerateToken(
	subject string,
	claims ctype.Dict,
	secret string,
	expMins int,
) (string, error) {
	localizer := localeutil.Get()
	// Create a new JWT builder
	builder := jwt.NewBuilder()
	// Set the subject from the parameter
	if subject != "" {
		builder = builder.Subject(subject)
	} else {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MissingSubject,
		})
		return "", errutil.New("", []string{msg})
	}
	// Add all claims from the map (skip "sub" since we handle it separately)
	for key, value := range claims {
		if key != "sub" { // Skip "sub" since we're using the subject parameter
			builder = builder.Claim(key, value)
		}
	}
	// Add expiration only if expMins is not zero
	if expMins > 0 {
		duration := time.Duration(expMins) * time.Minute
		builder = builder.Expiration(time.Now().Add(duration))
	}
	// Build the token
	token, err := builder.Build()
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToBuildToken,
		})
		return "", errutil.New("", []string{msg})
	}
	// Sign the token using the provided secret
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToSignToken,
		})
		return "", errutil.New("", []string{msg})
	}
	return string(signed), nil
}

func VerifyToken(tokenStr string, secret string) (ctype.Dict, error) {
	localizer := localeutil.Get()
	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKey(jwa.HS256, []byte(secret)),
	)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseToken,
		})
		return nil, errutil.New("", []string{msg})
	}

	if err := jwt.Validate(token); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToVerifyToken,
		})
		return nil, errutil.New("", []string{msg})
	}

	subject := token.Subject()
	if subject == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.MissingSubject,
		})
		return nil, errutil.New("", []string{msg})
	}

	// Extract all claims as a map
	claims := make(map[string]interface{})
	claims["sub"] = subject
	ctx := context.Background()
	for iter := token.Iterate(ctx); iter.Next(ctx); {
		pair := iter.Pair()
		claims[pair.Key.(string)] = pair.Value
	}

	return claims, nil
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
			DefaultMessage: localeutil.FailedToBuildToken,
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
			DefaultMessage: localeutil.FailedToBuildToken,
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
