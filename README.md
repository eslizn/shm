# Go Shared Memory Library (shm)

A cross-platform Go library for creating and managing shared memory objects.

## Features

- Cross-platform support (Linux, Unix, Windows)
- Generic API supporting any type
- Simple and intuitive interface
- Thread-safe operations
- Configurable memory location

## Installation

```bash
go get github.com/eslizn/shm
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/eslizn/shm"
)

type Data struct {
	Value int
}

func main() {
	// Create a new shared memory object
	obj, err := shm.New[Data]()
	if err != nil {
		panic(err)
	}
	defer obj.Close()

	// Access the shared data
	obj.Value = 42
	fmt.Println("Value:", obj.Value)

	// Reset the memory to zero values
	err = obj.Memset()
	if err != nil {
		panic(err)
	}
}
```

## API Reference

### `func New[T any](opts ...Option) (*T, error)`

Creates or opens a shared memory object of type T.

Options:
- `WithName(name string)` - Specifies a custom name for the shared memory
- `WithFinder(finder Finder)` - Specifies a custom file finder

### `func (obj *T) Memset() error`

Resets the shared memory object to its zero value.

### `func (obj *T) Close() error`

Releases the shared memory object.

### `func Sizeof[T any]() uintptr`

Returns the size in bytes of type T.

## Platform Support

- **Linux**: Uses `/dev/shm` directory
- **Unix**: Uses system temp directory
- **Windows**: Uses Windows file mapping API

## Configuration

You can customize the shared memory location:

```go
// Custom name
obj, err := shm.New[Data](shm.WithName("my-shared-data"))

// Custom finder
finder := func(name string) string {
    return filepath.Join("/custom/path", name)
}
obj, err := shm.New[Data](shm.WithFinder(finder))
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

MIT
