package srv

import (
	"src/common/ctype"
	"src/module/account/domain/model"
	"src/module/account/schema"
	"src/util/errutil"
	"src/util/localeutil"
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
		return model.TokenPair{}, errutil.New(localeutil.LockedAccount)
	}

	return srv.authTokenSrv.GenerateTokenPair(userID)

}
