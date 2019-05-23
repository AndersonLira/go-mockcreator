package cache

import (
	"log"
	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/xml"
)


var memCache = make(map[string] string)
type MemoryCacheExecutor struct {
	Next chain.Executor
}

func (self MemoryCacheExecutor) Get(xmlS string) (string,error) {

	fileName :=  xml.NameSugested(xmlS)
	var err error

	content ,ok  := memCache[fileName]
	if !ok || content == "" {
		content, err = self.GetNext().Get(xmlS)
		memCache[fileName] = content
	}else{
		log.Printf("Read from cache: %s",fileName)

	}
	return content, err
}

func (self *MemoryCacheExecutor) GetNext() chain.Executor{
	return self.Next
}

