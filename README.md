# shared memory for go

support unix(linux、darwin)/windows

## Usage

* New[T any](options ...Option) (*T, error)

```go
func TestNew(t *testing.T) {
	layer, err := New[syscall.RawSockaddrLinklayer]()
	require.NoError(t, err)
	layer.Protocol++
	t.Logf("%+v", layer)
}
```
