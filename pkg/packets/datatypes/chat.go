package datatypes

// Chat is a chat message
// This is used in any location where text can be formatted such as the server status, chat messages, and books.
// For more info see https://wiki.vg/Chat
type Chat struct {
	Text string `json:"text"`

	Bold          bool `json:"bold,omitempty"`
	Italic        bool `json:"italic,omitempty"`
	Underlined    bool `json:"underlined,omitempty"`
	Strikethrough bool `json:"strikethrough,omitempty"`
	Obfuscated    bool `json:"obfuscated,omitempty"`

	Font  string `json:"font,omitempty"`
	Color string `json:"color,omitempty"`

	Insertion  string         `json:"insertion,omitempty"`
	ClickEvent ChatClickEvent `json:"clickEvent,omitempty"`
	HoverEvent ChatHoverEvent `json:"hoverEvent,omitempty"`

	Extra []Chat `json:"extra,omitempty"`
}

// ChatClickEvent is an event that happens when a chat message is clicked
// For more info see https://wiki.vg/Chat
type ChatClickEvent struct {
	OpenUrl         string
	RunCommand      string
	SuggestCommand  string
	ChangePage      string
	CopyToClipboard string
}

// ChatHoverEvent is an event that happens when a chat message is hovered over
// For more info see https://wiki.vg/Chat
type ChatHoverEvent struct {
	ShowText   string
	ShowItem   string
	ShowEntity string
}
