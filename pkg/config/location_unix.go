// +build linux darwin

package config

import (
	"os"
	"path"
)

func DefaultConfig() string {
	return path.Join(os.Getenv("HOME"), FileName)
}
