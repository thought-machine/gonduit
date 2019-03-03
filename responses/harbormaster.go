package responses

// HarbormasterCreateArtifactResponse is the response of calling harbormaster.createartifact.
type HarbormasterCreateArtifactResponse struct {
	Result struct {
		Data []struct {
			PHID string `json:"phid"`
		} `json:"data"`
	} `json:"result"`
}
