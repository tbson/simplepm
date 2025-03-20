package auth

import (
	"src/module/account"

	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/numberutil"
	"src/util/tokenutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type srv struct {
}

func New() *srv {
	return &srv{}
}

func (r srv) VerifyAccessToken(token string) (uint, error) {
	localizer := localeutil.Get()
	tokenSecret := setting.ACCESS_TOKEN_SECRET()
	claims, err := tokenutil.VerifyToken(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	typ := claims["typ"].(string)
	if typ != "access" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidAccessToken,
		})
		return 0, errutil.New("", []string{msg})
	}
	userID := uint(claims["sub"].(float64))
	return userID, nil
}

func (r srv) VerifyRefreshToken(token string) (uint, error) {
	localizer := localeutil.Get()
	tokenSecret := setting.REFRESH_TOKEN_SECRET()
	claims, err := tokenutil.VerifyToken(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	typ := claims["typ"].(string)
	if typ != "refresh" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidRefreshToken,
		})
		return 0, errutil.New("", []string{msg})
	}
	userID := uint(claims["sub"].(float64))
	return userID, nil
}

func (r srv) GenerateTokenPair(userID uint) (account.TokenPair, error) {
	accessTokenSecret := setting.ACCESS_TOKEN_SECRET()
	accessTokenLifetime := setting.ACCESS_TOKEN_LIFETIME()

	refreshTokenSecret := setting.REFRESH_TOKEN_SECRET()
	refreshTokenLifetime := setting.REFRESH_TOKEN_LIFETIME()

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
		return account.TokenPair{}, err
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
		return account.TokenPair{}, err
	}

	result := account.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}

func (r *srv) SetPwd(userID uint, pwd string) error {
	return nil
}

func (r *srv) SendVerifyEmail(userID uint) error {
	return nil
}

func (r *srv) VerifyOTP(code string) error {
	return nil
}

func (r *srv) RefreshToken(refreshToken uint) (account.TokenPair, error) {
	result := account.TokenPair{}
	return result, nil
}
