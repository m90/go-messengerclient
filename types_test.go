package msngrclient

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAddButton_Attachment(t *testing.T) {
	tests := []struct {
		initial  MessageAttachmentPayload
		addition AttachmentButton
		expected MessageAttachmentPayload
	}{
		{
			MessageAttachmentPayload{},
			AttachmentButton{
				Type:  "foo",
				Title: "bar",
			},
			MessageAttachmentPayload{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
				},
			},
		},
		{
			MessageAttachmentPayload{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
				},
			},
			AttachmentButton{
				Type:  "bar",
				Title: "baz",
			},
			MessageAttachmentPayload{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
					AttachmentButton{
						Type:  "bar",
						Title: "baz",
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.initial.AddButton(test.addition)
		if !reflect.DeepEqual(test.initial, test.expected) {
			t.Errorf("Expected result of %+v, got %+v", test.expected, test.initial)
		}
	}
}

func TestAddElement(t *testing.T) {
	tests := []struct {
		initial  MessageAttachmentPayload
		addition AttachmentElement
		expected MessageAttachmentPayload
	}{
		{
			MessageAttachmentPayload{},
			AttachmentElement{
				Title: "foo",
			},
			MessageAttachmentPayload{
				Elements: &[]AttachmentElement{
					AttachmentElement{
						Title: "foo",
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.initial.AddElement(test.addition)
		if !reflect.DeepEqual(test.initial, test.expected) {
			t.Errorf("Expected result of %+v, got %+v", test.expected, test.initial)
		}
	}
}

func TestAddButton_Element(t *testing.T) {
	tests := []struct {
		initial  AttachmentElement
		addition AttachmentButton
		expected AttachmentElement
	}{
		{
			AttachmentElement{},
			AttachmentButton{
				Type:  "foo",
				Title: "bar",
			},
			AttachmentElement{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
				},
			},
		},
		{
			AttachmentElement{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
				},
			},
			AttachmentButton{
				Type:  "bar",
				Title: "baz",
			},
			AttachmentElement{
				Buttons: &[]AttachmentButton{
					AttachmentButton{
						Type:  "foo",
						Title: "bar",
					},
					AttachmentButton{
						Type:  "bar",
						Title: "baz",
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.initial.AddButton(test.addition)
		if !reflect.DeepEqual(test.initial, test.expected) {
			t.Errorf("Expected result of %+v, got %+v", test.expected, test.initial)
		}
	}
}

type mockMarshaler struct{}

func (m *mockMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`{"key": "value"}`), nil
}
func TestMessagePayload(t *testing.T) {
	tests := []struct {
		name     string
		payload  MessagePayload
		expected string
	}{
		{
			"default",
			MessagePayload{
				Message: &Message{
					Text: "O HAI!",
				},
				Recipient: &MessageRecipient{
					ID: "zuck",
				},
			},
			`{"message":{"text":"O HAI!"},"recipient":{"id":"zuck"}}`,
		},
		{
			"custom marshaler",
			MessagePayload{
				Message: &mockMarshaler{},
				Recipient: &MessageRecipient{
					ID: "zuck",
				},
			},
			`{"message":{"key":"value"},"recipient":{"id":"zuck"}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := json.Marshal(test.payload)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if string(result) != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, string(result))
			}
		})
	}
}
