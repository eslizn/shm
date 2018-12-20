package shm

type Segment interface {
	Size() int
	Attach() (uintptr, error)
	Detach() error
}
