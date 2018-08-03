package responses

import "github.com/thought-machine/gonduit/entities"

// ManiphestGetTaskTransactionsResponse is the response of calling maniphest.query.
type ManiphestGetTaskTransactionsResponse map[string][]*entities.ManiphestTaskTranscation
