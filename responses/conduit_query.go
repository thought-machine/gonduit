package responses

import "github.com/thought-machine/gonduit/entities"

// ConduitQueryResponse is the response of calling conduit.query.
type ConduitQueryResponse map[string]*entities.ConduitMethod
