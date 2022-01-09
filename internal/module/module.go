package module

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/tftest/internal/resource"
	"github.com/samy-dougui/tftest/internal/variable"
)

type Module struct {
	Directory string
	Variables map[string]*variable.Variable
	//Locals    map[string]*Local
	//Outputs   map[string]*Output
	ManagedResources map[string]*resource.Resource
	Context          hcl.EvalContext
}
