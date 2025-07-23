package shm

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"
)

var (
	ErrInvalidType = errors.New("invalid type")
	ErrSystemCall  = errors.New("system call failed")
	ErrInvalidSize = errors.New("invalid shared memory size")
)

// New create or open object in shared memory
func New[T any](options ...Option) (*T, error) {
	opt := getOptions(options...)
	typ := reflect.TypeOf((*T)(nil)).Elem()
	if len(opt.name) == 0 {
		opt.name = typ.Name()
	}
	size, err := Sizeof(typ)
	if err != nil {
		return nil, err
	}

	file := opt.finder(opt.name)
	err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
	if err != nil {
		return nil, err
	}

	ptr, err := Open(file, size)
	if err != nil {
		return nil, err
	}
	return (*T)(ptr), nil
}

// Close free object
func Close[T any](p *T) error {
	if p == nil {
		return nil
	}
	size, err := Sizeof(*p)
	if err != nil {
		return err
	}
	err = Free(unsafe.Pointer(p), size)
	if err != nil {
		return err
	}
	p = nil
	return nil
}
