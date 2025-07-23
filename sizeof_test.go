package shm

import (
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

func Test_Sizeof(t *testing.T) {
	object := testStruct{}
	size, err := Sizeof(object)
	require.NoError(t, err, err)
	require.Equal(t, int(unsafe.Sizeof(object)), size)
	t.Logf("[%d]%+v", size, object)
}

func BenchmarkSizeof(b *testing.B) {
	object := testStruct{}
	for i := 0; i < b.N; i++ {
		_, err := Sizeof(object)
		if err != nil {
			b.Error(err)
		}
	}
}
