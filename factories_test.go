package msngrclient

import (
	"reflect"
	"testing"
)

func TestNewMessagePayload(t *testing.T) {
	tests := []struct {
		senderID string
		message  *Message
		expected *MessagePayload
	}{
		{
			"abc123",
			nil,
			&MessagePayload{
				Recipient:    &MessageRecipient{"abc123"},
				Message:      nil,
				SenderAction: "",
			},
		},
		{
			"abc123",
			&Message{
				Text: "hey",
			},
			&MessagePayload{
				Recipient: &MessageRecipient{"abc123"},
				Message: &Message{
					Text: "hey",
				},
				SenderAction: "",
			},
		},
	}
	for _, test := range tests {
		result := NewMessagePayload(test.senderID, test.message)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %+v, got %+v", test.expected, result)
		}
	}
}
