package cache

import (
	"log"
	"strings"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"
)


var memCache = make(map[string] string)
type MemoryCacheExecutor struct {
	Next chain.Executor
}

func (self MemoryCacheExecutor) Get(xmlS string) (string,error) {

	fileName :=  xml.NameSugested(xmlS)
	methodName := xml.ExtractXmlMethodName(xmlS)
	var err error

	if config.GetConfig().IsCacheEvict(fileName) {
		return self.GetNext().Get(xmlS)
	}

	content ,ok  := memCache[fileName]
	if !ok || content == "" {
		content, err = self.GetNext().Get(xmlS)
		if err == nil {
			memCache[fileName] = content
			manageCache(methodName)
		}
	}else{
		log.Printf("Read from cache: %s",fileName)

	}
	return content, err
}

func (self *MemoryCacheExecutor) GetNext() chain.Executor{
	return self.Next
}

func manageCache(methodName string){
	if list, ok := config.GetConfig().ShouldClearCache(methodName); ok {
		for _,item := range list {
			for k := range memCache {
				if strings.HasPrefix(k, item){
					log.Printf("%sRemoving memory cache: %s%s",ft.BLUE,k,ft.NONE)
					delete(memCache,k)
				}
			}
		}
	}
}
