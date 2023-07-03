package meta

import (
	"github.com/honeweimimeng/eucalyptus/utils"
	"os"
)

type ObFile struct {
	path    string
	model   int
	process DataProcessor
	file    *os.File
}

func NewNioObFile(path string, model int) *ObFile {
	var obFile ObFile
	obFile.path = path
	obFile.model = model
	return &obFile
}

func (o *ObFile) initFile() {
	if o.model == 0 {
		o.model = os.O_RDWR | os.O_CREATE
	}
	o.file = utils.OpenFileSys(o.path, o.model, 0666)
}

func (o *ObFile) InitProcessor(processor DataProcessor) {
	o.process = processor
}

func (o *ObFile) Processor() DataProcessor {
	return o.process
}

func (o *ObFile) Name() string {
	return o.path
}

func (o *ObFile) Close() error {
	return o.file.Close()
}
