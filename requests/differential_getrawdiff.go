package requests

import "github.com/samwestmoreland/gonduit/responses"

// DifferentialGetRawDiffRequest represents a request to
// differential.getrawdiff.
type DifferentialGetRawDiffRequest struct {
	DiffID responses.SearchData `json:"diffID"`
	Request
}
