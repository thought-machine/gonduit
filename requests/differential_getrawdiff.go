package requests

// DifferentialGetRawDiffRequest represents a request to
// differential.getrawdiff.
type DifferentialGetRawDiffRequest struct {
	DiffID string `json:"diffID,omitempty"`
	Request
}
