package dependencies

import types "github.com/JorgeGorrito/anise-with-gin/anise/dependencies/types"

type Binder interface {
	Bind(abstract types.Abstract, concrete types.Concrete)
}
