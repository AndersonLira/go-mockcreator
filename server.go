package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/andersonlira/go-mockcreator/cache"
	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/net"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"

)


func HelloServer(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	reqText := string(reqData)
	
	fileExecutor := cache.FileCacheExecutor{}
	memExecutor := cache.CreateMemoryCacheExecutor(&fileExecutor)
	wsdlExecutor := net.WsdlExecutor{}

	fileExecutor.Next = &wsdlExecutor
	var executor  chain.Executor
	if cfg.WorkAsProxy {
		executor = &wsdlExecutor
	}else{
		executor = memExecutor
	}
	content, err := executor.Get(reqText)
	if err != nil {
		http.Error(w,content, http.StatusInternalServerError)
		if cfg.LoopWhenErrorInterval > 30 {
			worker := Worker{executor, reqText, cfg.LoopWhenErrorInterval}
			worker.Run()
		}
	}
	methodName := xml.ExtractXmlMethodName(reqText)
	if cfg.ReturnDelay > 0 && cfg.IsDelayedMethod(methodName) {
		log.Printf("%s Sleeping %d milliseconds for %s request%s",ft.GREEN,cfg.ReturnDelay,methodName,ft.NONE )
		time.Sleep(time.Duration(cfg.ReturnDelay) * time.Millisecond)
	}
	if cfg.LogRequestBody && err == nil {
		log.Printf("Request Body %s \n\n",methodName)
		log.Println(reqText)
		fmt.Println("")
	}
	if cfg.LogResponseBody && err == nil {
		log.Printf("Response Body %s\n\n",methodName)
		log.Println(content)
		fmt.Println("")
	}

	for k,v := range cfg.ManipulateData {
		content = strings.ReplaceAll(content,k,v)
	}

	fmt.Fprint(w, content)
}

