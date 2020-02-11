package entities

// ResolvedRef is the result item for resolverefs queries.
type ResolvedRef struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Alternate  string `json:"alternate,omitempty"`
}
