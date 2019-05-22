package main

import (
    "fmt"
    "net/http"

    "github.com/andersonlira/goutils/io"

)




func HelloServer(w http.ResponseWriter, r *http.Request) {
    content, _ := io.ReadFile("D:\\programs\\mockcreator\\payloads\\geBillingInfoStatic.xml")
    fmt.Fprint(w, content)
}

