package shm

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
	"unsafe"
)

// Open create or open memory block
func Open(file string, size int) (unsafe.Pointer, error) {
	fd, err := unix.Open(file, unix.O_CREAT|unix.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer unix.Close(fd)

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		return nil, err
	}

	if stat.Size == 0 {
		err = unix.Ftruncate(fd, int64(size))
		if err != nil {
			return nil, errors.Wrap(ErrSystemCall, err.Error())
		}
	} else if stat.Size != int64(size) {
		return nil, errors.Wrapf(ErrInvalidSize, "file size: %d, apply: %d", stat.Size, size)
	}

	mapping, err := unix.Mmap(fd, 0, size, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(&mapping[0]), nil
}

// Free freeze memory block
func Free(ptr unsafe.Pointer, size int) error {
	err := unix.MunmapPtr(ptr, uintptr(size))
	if err != nil {
		return err
	}
	return nil
}
