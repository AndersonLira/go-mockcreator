package main

import (
	"log"
	"net/http"
	"fmt"
	
	"github.com/andersonlira/go-mockcreator/config"
)

var cfg = config.GetConfig()

var logo = `
 _____                                                  
/ ____|                                                 
| |  __  ___                                             
| | |_ |/ _ \                                            
| |__| | (_) |                                           
\_____| \___/      _                       _             
|  \/  |          | |                     | |            
| \  / | ___   ___| | _____ _ __ ___  __ _| |_ ___  _ __ 
| |\/| |/ _ \ / __| |/ / __| '__/ _ \/ _  | __/ _  | '__|
| |  | | (_) | (__|   < (__| | |  __/ (_| | || (_) | |   
|_|  |_|\___/ \___|_|\_\___|_|  \___|\__,_|\__\___/|_| V1.3
`

func main() {
	fmt.Println(logo)
	http.HandleFunc(cfg.GetContext(), HelloServer)
	log.Printf("Server started on %s port in %s context",cfg.GetPort(),cfg.GetContext());
    http.ListenAndServe(cfg.GetPort(), nil)
}

