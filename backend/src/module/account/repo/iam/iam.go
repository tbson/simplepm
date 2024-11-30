package iam

import (
	"context"
	"fmt"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/ssoutil"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Repo struct {
	client *gocloak.GoCloak
}

func New(client *gocloak.GoCloak) Repo {
	return Repo{client: client}
}

func getJwksUrl(realm string) string {
	jwksURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/certs",
		setting.KEYCLOAK_URL,
		realm,
	)
	return jwksURL
}

func (r Repo) GetAuthUrl(
	realm string,
	clientId string,
	state ctype.Dict,
) string {
	stateStr := ssoutil.EncodeState(state)
	keycloakAuthURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/auth?client_id=%s&response_type=code&redirect_uri=%s&scope=openid profile email&state=%s",
		setting.KEYCLOAK_URL,
		realm,
		clientId,
		setting.KEYCLOAK_REDIRECT_URI,
		stateStr,
	)
	return keycloakAuthURL
}

func (r Repo) GetLogoutUrl(
	realm string,
	clientId string,
) string {
	keycloakAuthURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/logout?client_id=%s&post_logout_redirect_uri=%s",
		setting.KEYCLOAK_URL,
		realm,
		clientId,
		setting.KEYCLOAK_POST_LOGOUT_URI,
	)
	return keycloakAuthURL
}

func (r Repo) ValidateCallback(
	ctx context.Context,
	realm string,
	clientId string,
	clientSecret string,
	code string,
) (ssoutil.TokensAndClaims, error) {
	var result ssoutil.TokensAndClaims
	localizer := localeutil.Get()

	if code == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.AuthorizationCodeNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	// Exchange the code for tokens
	token, err := r.client.GetToken(ctx, realm, gocloak.TokenOptions{
		ClientID:     gocloak.StringP(clientId),
		ClientSecret: gocloak.StringP(clientSecret),
		RedirectURI:  gocloak.StringP(setting.KEYCLOAK_REDIRECT_URI),
		Code:         gocloak.StringP(code),
		GrantType:    gocloak.StringP("authorization_code"),
	})

	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeAuthorizationCode,
		})
		return result, errutil.New("", []string{msg})
	}

	idToken := token.IDToken
	accesToken := token.AccessToken
	refreshToken := token.RefreshToken

	userInfo, err := r.ValidateToken(idToken, realm)
	if err != nil {
		r.Logout(clientId, clientSecret, realm, refreshToken)
		return result, err
	}

	ssoUserInfo, ssoResult := r.client.GetUserInfo(ctx, token.AccessToken, realm)
	if ssoResult != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotGetUserInfo,
		})
		return result, errutil.New("", []string{msg})
	}
	userInfo.Sub = ssoUserInfo.Sub

	result = ssoutil.TokensAndClaims{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
		Realm:        realm,
		UserInfo:     userInfo,
	}
	return result, nil
}

// TODO: Implement checking kid
func (r Repo) ValidateToken(tokenStr string, realm string) (ssoutil.UserInfo, error) {
	localizer := localeutil.Get()
	result := ssoutil.UserInfo{}
	jwksURL := getJwksUrl(realm)

	// Fetch the JWKS (public keys)
	keySet, err := ssoutil.GetKeySet(jwksURL)
	if err != nil {
		return result, err
	}

	// Parse the JWT token to extract headers and claims
	clockSkew := time.Duration(setting.KEYCLOAK_CLOCK_SKEW) * time.Minute
	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKeySet(keySet),
		jwt.WithClock(jwt.ClockFunc(time.Now().UTC)), // Use UTC time directly
		jwt.WithAcceptableSkew(clockSkew),            // Set clock skew tolerance
	)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseToken,
		})
		return result, errutil.New("", []string{msg})
	}

	// Check if the token is expired by inspecting the "exp" claim
	if err := jwt.Validate(
		token,
		jwt.WithClock(jwt.ClockFunc(time.Now().UTC)),
		jwt.WithAcceptableSkew(clockSkew),
	); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.TokenHasExpired,
		})
		return result, errutil.New("", []string{msg})
	}

	// If verification is successful, print the claims
	claims := token.PrivateClaims()

	result = ssoutil.UserInfo{
		ExternalID: claims["preferred_username"].(string),
		Email:      claims["email"].(string),
		FirstName:  claims["given_name"].(string),
		LastName:   claims["family_name"].(string),
	}

	return result, nil
}

