package datatypes

// StatusResponse is the response to a status request
// More info at https://wiki.vg/Server_List_Ping#Response
type StatusResponse struct {
	Version StatusVersion `json:"version"`

	Players StatusPlayers `json:"players"`

	Description        Chat   `json:"description"`
	Favicon            string `json:"favicon"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat"`
	PreviewsChat       bool   `json:"previewsChat"`
}

// UserSample is a single user that gets shown as an online example when hovered over the player count
type UserSample struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// StatusVersion is the minecraft version and protocol version of the server
type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

// StatusPlayers is the player count/max and sample list of online players
type StatusPlayers struct {
	Max    int          `json:"max"`
	Online int          `json:"online"`
	Sample []UserSample `json:"sample"`
}
