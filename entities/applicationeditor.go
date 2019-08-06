package entities

// ObjectIdentifier is used for Edit methods.
type ObjectIdentifier struct {
	ID       int
	PHID     PHID
	Monogram string
}

// Transaction is used for Edit methods.
type Transaction struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
