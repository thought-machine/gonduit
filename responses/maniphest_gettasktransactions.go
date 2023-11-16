package responses

import "github.com/samwestmoreland/gonduit/entities"

// ManiphestGetTaskTransactionsResponse is the response of calling maniphest.query.
type ManiphestGetTaskTransactionsResponse map[string][]*entities.ManiphestTaskTransaction
