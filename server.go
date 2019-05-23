package main

import (
	"fmt"
	"io/ioutil"
    "net/http"

	"github.com/andersonlira/go-mockcreator/cache"
	"github.com/andersonlira/go-mockcreator/net"

)

var fileExecutor = cache.FileCacheExecutor{}
var wsdlExecutor = net.WsdlExecutor{}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	reqText := string(reqData)
	fileExecutor.Next = &wsdlExecutor
    content, _ := fileExecutor.Get(reqText)
    fmt.Fprint(w, content)
}

