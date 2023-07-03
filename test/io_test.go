package test

import (
	"github.com/honeweimimeng/eucalyptus/core"
	"github.com/honeweimimeng/eucalyptus/meta"
	"os"
	"testing"
)

func TestDirectIO(t *testing.T) {
	infoFile := meta.NewNioObFile("/workspace/eucalyptus/test/info.txt", os.O_RDWR)
	buf := make([]byte, 2048)
	directIO := core.NewDirectIO()
	size, err := directIO.ProcessDataOb(infoFile).Reader().Read(buf)
	if err != nil {
		panic(err.Error())
	}
	println("read file ", infoFile.Name(), "success sizeOf:", size)
	println(string(buf))
}
