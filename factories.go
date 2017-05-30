package msngrclient

// NewMessagePayload wraps the passed message and creates a ready
// to send request body
func NewMessagePayload(senderID string, msg *Message) *MessagePayload {
	return &MessagePayload{
		Recipient: &MessageSender{ID: senderID},
		Message:   msg,
	}
}

// NewTextPayload creates the payload for a basic text message
func NewTextPayload(message string) *Message {
	return &Message{
		Text: message,
	}
}

// NewImagePayload creates the payload for a basic image message
func NewImagePayload(imageURL string) *Message {
	return &Message{
		Attachment: &MessageAttachment{
			Type: "image",
			Payload: &MessageAttachmentPayload{
				URL: imageURL,
			},
		},
	}
}

// NewPostbackButton creates a new AttachmentButton with associated
// postback information
func NewPostbackButton(title, postback string) *AttachmentButton {
	return &AttachmentButton{
		Title:   title,
		Payload: postback,
		Type:    "postback",
	}
}
