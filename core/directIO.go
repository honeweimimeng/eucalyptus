package core

import "github.com/honeweimimeng/eucalyptus/meta"

type DirectIO struct {
}

func (d *DirectIO) ProcessRead(warp *meta.DataWrap) {

}

func (d *DirectIO) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *DirectIO) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (d *DirectIO) ProcessDataOb(ob meta.DataOb) meta.DataOption {
	if init, isInit := ob.(meta.DataObInitializer); isInit {
		init.InitProcessor(d)
	}
	return meta.NewGatherOption(d, d)
}

func NewDirectIO() *DirectIO {
	return &DirectIO{}
}
