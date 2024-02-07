package entities

import (
	"encoding/json"

	"github.com/thought-machine/gonduit/util"
)

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
	Properties                json.RawMessage      `json:"properties"`
	AuthorName                string               `json:"authorName"`
	AuthorEmail               string               `json:"authorEmail"`
}

// A DifferentialChange represents a change to a file in Differential.
type DifferentialChange struct {
	ID          string             `json:"id"`
	OldPath     string             `json:"oldPath"`
	CurrentPath string             `json:"currentPath"`
	Type        string             `json:"type"`
	FileType    string             `json:"fileType"`
	CommitHash  string             `json:"commitHash"`
	AddLines    string             `json:"addLines"`
	DelLines    string             `json:"delLines"`
	Hunks       []DifferentialHunk `json:"hunks"`
}

// A DifferentialHunk is an absolute unit.
type DifferentialHunk struct {
	OldOffset string `json:"oldOffset"`
	NewOffset string `json:"newOffset"`
	OldLength string `json:"oldLength"`
	NewLength string `json:"newLength"`
	Corpus    string `json:"corpus"`
}

// A DifferentialProperties contains dynamic metadata about the Diff.
type DifferentialProperties struct {
	ArcOnto []DifferentialArcOnto `json:"arc:onto"`
}

// A DifferentialArcOnto contains a branch destination set with arc.
type DifferentialArcOnto struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// A DifferentialRawDiff is a raw diff.
type DifferentialRawDiff struct {
	Diff string `json:"diff"`
}

// DifferentialDiffUpdateTransaction adds a diff (by PHID) to the given revision
func DifferentialDiffUpdateTransaction(diff PHID) Transaction {
	return Transaction{
		Type:  "update",
		Value: diff,
	}
}

// DifferentialTitleTransaction retitles a revision
func DifferentialTitleTransaction(title string) Transaction {
	return Transaction{
		Type:  "title",
		Value: title,
	}
}

// DifferentialTitleTransaction changes the revision summary
func DifferentialSummaryTransaction(summary string) Transaction {
	return Transaction{
		Type:  "summary",
		Value: summary,
	}
}

// DifferentialTestPlanTransaction changes the revision test plan
func DifferentialTestPlanTransaction(testPlan string) Transaction {
	return Transaction{
		Type:  "testPlan",
		Value: testPlan,
	}
}

// DifferentialAddReviewersTransaction adds a list of reviewers by PHIDs
func DifferentialAddReviewersTransaction(reviewers []PHID) Transaction {
	return Transaction{
		Type:  "reviewers.add",
		Value: reviewers,
	}
}

// DifferentialRemoveReviewersTransaction adds a list of reviewers by PHIDs
func DifferentialRemoveReviewersTransaction(reviewers []PHID) Transaction {
	return Transaction{
		Type:  "reviewers.remove",
		Value: reviewers,
	}
}

// DifferentialSetReviewersTransaction changes the list of reviewers for this revision
func DifferentialSetReviewersTransaction(reviewers []PHID) Transaction {
	return Transaction{
		Type:  "reviewers.set",
		Value: reviewers,
	}
}

// DifferentialRepositoryTransaction changes the repository of the revision
func DifferentialRepositoryTransaction(repo PHID) Transaction {
	return Transaction{
		Type:  "repositoryPHID",
		Value: repo,
	}
}

// DifferentialAddTasksTransaction adds a list of tasks by PHIDs
func DifferentialAddTasksTransaction(tasks []PHID) Transaction {
	return Transaction{
		Type:  "tasks.add",
		Value: tasks,
	}
}

// DifferentialAddTasksTransaction removes a list of tasks by PHIDs
func DifferentialRemoveTasksTransaction(tasks []PHID) Transaction {
	return Transaction{
		Type:  "tasks.remove",
		Value: tasks,
	}
}

// DifferentialSetTasksTransaction updates the list of tasks by PHIDs
func DifferentialSetTasksTransaction(tasks []PHID) Transaction {
	return Transaction{
		Type:  "tasks.set",
		Value: tasks,
	}
}

// DifferentialAddParentsTransaction adds a list of parents by PHIDs
func DifferentialAddParentsTransaction(parents []PHID) Transaction {
	return Transaction{
		Type:  "parents.add",
		Value: parents,
	}
}

