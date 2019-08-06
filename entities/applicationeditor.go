package entities

import (
	"errors"
	"fmt"
)

// ObjectIdentifier is used for Edit methods. The PHID is preferred if it is set, then the ID, then the Monogram.
type ObjectIdentifier struct {
	ID       int
	PHID     PHID
	Monogram string
}

func (o ObjectIdentifier) MarshalJSON() ([]byte, error) {
	if o.PHID != "" {
		return []byte(fmt.Sprintf(`"%s"`, o.PHID)), nil
	} else if o.ID != 0 {
		return []byte(fmt.Sprintf(`%d`, o.ID)), nil
	} else if o.Monogram != "" {
		return []byte(fmt.Sprintf(`"%s"`, o.Monogram)), nil
	}
	return nil, errors.New("no object identifier specified")
}

// Transaction is used for Edit methods.
type Transaction struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
