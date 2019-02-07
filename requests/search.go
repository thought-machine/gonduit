package requests

import "github.com/thought-machine/gonduit/constants"

// SearchRequest represents a request to one of the .search endpoints.
type SearchRequest struct {
	QueryKey    string                 `json:"queryKey,omitempty"`
	Constraints map[string]interface{} `json:"constraints,omitempty"`
	Attachments map[string]bool        `json:"attachments,omitempty"`
	Order       constants.SearchOrder  `json:"order,omitempty"`
	Before      string                 `json:"before,omitempty"`
	After       string                 `json:"after,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
}
