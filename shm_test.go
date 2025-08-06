package shm

import (
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
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
	Bytes    [32]byte
}

func TestNewAndClose(t *testing.T) {
	object, err := New[testStruct]()
	require.NoError(t, err)
	require.NotNil(t, object)
	object.Int8++
	t.Logf("%+v", object)
	err = Close(object)
	require.NoError(t, err)
}

func TestOpen(t *testing.T) {
	name := `test`
	size := 1024
	ptr, err := open(defaultFinder(name), size, &Options{})
	require.NoError(t, err)
	require.NotEmpty(t, ptr)
	err = free(ptr, size)
	require.NoError(t, err)
}

func TestSizeof(t *testing.T) {
	object := testStruct{}
	size, err := Sizeof(object)
	require.NoError(t, err, err)
	require.Equal(t, int(unsafe.Sizeof(object)), size)
	t.Logf("[%d]%+v", size, object)
}

func TestUnsafeSizeof(t *testing.T) {
	object := testStruct{}
	size := unsafe.Sizeof(object)
	require.NotEmpty(t, size)
	t.Logf("unsafe.Sizeof: %d", size)
}

func TestMemset(t *testing.T) {
	object, err := New[testStruct]()
	require.NoError(t, err)
	require.NotNil(t, object)
	object.Int8++
	require.NotEmpty(t, object)
	Memset(object)
	require.Empty(t, object)
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

func BenchmarkSizeof(b *testing.B) {
	object := testStruct{}
	for i := 0; i < b.N; i++ {
		_, err := Sizeof(object)
		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func BenchmarkUnsafeSizeof(b *testing.B) {
	object := testStruct{}
	for i := 0; i < b.N; i++ {
		_ = unsafe.Sizeof(object)
	}
}

func BenchmarkMemset(b *testing.B) {
	object, err := New[testStruct]()
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		Memset(object)
	}
}
