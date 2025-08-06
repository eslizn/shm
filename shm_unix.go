//go:build linux || darwin || freebsd

package shm

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
	"unsafe"
)

// open create or open memory block
func open(file string, size int, options *Options) (unsafe.Pointer, error) {
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

	if stat.Size != int64(size) {
		if !options.force {
			return nil, errors.Wrapf(ErrInvalidSize, "file size: %d, apply: %d", stat.Size, size)
		}
		err = unix.Ftruncate(fd, 0)
		if err != nil {
			return nil, errors.Wrap(ErrSystemCall, err.Error())
		}
		_, err = unix.Seek(fd, 0, 0)
		if err != nil {
			return nil, errors.Wrap(ErrSystemCall, err.Error())
		}
		err = unix.Fstat(fd, &stat)
		if err != nil {
			return nil, err
		}
	}

	if stat.Size == 0 {
		err = unix.Ftruncate(fd, int64(size))
		if err != nil {
			return nil, errors.Wrap(ErrSystemCall, err.Error())
		}
	}

	mapping, err := unix.Mmap(fd, 0, size, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(&mapping[0]), nil
}

// free freeze memory block
func free(ptr unsafe.Pointer, size int) error {
	err := unix.MunmapPtr(ptr, uintptr(size))
	if err != nil {
		return err
	}
	return nil
}
