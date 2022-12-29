package loader

import (
	"fmt"
	"github.com/samy-dougui/ptf/internal/policy"
	"path"
	"strings"
)

var PolicyFileExtension = ".hcl"

func LoadPolicies(dirPath string) ([]*policy.Policy, error) {
	// TODO: better manage file loader
	var policies []*policy.Policy
	var fileLoader Loader
	fileLoader.Init()

	if isDir, err := fileLoader.FileSystem.DirExists(dirPath); !isDir {
		if err != nil {
			fmt.Printf("Path %s does not exist or is not a directory.", dirPath)
			return nil, err
		}
	}

	dirEntries, err := fileLoader.FileSystem.ReadDir(path.Clean(dirPath))
	if err != nil {
		return nil, err
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			subDir := path.Join(dirPath, entry.Name())
			subDirPolicies, err := LoadPolicies(subDir)
			if err != nil {
				fmt.Printf("Error loading policies in directory %s", subDir)
			} else {
				policies = append(policies, subDirPolicies...)
			}
		} else if strings.HasSuffix(entry.Name(), PolicyFileExtension) {
			fullPath := path.Join(dirPath, entry.Name())
			policies = append(policies, loadPoliciesFromFile(&fileLoader, fullPath)...)
		}
	}
	return policies, nil
}

func loadPoliciesFromFile(loader *Loader, path string) []*policy.Policy {
	var policies []*policy.Policy
	policiesBody, _ := loader.LoadHCLFile(path)
	policiesBlock, _ := policiesBody.Content(policy.PolicyFileSchema)
	for _, policyBlock := range policiesBlock.Blocks {
		switch policyBlock.Type {
		case "policy":
			var p policy.Policy
			p.Init(policyBlock)
			if !p.Disabled {
				policies = append(policies, &p)
			}
		default:
			continue
		}
	}
	return policies
}
