package srv

import (
	"src/common/ctype"
	"src/module/account"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type authProvider interface {
	VerifyRefreshToken(token string) (uint, error)
	GenerateTokenPair(userID uint) (account.TokenPair, error)
}

type userProvider interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.User, error)
}

type srv struct {
	authSrv  authProvider
	userRepo userProvider
}

func New(authSrv authProvider, userRepo userProvider) srv {
	return srv{authSrv, userRepo}
}

func (srv srv) RefreshToken(refreshToken string) (account.TokenPair, error) {
	localizer := localeutil.Get()

	userID, err := srv.authSrv.VerifyRefreshToken(refreshToken)
	if err != nil {
		return account.TokenPair{}, err
	}

	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"id": userID},
	}
	user, err := srv.userRepo.Retrieve(queryOptions)
	if err != nil {
		return account.TokenPair{}, err
	}

	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return account.TokenPair{}, errutil.New("", []string{msg})
	}

	return srv.authSrv.GenerateTokenPair(userID)

}
