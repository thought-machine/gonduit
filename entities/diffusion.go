package entities

// DiffusionCommit represents a commit in Diffusion.
type DiffusionCommit struct {
	ID             string `json:"id"`
	PHID           string `json:"phid"`
	RepositoryPHID string `json:"repositoryPHID"`
	Identifier     string `json:"identifier"`
	Epoch          string `json:"epoch"`
	URI            string `json:"uri"`
	IsImporting    bool   `json:"isImporting"`
	Summary        string `json:"summary"`
	AuthorPHID     string `json:"authorPHID"`
	CommitterPHID  string `json:"committerPHID"`
	Author         string `json:"author"`
	AuthorName     string `json:"authorName"`
	AuthorEmail    string `json:"authorEmail"`
	Committer      string `json:"committer"`
	CommitterName  string `json:"committerName"`
}

// DiffusionBranch represents a branch in Diffusion.
type DiffusionBranch struct {
	ShortName string `json:"shortName"`
	CommitID  string `json:"commitIdentifier"`
	RefType   string `json:"refType"`
	RawFields struct {
		ObjectName string `json:"objectName"`
		ObjectType string `json:"objectType"`
		Refname    string `json:"refname"`
		Subject    string `json:"subject"`
		Creator    string `json:"creator"`
		Author     string `json:"author"`
		Epoch      string `json:"epoch"`
	} `json:"RawFields"`
}
