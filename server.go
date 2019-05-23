package main

import (
	"fmt"
	"io/ioutil"
    "net/http"

	"github.com/andersonlira/go-mockcreator/net"

)


var executor = net.WsdlExecutor{}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	reqText := string(reqData)
    content, _ := executor.Get(reqText)
    fmt.Fprint(w, content)
}

