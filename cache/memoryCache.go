package cache

import (
	"log"
	"strings"
	"sync"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"
	"github.com/andersonlira/goutils/watcher"
)

var memCache = make(map[string] string)
var listToClear = make(map[string]bool)
var once  sync.Once

type memoryCacheExecutor struct {
	next chain.Executor
}

func CreateMemoryCacheExecutor(next chain.Executor) (executor chain.Executor) {
	executor = &memoryCacheExecutor{next:next}
	once.Do(func(){
		wf := watcher.WatcherFile{
			Paths:[]string{config.GetConfig().PayloadFolder},
		}
		wf.Start()
		go func(){
			for {
				writeMapChanged(<-wf.FileChanged,true)
			}
		}()
	})
	return 
}

func (self memoryCacheExecutor) Get(xmlS string) (string,error) {
	manageListToClear()
	if self.next == nil {
		panic("next can not be nil. Use CreateMemoryCacheExecutor to create instance and pass executor reference")
	}

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

func (self *memoryCacheExecutor) GetNext() chain.Executor{
	return self.next
}

func manageCache(methodName string){
	if list, ok := config.GetConfig().ShouldClearCache(methodName); ok {
		for _,item := range list {
			for k := range getMap() {
				if strings.HasPrefix(k, item){
					log.Printf("%sRemoving memory cache: %s%s",ft.BLUE,k,ft.NONE)
					deleteFromMap(k)
				}
			}
		}
	}
}

func manageListToClear(){
	for fileName, _ := range getMapChanged() {
		key := strings.Replace(fileName,config.GetConfig().PayloadFolder,"",-1)
		for key2 := range getMap() {
			if strings.HasSuffix(fileName,key2){
				log.Printf("%sRemoving memory cache: %s because file has been changed.%s",ft.BLUE,key2,ft.NONE)
				deleteFromMapChanged(key)
				deleteFromMap(key2)
			}
		}
	}
}

var m = sync.Mutex{}

func deleteFromMap(key string){
	m.Lock()
	defer m.Unlock()
	delete(memCache,key)
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

func getMap() map[string]string {
	m.Lock()
	defer m.Unlock()
	return memCache
}


var mChanged = sync.Mutex{}


func deleteFromMapChanged(key string){
	mChanged.Lock()
	defer mChanged.Unlock()
	delete(listToClear,key)
}



func readMapChanged(key string) (bool,bool) {
	mChanged.Lock()
	defer mChanged.Unlock()
	r,ok := listToClear[key]
	return r,ok
}

func writeMapChanged(key string, value bool){
	mChanged.Lock()
	defer mChanged.Unlock()
	listToClear[key] = value
}

func getMapChanged() map[string]bool {
	mChanged.Lock()
	defer mChanged.Unlock()
	return listToClear
}

