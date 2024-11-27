package ssoutil

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/Nerzal/gocloak/v13"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

var client *gocloak.GoCloak

type UserInfo struct {
	ID          uint    `json:"id"`
	ExternalID  string  `json:"external_id"`
	Sub         *string `json:"sub"`
	Email       string  `json:"email"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	ProfileType string  `json:"profile_type"`
	TenantUid   string  `json:"tenant_uid"`
}

type TokensAndClaims struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Realm        string   `json:"realm"`
	UserInfo     UserInfo `json:"user_info"`
}

func Client() *gocloak.GoCloak {
	if client != nil {
		return client
	}
	client = gocloak.NewClient(setting.KEYCLOAK_URL)
	return client
}

func GetKeySet(jwksURL string) (jwk.Set, error) {
	localizer := localeutil.Get()
	ctx := context.Background()
	keySet, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeAuthorizationCode,
		})
		return nil, errutil.New("", []string{msg})
	}
	return keySet, nil
}

func EncodeState(stateData ctype.Dict) string {
	jsonData, _ := json.Marshal(stateData)
	return base64.URLEncoding.EncodeToString(jsonData)
}

func DecodeState(stateStr string) (ctype.Dict, error) {
	jsonData, err := base64.URLEncoding.DecodeString(stateStr)
	if err != nil {
		return nil, err
	}
	var stateData ctype.Dict
	err = json.Unmarshal(jsonData, &stateData)
	if err != nil {
		return nil, err
	}
	return stateData, nil
}
