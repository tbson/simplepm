package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"src/common/setting"
	"src/util/errutil"
	"src/util/i18nmsg"
)

type adapter struct {
	apiKey string
	url    string
}

func New() adapter {
	apiKey := setting.CENTRIFUGO_API_KEY()
	url := fmt.Sprintf("%s/publish", setting.CENTRIFUGO_API_ENDPOINT())
	return adapter{
		apiKey: apiKey,
		url:    url,
	}
}

func (r adapter) Publish(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errutil.New(i18nmsg.CannotReadRequestBody)
	}

	reqBody := bytes.NewBuffer(jsonData)

	// Create the HTTP request
	req, err := http.NewRequest("POST", r.url, reqBody)
	if err != nil {
		return errutil.New(i18nmsg.CanNotCreateRequest)
	}

	// Set headers
	req.Header.Set("X-API-Key", r.apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errutil.New(i18nmsg.CanNotSendRequest)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return errutil.New(i18nmsg.BadRequest)
	}
	return nil
}
