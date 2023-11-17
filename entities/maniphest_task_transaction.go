package entities

import "github.com/thought-machine/gonduit/util"

// ManiphestTaskTransaction represents a single task's transaction on Maniphest.
type ManiphestTaskTransaction struct {
	TaskID          string             `json:"taskID"`
	TransactionID   string             `json:"transactionID"`
	TransactionPHID string             `json:"transactionPHID"`
	TransactionType string             `json:"transactionType"`
	OldValue        interface{}        `json:"oldValue"`
	NewValue        interface{}        `json:"newValue"`
	Comments        string             `json:"comments"`
	AuthorPHID      string             `json:"authorPHID"`
	DateCreated     util.UnixTimestamp `json:"dateCreated"`
}

// Backwards-compatible alias to fix typo.
// Deprecated: Use ManiphestTaskTransaction instead.
type ManiphestTaskTranscation = ManiphestTaskTransaction
