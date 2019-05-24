package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/andersonlira/go-mockcreator/cache"
	"github.com/andersonlira/go-mockcreator/net"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"

)

var memExecutor = cache.MemoryCacheExecutor{}
var fileExecutor = cache.FileCacheExecutor{}
var wsdlExecutor = net.WsdlExecutor{}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	reqText := string(reqData)
	memExecutor.Next = &fileExecutor
	fileExecutor.Next = &wsdlExecutor
	content, _ := memExecutor.Get(reqText)
	if cfg.ReturnDelay > 0 {
		log.Printf("%s Sleeping %d milliseconds for %s request%s",ft.GREEN,cfg.ReturnDelay,xml.ExtractXmlMethodName(reqText),ft.NONE )
		time.Sleep(time.Duration(cfg.ReturnDelay) * time.Millisecond)
	}
    fmt.Fprint(w, content)
}

