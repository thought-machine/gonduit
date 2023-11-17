package responses

import "github.com/thought-machine/gonduit/entities"

// ResolveRefsResponse represents the result of calling diffusion.resolverefs.
type ResolveRefsResponse map[string][]*entities.ResolvedRef
