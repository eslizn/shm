package shm

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	name := filepath.Join(DefaultDir, "test")
	size := 1024
	ptr, err := Open(name, size)
	require.NoError(t, err)
	require.NotEmpty(t, ptr)
	err = Free(ptr, size)
	require.NoError(t, err)
}
