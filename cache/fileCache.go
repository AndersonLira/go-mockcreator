package cache

import (
	"log"
	"io/ioutil"
	"strings"
	"os"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/ft"
)


type FileCacheExecutor struct {
	Next chain.Executor
}

var payloadFolder = config.GetConfig().PayloadFolder + "/"

func (self FileCacheExecutor) Get(xmlS string) (string,error) {

	fileName := payloadFolder + xml.NameSugested(xmlS)
	methodName := xml.ExtractXmlMethodName(xmlS)

	content ,err  := io.ReadFile(fileName)
	if err != nil || content == "" {
		content, err = self.GetNext().Get(xmlS)
		if err == nil {
			writeNewContent(fileName,content)
		}
	}else{
		log.Printf("Read from file: %s",fileName)

	}
	manageFileCache(methodName)
	return content, err
}

func (self *FileCacheExecutor) GetNext() chain.Executor{
	return self.Next
}

func writeNewContent(fileName,content string){
	io.WriteFile(fileName,content)
}

func manageFileCache(methodName string){
	if list, ok := config.GetConfig().ShouldClearCache(methodName); ok {
		files, _ := ioutil.ReadDir(payloadFolder)
		for _,f := range files {
			for _,item := range list {
				if strings.HasPrefix(f.Name(), item){
					log.Printf("%sRemoving file: %s%s",ft.BLUE,payloadFolder + f.Name(),ft.NONE)
					os.Remove(payloadFolder + f.Name())
				}
			}
		}
	}
}