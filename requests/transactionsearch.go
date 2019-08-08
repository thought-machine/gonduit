package requests

type TransactionSearchRequest struct {
	ObjectID    string                 `json:"objectIdentifier,omitempty"`
	Constraints map[string]interface{} `json:"constraints,omitempty"`
	Before      string                 `json:"before,omitempty"`
	After       string                 `json:"after,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Request
}