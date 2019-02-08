package entities

import "github.com/thought-machine/gonduit/util"

// DifferentialRevision represents a revision in Differential.
type DifferentialRevision struct {
	ID             string                 `json:"id"`
	PHID           string                 `json:"phid"`
	Title          string                 `json:"title"`
	URI            string                 `json:"uri"`
	DateCreated    util.UnixTimestamp     `json:"dateCreated"`
	DateModified   util.UnixTimestamp     `json:"dateModified"`
	AuthorPHID     string                 `json:"authorPHID"`
	Status         string                 `json:"status"`
	StatusName     string                 `json:"statusName"`
	Branch         string                 `json:"branch"`
	Summary        string                 `json:"summary"`
	TestPlan       string                 `json:"testPlan"`
	LineCount      string                 `json:"lineCount"`
	ActiveDiffPHID string                 `json:"activeDiffPHID"`
	Diffs          []string               `json:"diffs"`
	Commits        []string               `json:"commits"`
	Reviewers      map[string]string      `json:"reviewers"`
	CCs            []string               `json:"ccs"`
	Hashes         [][]string             `json:"hashes"`
	Auxiliary      map[string]interface{} `json:"auxiliary"`
	RepositoryPHID string                 `json:"repositoryPHID"`
}

// A DifferentialDiff represents a diff in Differential.
type DifferentialDiff struct {
	ID                        string               `json:"id"`
	RevisionID                string               `json:"revisionID"`
	DateCreated               util.UnixTimestamp   `json:"dateCreated"`
	DateModified              util.UnixTimestamp   `json:"dateModified"`
	SourceControlBaseRevision string               `json:"sourceControlBaseRevision"`
	SourceControlPath         string               `json:"sourceControlPath"`
	SourceControlSystem       string               `json:"sourceControlSystem"`
	Branch                    string               `json:"branch"`
	Bookmark                  string               `json:"bookmark"`
	CreationMethod            string               `json:"creationMethod"`
	Description               string               `json:"description"`
	UnitStatus                string               `json:"unitStatus"`
	LintStatus                string               `json:"lintStatus"`
	Changes                   []DifferentialChange `json:"changes"`
	AuthorName                string               `json:"authorName"`
	AuthorEmail               string               `json:"authorEmail"`
}

// A DifferentialChange represents a change to a file in Differential.
type DifferentialChange struct {
	ID          string                 `json:"id"`
	Metadata    map[string]interface{} `json:"metadata"`
	OldPath     string                 `json:"oldPath"`
	CurrentPath string                 `json:"currentPath"`
	Type        string                 `json:"type"`
	FileType    string                 `json:"fileType"`
	CommitHash  string                 `json:"commitHash"`
	AddLines    string                 `json:"addLines"`
	DelLines    string                 `json:"delLines"`
	Hunks       []DifferentialHunk     `json:"hunks"`
}

// A DifferentialHunk is an absolute unit.
type DifferentialHunk struct {
	OldOffset string `json:"oldOffset"`
	NewOffset string `json:"newOffset"`
	OldLength string `json:"oldLength"`
	NewLength string `json:"newLength"`
	Corpus    string `json:"corpus"`
}
