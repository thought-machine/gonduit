package requests

import "github.com/thought-machine/gonduit/constants"

// DifferentialSetDiffPropertyRequest represents a request to the differential.setdiffproperty endpoint.
type DifferentialSetDiffPropertyRequest struct {
	// ID of the diff (not revision) to set the property on.
	DiffID uint64 `json:"diff_id"`
	// Name of the property to set
	Name string `json:"name"`
	// Value of the property. Must be a JSON-encoded string representing an object or array of objects.
	Data string `json:"data"`
	Request
}
