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
	builder := jwt.NewBuilder()
	// Set the subject from the parameter
	if subject != "" {
		builder = builder.Subject(subject)
	} else {
		return "", errutil.New(localeutil.MissingSubject)
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
		return "", errutil.New(localeutil.FailedToBuildToken)
	}
	// Sign the token using the provided secret
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		return "", errutil.New(localeutil.FailedToSignToken)
	}
	return string(signed), nil
}

func VerifyToken(tokenStr string, secret string) (ctype.Dict, error) {
	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKey(jwa.HS256, []byte(secret)),
	)
	if err != nil {
		return nil, errutil.New(localeutil.FailedToParseToken)
	}

	if err := jwt.Validate(token); err != nil {
		return nil, errutil.New(localeutil.FailedToVerifyToken)
	}

	subject := token.Subject()
	if subject == "" {
		return nil, errutil.New(localeutil.MissingSubject)
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
	userIDStr := numberutil.UintToStr(userID)
	token, err := jwt.NewBuilder().
		Subject(userIDStr).
		Claim("client", clientID).
		Expiration(time.Now().Add(time.Duration(expSeconds) * time.Second)).
		Build()
	if err != nil {
		return "", errutil.New(localeutil.FailedToBuildToken)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(clientSecret)))
	if err != nil {
		return "", errutil.New(localeutil.FailedToSignToken)
	}

	return string(signed), nil
}

func GenerateSubscriptionJWT(
	clientID string,
	clientSecret string,
	channel string,
) (string, error) {
	// Create a new token with the desired claims.
	token, err := jwt.NewBuilder().
		Claim("client", clientID).
		Claim("channel", channel).
		Build()
	if err != nil {
		return "", errutil.New(localeutil.FailedToBuildToken)
	}

	// Sign the token using HS256 with "secret" as the key.
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(clientSecret)))
	if err != nil {
		return "", errutil.New(localeutil.FailedToSignToken)
	}

	return string(signed), nil
}
