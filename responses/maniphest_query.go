package responses

import "github.com/etcinit/gonduit/entities"

// ManiphestQueryResponse is the response of calling maniphest.query.
type ManiphestQueryResponse map[string]*entities.ManiphestTask
