package cache

import (
	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
)


var memCache = make(map[string] string)
type MemoryCacheExecutor struct {
	Next chain.Executor
}

func (self MemoryCacheExecutor) Get(xmlS string) (string,error) {
	cfg := config.GetConfig()

	fileName := cfg.PayloadFolder + "/" + xml.NameSugested(xmlS)
	var err error

	content ,ok  := memCache[fileName]
	if !ok || content == "" {
		content, err = self.GetNext().Get(xmlS)
		memCache[fileName] = content
	}
	return content, err
}

func (self *MemoryCacheExecutor) GetNext() chain.Executor{
	return self.Next
}

