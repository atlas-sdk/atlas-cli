package sdk

import (
	"fmt"
	"os"
	"strings"
)

func AddToPath(newDir string) error {
	path := os.Getenv("PATH")
	if path == "" {
		return fmt.Errorf("PATH não está definida")
	}

	for _, dir := range strings.Split(path, string(os.PathListSeparator)) {
		if dir == newDir {
			return nil
		}
	}

	newPath := newDir + string(os.PathListSeparator) + path
	return os.Setenv("PATH", newPath)
}
