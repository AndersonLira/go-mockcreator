package cache

import (
	"log"
	"strings"
	"sync"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"
)

var m = sync.Mutex{}
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

	content ,ok  := readMap(fileName)
	if !ok || content == "" {
		content, err = self.GetNext().Get(xmlS)
		if err == nil {
			writeMap(fileName, content)
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

func readMap(key string) (string,bool) {
	m.Lock()
	defer m.Unlock()
	r,ok := memCache[key]
	return r,ok
}

func writeMap(key, value string){
	m.Lock()
	defer m.Unlock()
	memCache[key] = value
}
