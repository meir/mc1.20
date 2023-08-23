package datatypes

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

type ChatClickEvent struct {
	OpenUrl         string
	RunCommand      string
	SuggestCommand  string
	ChangePage      string
	CopyToClipboard string
}

type ChatHoverEvent struct {
	ShowText   string
	ShowItem   string
	ShowEntity string
}
