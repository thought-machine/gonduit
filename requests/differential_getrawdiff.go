package requests

// DifferentialGetRawDiffRequest represents a request to
// differential.getrawdiff.
type DifferentialGetRawDiffRequest struct {
	DiffID int64 `json:"diffID"`
	Request
}
