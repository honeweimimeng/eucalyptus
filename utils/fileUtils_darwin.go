//go:build darwin

package utils

import (
	"fmt"
	"os"
	"syscall"
)

const (
	// OSX doesn't need any alignment
	AlignSize = 0

	// Minimum block size
	BlockSize = 4096
)

func OpenFileSys(fileName string, flag int, model os.FileMode) *os.File {
	file, err := os.OpenFile(fileName, flag, model)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %s,may be that is not exit", fileName))
		return nil
	}
	_, _, e1 := syscall.Syscall(syscall.SYS_FCNTL, file.Fd(), syscall.F_NOCACHE, 1)
	if e1 != 0 {
		err = fmt.Errorf("failed to set F_NOCACHE: %s", e1)
		err = file.Close()
		if err != nil {
			panic("option of file close is field")
		}
		file = nil
	}
	return file
}
