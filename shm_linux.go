package shm

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"
)

const (
	IPC_CREATE  = 00001000
	IPC_EXCL    = 00002000
	IPC_NOWAIT  = 00004000
	IPC_DIPC    = 00010000
	IPC_OWN     = 00020000
	IPC_PRIVATE = 0
	IPC_RMID    = 0
	IPC_SET     = 1
	IPC_STAT    = 2
	IPC_INFO    = 3
	IPC_OLD     = 0
	IPC_64      = 0x0100
)

//from bits/ipc.h
type shmid_ds struct {
	ipc_perm struct {
		key     uint32
		uid     uint32
		gid     uint32
		cuid    uint32
		cgid    uint32
		mode    uint32
		pad1    uint16
		seq     uint16
		pad2    uint16
		unused1 uint
		unused2 uint
	}
	shm_segsz   uint32
	shm_atime   uint64
	shm_dtime   uint64
	shm_ctime   uint64
	shm_cpid    uint32
	shm_lpid    uint32
	shm_nattch  uint16
	shm_unused  uint16
	shm_unused2 uintptr
	shm_unused3 uintptr
}

type segment struct {
	sync.Mutex
	id   uintptr
	addr uintptr
	*shmid_ds
}

func (s *segment) Id() int {
	s.Lock()
	defer s.Unlock()
	return int(s.id)
}

func (s *segment) Size() int {
	s.Lock()
	defer s.Unlock()
	if s.shmid_ds != nil {
		return int(s.shmid_ds.shm_segsz)
	}
	return 0
}

func (s *segment) Attach() (uintptr, error) {
	s.Lock()
	defer s.Unlock()
	if s.addr == 0 {
		var (
			err syscall.Errno
		)
		s.addr, _, err = syscall.Syscall(syscall.SYS_SHMAT, s.id, 0, 0)
		if err != 0 {
			return 0, errors.New(err.Error())
		}
	}
	return s.addr, nil
}

func (s *segment) Detach() error {
	s.Lock()
	defer s.Unlock()
	if s.addr != 0 {
		var (
			err syscall.Errno
		)
		_, _, err = syscall.Syscall(syscall.SYS_SHMDT, s.addr, 0, 0)
		if err != 0 {
			return errors.New(err.Error())
		}
		s.addr = 0
	}
	return nil
}

func OpenSegment(size int, key int, flags int, mode int) (Segment, error) {
	var (
		err syscall.Errno
		seg = &segment{}
	)
	seg.id, _, err = syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), uintptr(mode&flags))
	if err != 0 {
		return nil, errors.New(err.Error())
	}
	seg.shmid_ds = &shmid_ds{}
	_, _, err = syscall.Syscall(syscall.SYS_SHMCTL, seg.id, uintptr(IPC_STAT), uintptr(unsafe.Pointer(seg.shmid_ds)))
	if err != 0 {
		return nil, errors.New(err.Error())
	}
	return seg, nil
}

func Open(key int) (Segment, error) {
	return OpenSegment(0, key, ^IPC_CREATE, 0666)
}

func Create(size int, flags int, mode int) (Segment, error) {
	return OpenSegment(size, IPC_PRIVATE, flags, mode)
}
