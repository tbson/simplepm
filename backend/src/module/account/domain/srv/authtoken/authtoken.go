package authtoken

import (
	"src/module/account/domain/model"

	"src/util/errutil"
	"src/util/localeutil"
	"src/util/numberutil"
	"src/util/tokenutil"
)

type srv struct {
	accessTokenSecret    string
	refreshTokenSecret   string
	accessTokenLifetime  int
	refreshTokenLifetime int
}

func New(
	accessTokenSecret string,
	refreshTokenSecret string,
	accessTokenLifetime int,
	refreshTokenLifetime int,
) *srv {
	return &srv{
		accessTokenSecret:    accessTokenSecret,
		refreshTokenSecret:   refreshTokenSecret,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

func (r srv) VerifyAccessToken(token string) (uint, error) {
	tokenSecret := r.accessTokenSecret
	claims, err := tokenutil.VerifyToken(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	typ := claims["typ"].(string)
	if typ != "access" {
		return 0, errutil.New(localeutil.InvalidAccessToken)
	}
	userID := numberutil.StrToUint(claims["sub"].(string), 0)
	return userID, nil
}

func (r srv) VerifyRefreshToken(token string) (uint, error) {
	tokenSecret := r.refreshTokenSecret
	claims, err := tokenutil.VerifyToken(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	typ := claims["typ"].(string)
	if typ != "refresh" {
		return 0, errutil.New(localeutil.InvalidRefreshToken)
	}
	userID := numberutil.StrToUint(claims["sub"].(string), 0)
	return userID, nil
}

func (r srv) GenerateTokenPair(userID uint) (model.TokenPair, error) {
	accessTokenSecret := r.accessTokenSecret
	accessTokenLifetime := r.accessTokenLifetime

	refreshTokenSecret := r.refreshTokenSecret
	refreshTokenLifetime := r.refreshTokenLifetime

	sub := numberutil.UintToStr(userID)
	accessToken, err := tokenutil.GenerateToken(
		sub,
		map[string]interface{}{
			"sub": userID,
			"typ": "access",
		},
		accessTokenSecret,
		accessTokenLifetime,
	)
	if err != nil {
		return model.TokenPair{}, err
	}

	refreshToken, err := tokenutil.GenerateToken(
		sub,
		map[string]interface{}{
			"sub": userID,
			"typ": "refresh",
		},
		refreshTokenSecret,
		refreshTokenLifetime,
	)
	if err != nil {
		return model.TokenPair{}, err
	}

	result := model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}

func (r *srv) RefreshToken(refreshToken uint) (model.TokenPair, error) {
	// to be implemented
	result := model.TokenPair{}
	return result, nil
}
