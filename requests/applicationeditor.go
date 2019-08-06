package requests

import "github.com/thought-machine/gonduit/entities"

type EditRequest struct {
	ObjectIdentifier entities.ObjectIdentifier `json:"objectIdentifier"`
	Transactions     []entities.Transaction    `json:"transactions"`
	Request
}
