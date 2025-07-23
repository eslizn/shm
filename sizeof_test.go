package shm

import (
	"github.com/stretchr/testify/require"
	"syscall"
	"testing"
	"unsafe"
)

func Test_Sizeof(t *testing.T) {
	obj := syscall.RawSockaddrLinklayer{}
	size, err := Sizeof(obj)
	require.NoError(t, err, err)
	require.Equal(t, int(unsafe.Sizeof(obj)), size)
	t.Logf("[%d]%+v", size, obj)
}
