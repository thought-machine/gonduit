package responses

import "github.com/thought-machine/gonduit/entities"

// PHIDQueryResponse is the result of phid.query operations.
type PHIDQueryResponse map[string]*entities.PHIDResult
