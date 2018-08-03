package responses

import "github.com/thought-machine/gonduit/entities"

// RepositoryQueryResponse is the result of repository.query operations.
type RepositoryQueryResponse []*entities.Repository
