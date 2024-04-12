package osx

import (
	"os"
)

// IsFile reports wheter file exists.
func IsFile(filename string) bool {
	return exists(filename, false)
}

// IsDir reports wheter directory exists.
func IsDir(dirname string) bool {
	return exists(dirname, true)
}

func exists(name string, dir bool) bool {
	switch info, err := os.Stat(name); {
	case os.IsNotExist(err):
		return false
	case dir:
		return info.IsDir()
	default:
		return !info.IsDir()
	}
}
