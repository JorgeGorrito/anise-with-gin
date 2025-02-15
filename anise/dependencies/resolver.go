package dependencies

import types "github.com/JorgeGorrito/anise-with-gin/anise/dependencies/types"

type Resolver interface {
	Resolve(abstract types.Abstract) types.Concrete
}
