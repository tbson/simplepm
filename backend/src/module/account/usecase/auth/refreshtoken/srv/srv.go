package srv

import (
	"context"
	"src/common/authtype"
	"src/module/account/usecase/auth/checkauthurl/model"
)

type localDataProvider interface {
	GetAuthClientFromSub(sub string) (model.AuthClientInfo, error)
}

type iamProvider interface {
	GetSub(tokenStr string, realm string) (string, error)
	RefreshToken(
		ctx context.Context,
		realm string,
		refreshToken string,
		clientId string,
		clientSecret string,
	) (authtype.SsoCallbackResult, error)
}

type srv struct {
	localDataRepo localDataProvider
	iamRepo       iamProvider
}

func New(localDataRepo localDataProvider, iamRepo iamProvider) srv {
	return srv{localDataRepo, iamRepo}
}

func (srv srv) RefreshToken(
	ctx context.Context,
	realm string,
	refreshToken string,
) (authtype.SsoCallbackResult, error) {
	sub, err := srv.iamRepo.GetSub(refreshToken, realm)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	authClientInfo, err := srv.localDataRepo.GetAuthClientFromSub(sub)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	clientId := authClientInfo.ClientID
	clientSecret := authClientInfo.ClientSecret

	tokensAndClaims, err := srv.iamRepo.RefreshToken(
		ctx, realm, refreshToken, clientId, clientSecret,
	)
	if err != nil {
		return authtype.SsoCallbackResult{}, err
	}

	return tokensAndClaims, nil
}
