# shared memory for go

support unix(linux、darwin)/windows

## Usage

* New[T any](options ...Option) (*T, error)

```go
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
    require.NotNil(t, object)
}
```
