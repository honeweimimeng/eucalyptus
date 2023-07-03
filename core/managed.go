package core

import "github.com/honeweimimeng/eucalyptus/meta"

type Managed struct {
}

func (d *Managed) ProcessRead(warp *meta.DataWrap) {

}

func (d *Managed) ProcessDataOb(ob meta.DataOb) meta.DataOption {
	return nil
}

func NewMultiLevelIO() meta.DataProcessor {
	return nil
}
