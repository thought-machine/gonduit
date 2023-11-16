package responses

import "github.com/samwestmoreland/gonduit/entities"

// DiffusionQueryCommitsResponse represents a response of the
// diffusion.querycommits call.
type DiffusionQueryCommitsResponse struct {
	Data          map[string]entities.DiffusionCommit `json:"data"`
	IdentifierMap map[string]string                   `json:"identifierMap"`
	Cursor        entities.Cursor                     `json:"cursor"`
}

// DiffusionBranchQueryResponse represents a response of the
// diffusion.branchquery call.
type DiffusionBranchQueryResponse []*entities.DiffusionBranch
