// +build windows

package config

import (
	"os"
	"path/filepath"
)

func DefaultConfig() string {
	return filepath.Join(
		os.Getenv("HOMEDRIVE"),
		os.Getenv("HOMEPATH"),
		FileName,
	)
}
