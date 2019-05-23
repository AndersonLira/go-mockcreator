package cache

import (
	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/io"
)


type FileCacheExecutor struct {
	Next chain.Executor
}

func (self FileCacheExecutor) Get(xmlS string) (string,error) {
	cfg := config.GetConfig()

	fileName := cfg.PayloadFolder + "/" + xml.NameSugested(xmlS)

	content ,err  := io.ReadFile(fileName)
	if err != nil || content == "" {
		content, err = self.GetNext().Get(xmlS)
		if err == nil {
			writeNewContent(fileName,content)
		}
	}
	return content, err
}

func (self *FileCacheExecutor) GetNext() chain.Executor{
	return self.Next
}

func writeNewContent(fileName,content string){
	io.WriteFile(fileName,content)
}