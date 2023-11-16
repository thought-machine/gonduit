package responses

import "github.com/samwestmoreland/gonduit/entities"

// DifferentialQueryResponse is the response of calling differential.query.
type DifferentialQueryResponse []*entities.DifferentialRevision

// DifferentialQueryDiffsResponse is the response of calling differential.querydiffs.
type DifferentialQueryDiffsResponse map[string]*entities.DifferentialDiff
