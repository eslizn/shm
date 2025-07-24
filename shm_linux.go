package shm

import "path/filepath"

func init() {
	defaultFinder = func(name string) string {
		return filepath.Join("/dev/shm", name)
	}
}
