package shm

func init() {
	defaultFinder = func(name string) string {
		return filepath.Join("/dev/shm", name)
	}
}
