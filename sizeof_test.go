package shm

import (
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

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
