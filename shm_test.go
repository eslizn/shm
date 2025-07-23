package shm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testStructChild struct {
	Int8  int8
	int16 int16
	Int32 int32
	Int64 int64
}

type testStruct struct {
	Int8     int8
	int16    int16
	Int32    int32
	Int64    int64
	Uint8    uint8
	Uint16   uint16
	Uint32   uint32
	Uint64   uint64
	Float32  float32
	Float64  float64
	Bool     bool
	ArrayInt [8]int64
	Children [8]testStructChild
}

func TestNewAndClose(t *testing.T) {
	object, err := New[testStruct]()
	require.NoError(t, err)
	object.Int8++
	t.Logf("%+v", object)
	err = Close(object)
	require.NoError(t, err)
}

func TestOpen(t *testing.T) {
	name := `test`
	size := 1024
	ptr, err := Open(defaultNamer(name), size)
	require.NoError(t, err)
	require.NotEmpty(t, ptr)
	err = Free(ptr, size)
	require.NoError(t, err)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		object, err := New[testStruct]()
		if err != nil {
			b.Error(err)
		}
		b.StopTimer()
		err = Close(object)
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
