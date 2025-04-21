package srv

import (
	"src/common/ctype"
	"src/module/account/domain/model"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type authTokenProvider interface {
	VerifyRefreshToken(token string) (uint, error)
	GenerateTokenPair(userID uint) (model.TokenPair, error)
}

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
}

type srv struct {
	userRepo     userProvider
	authTokenSrv authTokenProvider
}

func New(userRepo userProvider, authTokenSrv authTokenProvider) srv {
	return srv{userRepo, authTokenSrv}
}

func (srv srv) RefreshToken(refreshToken string) (model.TokenPair, error) {
	localizer := localeutil.Get()

	userID, err := srv.authTokenSrv.VerifyRefreshToken(refreshToken)
	if err != nil {
		return model.TokenPair{}, err
	}

	opts := ctype.QueryOpts{
		Filters: ctype.Dict{"id": userID},
	}
	user, err := srv.userRepo.Retrieve(opts)
	if err != nil {
		return model.TokenPair{}, err
	}

	if user.LockedAt != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.LockedAccount,
		})
		return model.TokenPair{}, errutil.New("", []string{msg})
	}

	return srv.authTokenSrv.GenerateTokenPair(userID)

}
