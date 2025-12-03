package shm

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// open create or open memory block
func open(file string, size int, options *Options) (unsafe.Pointer, error) {
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

// flush sync to disk
func flush(ptr unsafe.Pointer, size int) error {
	//@TODO FlushFileBuffers
	return windows.FlushViewOfFile(uintptr(ptr), uintptr(size))
}

// free freeze memory block
func free(ptr unsafe.Pointer, size int) error {
	err := flush(ptr, size)
	if err != nil {
		return err
	}
	return windows.UnmapViewOfFile(uintptr(ptr))
}
