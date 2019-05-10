package shm

import "errors"

var (
	ErrUnspportedTarget = errors.New("unspported target")
)

func OpenSegment(size int, key int, flags int, mode int) (Segment, error) {
	return nil, ErrUnspportedTarget
}

func Open(key int) (Segment, error) {
	return OpenSegment(0, key, 0, 0666)
}

func Create(size int, flags int, mode int) (Segment, error) {
	return OpenSegment(size, 0, flags, mode)
}
