package objects

type Chat struct {
	Text string `json:"text"`

	Bold          string `json:"bold,omitempty"`
	Italic        string `json:"italic,omitempty"`
	Underlined    string `json:"underlined,omitempty"`
	Strikethrough string `json:"strikethrough,omitempty"`
	Obfuscated    string `json:"obfuscated,omitempty"`

	Font  string `json:"font,omitempty"`
	Color string `json:"color,omitempty"`

	Insertion  string         `json:"insertion,omitempty"`
	ClickEvent ChatClickEvent `json:"clickEvent,omitempty"`
	HoverEvent ChatHoverEvent `json:"hoverEvent,omitempty"`

	Extra []Chat `json:"extra,omitempty"`
}

type UserSample struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type StatusPlayers struct {
	Max    int          `json:"max"`
	Online int          `json:"online"`
	Sample []UserSample `json:"sample"`
}

type StatusResponse struct {
	Version StatusVersion `json:"version"`

	Players StatusPlayers `json:"players"`

	Description        Chat   `json:"description"`
	Favicon            string `json:"favicon"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat"`
	PreviewsChat       bool   `json:"previewsChat"`
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

type Position struct {
	X int
	Y int
	Z int16
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}
