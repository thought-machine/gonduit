package entities

// ConduitMethod is a conduit method representation returned by
// `conduit.query`.
type ConduitMethod struct {
	Description string            `json:"description"`
	Params      map[string]string `json:"params"`
	Return      string            `json:"return"`
}
