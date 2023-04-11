package file

import (
	"os"
	"path/filepath"
)

// LoadBytes loads file as byte slice
func LoadBytes(path string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(path))
}
