package responses

import "github.com/thought-machine/gonduit/entities"

// PasteQueryResponse represents the result of calling paste.query.
type PasteQueryResponse map[string]*entities.PasteItem
