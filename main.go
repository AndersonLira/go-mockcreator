package main

import (
	"fmt"
    "net/http"
)

var config = GetConfig()

func main() {
	http.HandleFunc(config.GetContext(), HelloServer)
	fmt.Printf("Server started on %s port in %s context",config.GetPort(),config.GetContext());
    http.ListenAndServe(config.GetPort(), nil)
}