// DifferentialAddParentsTransaction removes a list of parents by PHIDs
func DifferentialRemoveParentsTransaction(parents []PHID) Transaction {
	return Transaction{
		Type:  "parents.remove",
		Value: parents,
	}
}

// DifferentialSetParentsTransaction updates the list of parents by PHIDs
func DifferentialSetParentsTransaction(parents []PHID) Transaction {
	return Transaction{
		Type:  "parents.set",
		Value: parents,
	}
}

// DifferentialAddChildrenTransaction adds a list of children by PHIDs
func DifferentialAddChildrenTransaction(children []PHID) Transaction {
	return Transaction{
		Type:  "children.add",
		Value: children,
	}
}

// DifferentialAddChildrenTransaction removes a list of children by PHIDs
func DifferentialRemoveChildrenTransaction(children []PHID) Transaction {
	return Transaction{
		Type:  "children.remove",
		Value: children,
	}
}

// DifferentialSetChildrenTransaction updates the list of children by PHIDs
func DifferentialSetChildrenTransaction(children []PHID) Transaction {
	return Transaction{
		Type:  "children.set",
		Value: children,
	}
}

// DifferentialPlanChangesTransaction updates the revision status to plan-changes
func DifferentialPlanChangesTransaction() Transaction {
	return Transaction{
		Type:  "plan-changes",
		Value: true,
	}
}

// DifferentialRequestReviewTransaction updates the revision status to request-review
func DifferentialRequestReviewTransaction() Transaction {
	return Transaction{
		Type:  "request-review",
		Value: true,
	}
}

// DifferentialCloseTransaction updates the revision status to close
func DifferentialCloseTransaction() Transaction {
	return Transaction{
		Type:  "close",
		Value: true,
	}
}

// DifferentialReopenTransaction updates the revision status to re-open it
func DifferentialReopenTransaction() Transaction {
	return Transaction{
		Type:  "reopen",
		Value: true,
	}
}

// DifferentialAbandonTransaction updates the revision status to abandoned
func DifferentialAbandonTransaction() Transaction {
	return Transaction{
		Type:  "abandon",
		Value: true,
	}
}

// DifferentialRejectTransaction updates the revision status to rejected
func DifferentialRejectTransaction() Transaction {
	return Transaction{
		Type:  "reject",
		Value: true,
	}
}

// DifferentialCommandeerTransaction updates the revision status to commandeered
func DifferentialCommandeerTransaction() Transaction {
	return Transaction{
		Type:  "commandeer",
		Value: true,
	}
}

// DifferentialResignTransaction updates the revision status to resign
func DifferentialResignTransaction() Transaction {
	return Transaction{
		Type:  "resign",
		Value: true,
	}
}

// DifferentialDraftTransaction updates the revision status to draft
func DifferentialDraftTransaction() Transaction {
	return Transaction{
		Type:  "draft",
		Value: true,
	}
}

// DifferentialAddProjectsTransaction adds a list of projects by PHIDs
func DifferentialAddProjectsTransaction(projects []PHID) Transaction {
	return Transaction{
		Type:  "projects.add",
		Value: projects,
	}
}

// DifferentialAddProjectsTransaction removes a list of projects by PHIDs
func DifferentialRemoveProjectsTransaction(projects []PHID) Transaction {
	return Transaction{
		Type:  "projects.remove",
		Value: projects,
	}
}

// DifferentialSetProjectsTransaction updates the list of projects by PHIDs
func DifferentialSetProjectsTransaction(projects []PHID) Transaction {
	return Transaction{
		Type:  "projects.set",
		Value: projects,
	}
}

// DifferentialAddSubscribersTransaction adds a list of subscribers by PHIDs
func DifferentialAddSubscribersTransaction(subscribers []PHID) Transaction {
	return Transaction{
		Type:  "subscribers.add",
		Value: subscribers,
	}
}

// DifferentialAddSubscribersTransaction removes a list of subscribers by PHIDs
func DifferentialRemoveSubscribersTransaction(subscribers []PHID) Transaction {
	return Transaction{
		Type:  "subscribers.remove",
		Value: subscribers,
	}
}

// DifferentialSetSubscribersTransaction updates the list of subscribers by PHIDs
func DifferentialSetSubscribersTransaction(subscribers []PHID) Transaction {
	return Transaction{
		Type:  "subscribers.set",
		Value: subscribers,
	}
}
