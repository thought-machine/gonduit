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

// DiffusionRepository represents a VCS repository.
type DiffusionRepository struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	PHID        string                 `json:"phid"`
	Callsign    string                 `json:"callsign"`
	Monogram    string                 `json:"monogram"`
	VCS         string                 `json:"vcs"`
	URI         string                 `json:"uri"`
	RemoteURI   string                 `json:"remoteURI"`
	Description string                 `json:"description"`
	IsActive    bool                   `json:"isActive"`
	IsHosted    bool                   `json:"isHosted"`
	IsImporting bool                   `json:"isImporting"`
	Encoding    string                 `json:"encoding"`
	Staging     map[string]interface{} `json:"staging"`
}
