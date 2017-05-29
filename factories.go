package msngrclient

// NewMessagePayload wraps the passed message and creates a ready
// to send request body
func NewMessagePayload(senderID string, msg *Message) *MessagePayload {
	return &MessagePayload{
		Recipient: &MessageSender{ID: senderID},
		Message:   msg,
	}
}
