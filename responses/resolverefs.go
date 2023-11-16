package responses

import "github.com/samwestmoreland/gonduit/entities"

// ResolveRefsResponse represents the result of calling diffusion.resolverefs.
type ResolveRefsResponse map[string][]*entities.ResolvedRef
