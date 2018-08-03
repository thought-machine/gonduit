package responses

import "github.com/thought-machine/gonduit/entities"

// PHIDLookupResponse is the result of phid.lookup operations.
type PHIDLookupResponse map[string]*entities.PHIDResult
