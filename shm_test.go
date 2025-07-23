package shm

import (
	"github.com/stretchr/testify/require"
	"syscall"
	"testing"
)

func TestNew(t *testing.T) {
	layer, err := New[syscall.RawSockaddrLinklayer]()
	require.NoError(t, err)
	layer.Protocol++
	t.Logf("%+v", layer)
}
