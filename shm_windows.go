package shm

import (
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"unsafe"
)

var defaultNamer = func(name string) string {
	return filepath.Join(os.TempDir(), name)
}

// Open create or open memory block
func Open(file string, size int) (unsafe.Pointer, error) {
	fp, err := windows.CreateFile(
		windows.StringToUTF16Ptr(file),
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_ALWAYS,
		windows.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return nil, err
	}

	mapping, err := windows.CreateFileMapping(
		fp,
		nil,
		windows.PAGE_READWRITE,
		0,
		uint32(size),
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(mapping)

	addr, err := windows.MapViewOfFile(
		mapping,
		windows.FILE_MAP_WRITE|windows.FILE_MAP_READ,
		0, 0, uintptr(size),
	)
	if err != nil {
		return nil, err
	}
	return unsafe.Pointer(addr), nil
}

// Free freeze memory block
func Free(ptr unsafe.Pointer, size int) error {
	return windows.UnmapViewOfFile(uintptr(ptr))
}
