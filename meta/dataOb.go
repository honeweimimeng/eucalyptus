package meta

import (
	"io"
)

type DataProcessor interface {
	ProcessDataOb(ob DataOb) DataOption
	ProcessRead(warp *DataWrap)
}

type DataOption interface {
	Reader() io.Reader
	Writer() io.Writer
}

// GatherDataOption TODO: managed multiLevel io,gather io info
type GatherDataOption struct {
	reader io.Reader
	writer io.Writer
}

func NewGatherOption(writer io.Writer, reader io.Reader) *GatherDataOption {
	return &GatherDataOption{writer: writer, reader: reader}
}

func (op *GatherDataOption) Reader() io.Reader {
	return op.reader
}

func (op *GatherDataOption) Writer() io.Writer {
	return op.writer
}

type DataWrap struct {
	size int
	err  error
	data []byte
}
