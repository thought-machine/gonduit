package requests

import "github.com/thought-machine/gonduit/constants"

// SearchRequest represents a request to one of the .search endpoints.
type SearchRequest struct {
	QueryKey    string                 `json:"queryKey"`
	Constraints map[string]interface{} `json:"constraints"`
	Attachments map[string]bool        `json:"attachments"`
	Order       constants.SearchOrder  `json:"order"`
	Before      string                 `json:"before"`
	After       string                 `json:"after"`
	Limit       int                    `json:"limit"`
}
