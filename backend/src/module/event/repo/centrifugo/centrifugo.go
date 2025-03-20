package centrifugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Repo struct{}

func New() Repo {
	return Repo{}
}

func (r Repo) Publish(data interface{}) error {
	localizer := localeutil.Get()
	apiKey := setting.CENTRIFUGO_API_KEY()
	url := fmt.Sprintf("%s/publish", setting.CENTRIFUGO_API_ENDPOINT())

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
