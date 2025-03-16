package srv

import (
	"src/common/ctype"

	"src/module/account/usecase/auth/checkauthurl/model"
)

type localDataProvider interface {
	GetAuthClientFromTenantUid(tenantUid string) (model.AuthClientInfo, error)
}

type iamProvider interface {
	GetAuthUrl(realm string, clientId string, state ctype.Dict) string
}

type srv struct {
	localDataRepo localDataProvider
	iamRepo       iamProvider
}

func New(localDataRepo localDataProvider, iamRepo iamProvider) srv {
	return srv{localDataRepo, iamRepo}
}

func (srv srv) GetAuthUrl(tenantUid string, nextParam string) (string, error) {
	state := ctype.Dict{
		"tenantUid": tenantUid,
	}
	if nextParam != "" {
		state["next"] = nextParam
	}

	authClientInfo, err := srv.localDataRepo.GetAuthClientFromTenantUid(tenantUid)
	if err != nil {
		return "", err
	}

	realm := authClientInfo.Realm
	clientId := authClientInfo.ClientID

	url := srv.iamRepo.GetAuthUrl(realm, clientId, state)

	return url, nil
}
