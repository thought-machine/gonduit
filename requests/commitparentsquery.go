package requests

import "github.com/thought-machine/gonduit/constants"

// CommitParentsQueryRequest represents a request to the diffusion.commitparentsquery endpoint.
// One of "repository" or "callsign" is required.  Since "callsign" is deprecated we require "repository".
type CommitParentsQueryRequest struct {
	Commit     string `json:"commit"`
	Repository string `json:"repository"`
	Branch     string `json:"branch,omitempty"`
	Request
}
