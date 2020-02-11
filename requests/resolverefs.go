package requests

import "github.com/thought-machine/gonduit/constants"

// ResolveRefsRequest represents a request to the diffusion.resolverefs endpoint.
type ResolveRefsRequest struct {
	Refs       []string `json:"refs"`
	Types      []string `json:"types,omitempty"`
	Repository string   `json:"repository,omitempty"`
	Branch     string   `json:"branch,omitempty"`
	Request
}
