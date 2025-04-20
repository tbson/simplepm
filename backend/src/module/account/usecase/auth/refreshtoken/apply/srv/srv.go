package srv

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"

	"src/module/account/domain/srv/authtoken"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type authProvider interface {
	VerifyRefreshToken(token string) (uint, error)
	GenerateTokenPair(userID uint) (account.TokenPair, error)
}

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
}

type srv struct {
	userRepo userProvider
}

func New(userRepo userProvider) srv {
	return srv{userRepo}
}

func (srv srv) RefreshToken(refreshToken string) (account.TokenPair, error) {
	tokenSettings := setting.AUTH_TOKEN_SETTINGS()
	authTokenSrv := authtoken.New(
		tokenSettings.AccessTokenSecret,
		tokenSettings.RefreshTokenSecret,
		tokenSettings.AccessTokenLifetime,
		tokenSettings.RefreshTokenLifetime,
	)

	localizer := localeutil.Get()

	userID, err := authTokenSrv.VerifyRefreshToken(refreshToken)
	if err != nil {
		return account.TokenPair{}, err
	}

	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	}
	user, err := srv.userRepo.Retrieve(opts)
	if err != nil {
		return account.TokenPair{}, err
	}

	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return account.TokenPair{}, errutil.New("", []string{msg})
	}

	return authTokenSrv.GenerateTokenPair(userID)

}
