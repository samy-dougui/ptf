package utils

import (
	"path"
	"path/filepath"
)

func NormalizePath(p string) string {
	fullPath, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return path.Clean(fullPath)
}
