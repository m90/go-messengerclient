package msngrclient

import "testing"
import "net/http/httptest"
import "net/http"
import "io/ioutil"

func TestSend(t *testing.T) {
	tests := []struct {
		name         string
		token        string
		payload      *MessagePayload
		expectedBody string
		fail         bool
	}{
		{
			"default",
			"abc123",
			NewMessagePayload("zuck", NewTextMessage("i like that")),
			`{"message":{"text":"i like that"},"recipient":{"id":"zuck"}}`,
			false,
		},
		{
			"default",
			"",
			NewMessagePayload("zuck", NewTextMessage("i like that")),
			``,
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			previousEndpoint := endpoint
			defer func() {
				endpoint = previousEndpoint
			}()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if test.fail {
					http.Error(w, "Failing for testing purposes", http.StatusInternalServerError)
				} else {
					body, _ := ioutil.ReadAll(r.Body)
					if test.expectedBody != string(body) {
						t.Errorf("Expected body of %v, got %v", test.expectedBody, string(body))
					}
					if test.token != r.URL.Query().Get("access_token") {
						t.Errorf("Expected access_token of %v, got %v", test.token, r.URL.Query().Get("access_token"))
					}
				}
			}))
			endpoint = ts.URL
			client := New(test.token)
			err := client.Send(test.payload)
			if !test.fail && err != nil {
				t.Errorf("Expected clean test end, got %v", err)
			} else if test.fail && err == nil {
				t.Error("Expected Send to return error")
			}
		})
	}
}
