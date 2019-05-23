package net

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"

)

type WsdlExecutor struct {
	Next chain.Executor
}

func (self WsdlExecutor) Get(xmlS string) (string,error) {
	fileName := xml.ExtractXmlMethodName(xmlS)
	if self.GetNext() != nil {
		return self.GetNext().Get(xmlS)
	}else{
		log.Printf("%sRead from cache: %s%s",ft.YELLOW,fileName,ft.NONE)
	}
	return wsdl(xmlS)
}

func (self *WsdlExecutor) GetNext() chain.Executor{
	return self.Next
}


func wsdl(xmlRequest string) (string,error){
	// wsdl service url
	cfg := config.GetConfig()
	url := fmt.Sprintf("%s",
		cfg.URL,
	)

	// payload
	payload := []byte(strings.TrimSpace(xmlRequest))

	httpMethod := "POST"

	// soap action
	soapAction := xml.ExtractXmlMethodName(xmlRequest)

	// authorization credentials
	username := cfg.User
	password := cfg.Password


	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return "", err
	}

	// set the content type header, as well as the oter required headers
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)
	req.SetBasicAuth(username, password)

	// prepare the client request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}


	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return "",err
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	responseText := string(responseData)
	return responseText,nil
}