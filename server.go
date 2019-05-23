package main

import (
	"fmt"
	"io/ioutil"
    "net/http"

	"github.com/andersonlira/go-mockcreator/net"

)




func HelloServer(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	reqText := string(reqData)
    content 	:= net.Wsdl(reqText)
    fmt.Fprint(w, content)
}

