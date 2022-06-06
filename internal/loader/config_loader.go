package loader

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"log"
	"path"
	"strings"
)

func (l *Loader) LoadConfigDir(dirPath string) (hcl.Body, hcl.Diagnostics) {
	//var diag hcl.Diagnostics
	if isDir, err := l.FileSystem.DirExists(dirPath); !isDir {
		if err != nil {
			log.Fatalf("Error while loading the directory %e", err)
		}
		return hcl.EmptyBody(), hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf("The path %v is not a directory", dirPath),
			},
		}
	}
	dirEntries, _ := l.FileSystem.ReadDir(path.Clean(dirPath))
	var bodies []hcl.Body
	for _, entry := range dirEntries {
		if entry.IsDir() {
			subDirBody, _ := l.LoadConfigDir(path.Join(dirPath, entry.Name()))
			bodies = append(bodies, subDirBody)
		} else if strings.HasSuffix(entry.Name(), ".hcl") {
			fileBody, _ := l.LoadHCLFile(path.Join(dirPath, entry.Name()))
			bodies = append(bodies, fileBody)
		}
	}
	return hcl.MergeBodies(bodies), hcl.Diagnostics{}
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
