package health

type (
	// infoResponse – describes a response for status event.
	infoResponse struct {
		Name    string `json:"name"`
		Commit  string `json:"commit,omitempty"`
		Date    string `json:"date,omitempty"`
		Version string `json:"version"`
	}
)
