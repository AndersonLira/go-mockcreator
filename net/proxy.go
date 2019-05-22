package net

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/andersonlira/go-mockcreator/config"

)

func Wsdl(xmlRequest string) {
	// wsdl service url
	cfg := config.GetConfig()
	url := fmt.Sprintf("%s",
		cfg.URL,
	)

	// payload
	payload := []byte(strings.TrimSpace(xmlRequest))

	httpMethod := "POST"

	// soap action
	soapAction := ExtractXmlMethodName(xmlRequest)

	// authorization credentials
	username := cfg.User
	password := cfg.Password

	log.Println("-> Preparing the request")

	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
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

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	responseText := string(responseData)
	fmt.Printf("%s",responseText)
}