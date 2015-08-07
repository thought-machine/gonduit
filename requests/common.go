package requests

import "github.com/etcinit/gonduit/entities"

// Request is the base request struct.
type Request struct {
	Session *entities.Session `json:"__conduit__"`
}
