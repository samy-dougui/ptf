package loader

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
)

type Loader struct {
	Parser     *hclparse.Parser
	FileSystem afero.Afero
}

func (l *Loader) Init() {
	// TODO: refacto using init function instead of struct method
	l.Parser = hclparse.NewParser()
	l.FileSystem = afero.Afero{
		Fs: afero.NewOsFs(),
	}
}
