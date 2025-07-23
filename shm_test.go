package shm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testStruct struct {
	Family   uint16
	Protocol uint16
	Ifindex  int32
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]uint8
}

func TestNewAndClose(t *testing.T) {
	layer, err := New[testStruct]()
	require.NoError(t, err)
	layer.Protocol++
	t.Logf("%+v", layer)
	err = Close(layer)
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
