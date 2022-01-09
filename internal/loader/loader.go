package loader

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"strings"
)

type Loader struct {
	Parser     *hclparse.Parser
	FileSystem afero.Afero
}

func (l *Loader) Init() {
	l.Parser = hclparse.NewParser()
	l.FileSystem = afero.Afero{
		Fs: afero.NewOsFs(),
	}
}

func (l *Loader) LoadHCLFile(path string) (hcl.Body, hcl.Diagnostics) {
	src, err := l.FileSystem.ReadFile(path)
	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	var file *hcl.File
	var diags hcl.Diagnostics
	switch {
	case strings.HasSuffix(path, ".json"):
		file, diags = l.Parser.ParseJSON(src, path)
	default:
		file, diags = l.Parser.ParseHCL(src, path)
	}
	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), diags
	}

	return file.Body, diags
}

func (l *Loader) LoadTestConfig(path string) (map[string]string, hcl.Diagnostics) {
	config, err := l.FileSystem.ReadFile(path)
	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	data := make(map[string]map[string]string)
	errYaml := yaml.Unmarshal(config, &data)
	if errYaml != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	return data["variables"], nil
}
