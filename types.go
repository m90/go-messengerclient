package msngrclient

// MessagePayload describes the request body
// that will be sent to the facebook API
type MessagePayload struct {
	Message      *Message          `json:"message,omitempty"`
	Recipient    *MessageRecipient `json:"recipient"`
	SenderAction string            `json:"sender_action,omitempty"`
}

// Message describes a single message object
type Message struct {
	Text       string             `json:"text,omitempty"`
	Attachment *MessageAttachment `json:"attachment,omitempty"`
}

// MessageAttachment describes an attachment
type MessageAttachment struct {
	Type    string                    `json:"type"`
	Payload *MessageAttachmentPayload `json:"payload"`
}

// MessageAttachmentPayload describes the payload of an attachment
type MessageAttachmentPayload struct {
	URL          string               `json:"url,omitempty"`
	TemplateType string               `json:"template_type,omitempty"`
	Text         string               `json:"text,omitempty"`
	Buttons      *[]AttachmentButton  `json:"buttons,omitempty"`
	Elements     *[]AttachmentElement `json:"elements,omitempty"`
}

// AddButton adds an AttachmentButton to an AttachmentPayload
// in a nil-safe way
func (pl *MessageAttachmentPayload) AddButton(btn AttachmentButton) {
	if pl.Buttons == nil {
		pl.Buttons = &[]AttachmentButton{btn}
	} else {
		*pl.Buttons = append(*pl.Buttons, btn)
	}
}

// AddElement adds an AttachmentElement to an AttachmentPayload
// in a nil-safe way
func (pl *MessageAttachmentPayload) AddElement(el AttachmentElement) {
	if pl.Elements == nil {
		pl.Elements = &[]AttachmentElement{el}
	} else {
		*pl.Elements = append(*pl.Elements, el)
	}
}

// AttachmentButton describes a button contained in an attachment
type AttachmentButton struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	Payload string `json:"payload,omitempty"`
}

// AttachmentElement describes an element contained in an attachment
type AttachmentElement struct {
	Title         string              `json:"title,omitempty"`
	ImageURL      string              `json:"image_url,omitempty"`
	Subtitle      string              `json:"subtitle,omitempty"`
	DefaultAction *AttachmentAction   `json:"default_action,omitempty"`
	Buttons       *[]AttachmentButton `json:"buttons,omitempty"`
}

// AddButton adds an AttachmentButton to an AttachmentElement
// in a nil-safe way
func (ae *AttachmentElement) AddButton(btn AttachmentButton) {
	if ae.Buttons == nil {
		ae.Buttons = &[]AttachmentButton{btn}
	} else {
		*ae.Buttons = append(*ae.Buttons, btn)
	}
}

// AttachmentAction describes an action contained in an attachment
type AttachmentAction struct {
	Type                string `json:"type,omitempty"`
	URL                 string `json:"url,omitempty"`
	MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
	WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	FallbackURL         string `json:"fallback_url,omitempty"`
}

// MessageRecipient contains the ID of the message's recipient
type MessageRecipient struct {
	ID string `json:"id"`
}

// ResponseError describes an erroroneous response by the messenger API
type ResponseError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}
