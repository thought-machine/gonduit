package responses

import "github.com/samwestmoreland/gonduit/entities"

// EditResponse is likely to change as ApplicationEditor evolves.
type EditResponse struct {
	// Object contains information about the object that was created or edited.
	Object struct {
		PHID entities.PHID `json:"phid"`
	} `json:"object"`
	// Transactions contains information about the transactions that were actually applied.
	Transactions []struct {
		TransactionPHID entities.PHID `json:"phid"`
	} `json:"transactions"`
}
