package requests

// DiffusionQueryCommitsRequest represents a request to the
// diffusion.querycommits call.
type DiffusionQueryCommitsRequest struct {
	IDs            []uint64 `json:"ids"`
	PHIDs          []string `json:"phids"`
	Names          []string `json:"names"`
	RepositoryPHID string   `json:"repositoryPHID"`
	NeedMessages   bool     `json:"needMessages"`
	BypassCache    bool     `json:"bypassCache"`
	Before         string   `json:"before"`
	After          string   `json:"after"`
	Limit          uint64   `json:"limit"`
	Request
}

// DiffusionBranchQueryRequest represents a request to the
// diffusion.branchquery call.
type DiffusionBranchQueryRequest struct {
	Closed     bool     `json:"closed"`
	Contains   string   `json:"contains"`
	Patterns   []string `json:"patterns"`
	Repository string   `json:"repository"`
	Branch     string   `json:"branch"`
	Offset     uint64   `json:"offset"`
	Limit      uint64   `json:"limit"`
	Request
}
