package scripts

import "os"

// MakeDirectoryIfNotExists if not exist
func MakeDirectoryIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0755)
	}
}
