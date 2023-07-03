package core

import (
	"github.com/honeweimimeng/eucalyptus/utils"
	"log"
	"unsafe"
)

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func IsAligned(block []byte) bool {
	return alignment(block, utils.AlignSize) == 0
}

func AlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+utils.AlignSize)
	a := alignment(block, utils.AlignSize)
	offset := 0
	if a != 0 {
		offset = utils.AlignSize - a
	}
	block = block[offset : offset+BlockSize]
	if BlockSize != 0 {
		if !IsAligned(block) {
			log.Fatal("Failed to align block")
		}
	}
	return block
}
