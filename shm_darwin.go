package shm

import (
	"os"
	"path/filepath"
)

func init() {
	defaultFinder = func(name string) string {
		return filepath.Join(os.TempDir(), name)
	}
}