func (r Repo) GetSub(tokenStr string, realm string) (string, error) {
	localizer := localeutil.Get()
	result := ""

	// Parse the JWT token to extract headers and claims
	clockSkew := time.Duration(setting.KEYCLOAK_CLOCK_SKEW) * time.Minute
	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithVerify(false),                        // Skip signature verification
		jwt.WithClock(jwt.ClockFunc(time.Now().UTC)), // Use UTC time directly
		jwt.WithAcceptableSkew(clockSkew),            // Set clock skew tolerance
	)
	if err != nil {
		fmt.Println("Error parsing token")
		fmt.Println(err)
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseToken,
		})
		return result, errutil.New("", []string{msg})
	}

	// Check if the token is expired by inspecting the "exp" claim
	if err := jwt.Validate(
		token,
		jwt.WithClock(jwt.ClockFunc(time.Now().UTC)),
		jwt.WithAcceptableSkew(clockSkew),
	); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.TokenHasExpired,
		})
		return result, errutil.New("", []string{msg})
	}

	// If verification is successful, print the claims
	sub, ok := token.Get("sub")
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.SubClaimNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	return sub.(string), nil
}

func (r Repo) RefreshToken(
	ctx context.Context,
	realm string,
	refreshToken string,
	clientId string,
	clientSecret string,
) (ssoutil.TokensAndClaims, error) {
	var result ssoutil.TokensAndClaims
	localizer := localeutil.Get()
	if refreshToken == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.RefreshTokenNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	// Exchange the refresh token for new tokens
	token, err := r.client.RefreshToken(
		ctx,
		refreshToken,
		clientId,
		clientSecret,
		realm,
	)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeRefreshToken,
		})
		return result, errutil.New("", []string{msg})
	}

	idToken := token.IDToken
	accesToken := token.AccessToken
	refreshToken = token.RefreshToken

	userInfo, err := r.ValidateToken(idToken, realm)
	if err != nil {
		return result, err
	}

	result = ssoutil.TokensAndClaims{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
		Realm:        realm,
		UserInfo:     userInfo,
	}
	return result, nil
}

func (r Repo) GetAdminAccessToken() (string, error) {
	ctx := context.Background()
	adminUser := setting.KEYCLOAK_ADMIN
	adminPassword := setting.KEYCLOAK_ADMIN_PASSWORD
	token, err := r.client.LoginAdmin(ctx, adminUser, adminPassword, "master")
	if err != nil {
		msg := localeutil.Get().MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotLoginAdmin,
		})
		return "", errutil.New("", []string{msg})
	}
	return token.AccessToken, nil
}

func (r Repo) CreateUser(
	accessToken string,
	realm string,
	email string,
	firstName string,
	lastName string,
	mobile *string,
) (string, error) {
	ctx := context.Background()
	localizer := localeutil.Get()

	user := gocloak.User{
		FirstName: gocloak.StringP(firstName),
		LastName:  gocloak.StringP(lastName),
		Email:     gocloak.StringP(email),
		Enabled:   gocloak.BoolP(true),
		Attributes: &map[string][]string{
			"mobile": {fmt.Sprintf("%s", *mobile)},
		},
	}

	sub, err := r.client.CreateUser(ctx, accessToken, realm, user)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotCreateIAMUser,
		})
		return "", errutil.New("", []string{msg})
	}
	return sub, nil
}

func (r Repo) SendVerifyEmail(
	accessToken string,
	clientID string,
	sub string,
	realm string,
) error {
	ctx := context.Background()
	localizer := localeutil.Get()

	redirectUri := setting.KEYCLOAK_REDIRECT_URI
	params := gocloak.SendVerificationMailParams{
		ClientID:    &clientID,
		RedirectURI: &redirectUri,
	}

	err := r.client.SendVerifyEmail(ctx, accessToken, sub, realm, params)
	if err != nil {
		fmt.Println("Error sending verify email")
		fmt.Println(err)
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotSendVerifyEmail,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}

func (r Repo) UpdateUser(
	accessToken string,
	realm string,
	sub string,
	data ctype.Dict,
) error {
	ctx := context.Background()
	localizer := localeutil.Get()

	user, err := r.client.GetUserByID(context.Background(), accessToken, realm, sub)
	user.FirstName = gocloak.StringP(data["FirstName"].(string))
	user.LastName = gocloak.StringP(data["LastName"].(string))
	user.Attributes = &map[string][]string{
		"mobile": {data["Mobile"].(string)},
	}

	err = r.client.UpdateUser(ctx, accessToken, realm, *user)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotUpdateIAMUser,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}

func (r Repo) SetPassword(
	accessToken string,
	sub string,
	realm string,
	password string,
) error {
	ctx := context.Background()
	localizer := localeutil.Get()
	err := r.client.SetPassword(ctx, accessToken, sub, realm, password, false)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotSetPassword,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}

func (r Repo) Logout(clientId string, clientSecret string, realm string, refreshToken string) error {
	ctx := context.Background()
	localizer := localeutil.Get()
	err := r.client.Logout(ctx, clientId, clientSecret, realm, refreshToken)
	if err != nil {
		fmt.Println("Error logging out")
		fmt.Println(err)
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotLogout,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}
