package datatypes

type StatusResponse struct {
	Version StatusVersion `json:"version"`

	Players StatusPlayers `json:"players"`

	Description        Chat   `json:"description"`
	Favicon            string `json:"favicon"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat"`
	PreviewsChat       bool   `json:"previewsChat"`
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
