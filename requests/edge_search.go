package requests

// EdgeSearchRequest represents a request to the edge.search endpoints.
type EdgeSearchRequest struct {
	SourcePhids      []string `json:"sourcePHIDs"`
	Types            []string `json:"types"`
	DestinationPhids []string `json:"destinationPHIDs,omitempty"`
	Before           string   `json:"before,omitempty"`
	After            string   `json:"after,omitempty"`
	Limit            int      `json:"limit,omitempty"`
	Request
}
