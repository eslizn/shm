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

	ptr, err := open(file, size, opt)
	if err != nil {
		return nil, err
	}
	return (*T)(ptr), nil
}

// Memset reset object to zero val
func Memset[T any](p *T) {
	*p = *new(T)
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
	err = free(unsafe.Pointer(p), size)
	if err != nil {
		return err
	}
	p = nil
	return nil
}

// Sizeof calc object size
func Sizeof(in any) (int, error) {
	ref, assert := in.(reflect.Type)
	if !assert {
		ref = reflect.TypeOf(in)
	}

	total := 0
	switch ref.Kind() {
	case reflect.Struct:
		var structSize = 0
		var maxAlign = 1
		for i := 0; i < ref.NumField(); i++ {
			field := ref.Field(i)
			align := field.Type.Align()
			if align > maxAlign {
				maxAlign = align
			}
			// field align
			if structSize%align != 0 {
				structSize += align - structSize%align
			}
			size, err := Sizeof(field.Type)
			if err != nil {
				return 0, err
			}
			structSize += size
		}
		// struct align
		if structSize%maxAlign != 0 {
			structSize += maxAlign - structSize%maxAlign
		}
		total += structSize
	case reflect.Array:
		size, err := Sizeof(ref.Elem())
		if err != nil {
			return 0, err
		}
		total += size * ref.Len()
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Bool:
		total += int(ref.Size())
	default:
		return 0, errors.Wrapf(ErrInvalidType, "type: %s is unsupported", ref.Name())
	}
	return total, nil
}
