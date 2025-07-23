package shm

import (
	"os"
	"path/filepath"
)

func init() {
	defaultNamer = func(name string) string {
		return filepath.Join(os.TempDir(), name)
	}
}
