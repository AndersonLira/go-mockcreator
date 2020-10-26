package net

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/andersonlira/go-mockcreator/chain"
	"github.com/andersonlira/go-mockcreator/config"
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/ft"
	"github.com/andersonlira/goutils/io"
	"github.com/go-xmlfmt/xmlfmt"

)

type WsdlExecutor struct {
	Next chain.Executor
}

func (self WsdlExecutor) Get(xmlS string) (string,error) {
	fileName := xml.ExtractXmlMethodName(xmlS)
	if self.GetNext() != nil {
		return self.GetNext().Get(xmlS)
	}else{
		log.Printf("%sRead from server: %s%s",ft.YELLOW,fileName,ft.NONE)
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
		log.Println("%sError on creating request object. %v \n %s ",ft.RED ,err.Error(),ft.NONE)
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
		log.Printf("%sError on dispatching request. %v \n %s",ft.RED,  err.Error(), ft.NONE)
		return "",err
	}

	responseData, _ := ioutil.ReadAll(res.Body)
	responseText := string(responseData)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		if cfg.ShowErrorServer {
			log.Printf("%s%s%s",ft.MARGENT,xmlRequest,ft.NONE)
			log.Println("^SOAP IN ---------v SOAP OUT-------")
			log.Printf("%s%s%s",ft.MARGENT,responseText,ft.NONE)
		}
		persistError(cfg.LogErrorFile,xmlRequest,responseText)
		return responseText, errors.New(fmt.Sprintf("Server Error status %d",res.StatusCode))
	}
	return responseText,nil
}

func persistError(fileName, request, response string) {
	if fileName == "" {
		return 
	}
	reqFormatted := xmlfmt.FormatXML(request, "\t", "  ")
	respFormatted := xmlfmt.FormatXML(response, "\t", "  ")
	io.AppendFile(fileName,fmt.Sprintf("<!-- %v -->\n",time.Now()))
	io.AppendFile(fileName,fmt.Sprintf("<!-- SOAP IN -->\n%s\n\n<!-- SOAP OUT -->\n%s\n\n",reqFormatted,respFormatted))
}