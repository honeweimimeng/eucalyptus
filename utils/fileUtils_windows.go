//go:build windows

package utils

const (
	// Size to align the buffer to
	AlignSize = 4096

	// Minimum block size
	BlockSize = 4096

	// Extra flags for windows
	FILE_FLAG_NO_BUFFERING  = 0x20000000
	FILE_FLAG_WRITE_THROUGH = 0x80000000
)
