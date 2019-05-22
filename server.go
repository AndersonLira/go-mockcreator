package main

import (
    "fmt"
    "net/http"

	"github.com/andersonlira/go-mockcreator/net"

)




func HelloServer(w http.ResponseWriter, r *http.Request) {
    content 	:= net.Wsdl(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:prl="http://www.abcdefghhijkmln.pt/">
		<soapenv:Header/>
		<soapenv:Body>
		<prl:getUser>
			<identifier>contato@andersonlira.com</identifier>
		</prl:getUser>
		</soapenv:Body>
	</soapenv:Envelope>`)
    fmt.Fprint(w, content)
}

