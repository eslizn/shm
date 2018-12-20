package shm

type Segment interface {
	Id() int
	Size() int
	Attach() (uintptr, error)
	Detach() error
}
