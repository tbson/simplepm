package centrifugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/tokenutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Repo struct{}

func New() Repo {
	return Repo{}
}

func (r Repo) GetAuthJwt(userId uint) (string, error) {
	clientSecret := setting.CENTRIFUGO_CLIENT_SECRET
	lifeSpan := setting.CENTRIFUGO_JWT_LIFE_SPAN
	token, err := tokenutil.GenerateSimpleJWT(userId, clientSecret, lifeSpan)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r Repo) Publish(data ctype.SocketMessage) error {
	localizer := localeutil.Get()
	apiKey := setting.CENTRIFUGO_API_KEY
	url := fmt.Sprintf("%s/publish", setting.CENTRIFUGO_API_ENDPOINT)

	jsonData, err := json.Marshal(data)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotReadRequestBody,
		})
		return errutil.New("", []string{msg})
	}

	reqBody := bytes.NewBuffer(jsonData)

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CanNotCreateRequest,
		})
		return errutil.New("", []string{msg})
	}

	// Set headers
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CanNotSendRequest,
		})
		return errutil.New("", []string{msg})
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.BadRequest,
		})
		return errutil.New("", []string{msg})
	}
	return nil
}
