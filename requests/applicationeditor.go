package requests

import "github.com/samwestmoreland/gonduit/entities"

type EditRequest struct {
	ObjectIdentifier entities.ObjectIdentifier `json:"objectIdentifier"`
	Transactions     []entities.Transaction    `json:"transactions"`
	Request
}
