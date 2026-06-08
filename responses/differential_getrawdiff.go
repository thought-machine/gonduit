package responses

import (
	"encoding/json"

	"github.com/thought-machine/gonduit/entities"
)

// DifferentialGetRawDiffResponse is the response from calling differential.getrawdiff.
type DifferentialGetRawDiffResponse struct {
	*entities.DifferentialRawDiff
}

// UnmarshalJSON handles the fact that Phabricator returns the JSON as a string.
func (r *DifferentialGetRawDiffResponse) UnmarshalJSON(b []byte) error {
	var rawString string
	if err := json.Unmarshal(b, &rawString); err == nil {
		r.DifferentialRawDiff = &entities.DifferentialRawDiff{
			Diff: rawString,
		}
		return nil
	}

	// If it is a real JSON object, unmarshal it into the underlying struct normally.
	r.DifferentialRawDiff = &entities.DifferentialRawDiff{}
	return json.Unmarshal(b, r.DifferentialRawDiff)
}
