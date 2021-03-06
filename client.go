package msngrclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var endpoint = "https://graph.facebook.com/v2.6/me/messages"

// Client wraps a configured instance
type Client interface {
	Send(*MessagePayload) error
}

type client struct {
	token      string
	httpClient *http.Client
}

// New returns a new client instance that will authenticate
// using the given token
func New(token string) Client {
	instance := &client{token, http.DefaultClient}
	return instance
}

// Send tries to send the passed request body to the messenger API
func (c *client) Send(payload *MessagePayload) error {
	requestBody, requestBodyErr := json.Marshal(payload)
	if requestBodyErr != nil {
		return fmt.Errorf("failed marshaling outbound payload: %v", requestBodyErr)
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		endpoint,
		bytes.NewBuffer(requestBody),
	)

	q := req.URL.Query()
	q.Set("access_token", c.token)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, callErr := c.httpClient.Do(req)
	if callErr != nil {
		return callErr
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		info := ResponseError{}
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &info)
		return fmt.Errorf(
			"failed sending message to messenger: status %v, message: %v",
			res.Status,
			info.Error.Message,
		)
	}
	return nil
}
