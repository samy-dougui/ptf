package utils

import (
	"path"
	"path/filepath"
)

func NormalizePath(p string) (string, error) {
	fullPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return path.Clean(fullPath), nil
}
