package shm

import (
	"github.com/pkg/errors"
	"reflect"
)

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
		return 0, errors.Wrapf(ErrTypeUnsupported, "type: %s is unsupported", ref.Name())
	}
	return total, nil
}
