package main

import (
	"fmt"
	"net/http"
	
	"github.com/andersonlira/go-mockcreator/config"
)

var cfg = config.GetConfig()

func main() {
	http.HandleFunc(cfg.GetContext(), HelloServer)
	fmt.Printf("Server started on %s port in %s context",cfg.GetPort(),cfg.GetContext());
    http.ListenAndServe(cfg.GetPort(), nil)
}

