package net_test

import (
	"testing"
	"github.com/andersonlira/go-mockcreator/net"
)

func TestExtractMethodName(t *testing.T) {
	expectedName := "getUser"
	xml := `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:prl="http://www.andersonlira.com/service/prl">
	<soapenv:Header/>
	<soapenv:Body>
	   <prl:getUser>
		  <identifier>contato@andersonlira.com</identifier>
	   </prl:getUser>
	</soapenv:Body>
 	</soapenv:Envelope>	
	`
	methodName := net.ExtractMethodName(xml)

	if expectedName != methodName {
		t.Errorf("Method name should be %s but %s",expectedName,methodName)
	}


	xml = `<?xml version="1.0" encoding="UTF-8"?>
	<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	   <S:Header />
	   <S:Body>
		  <ns2:getAccountsResponse xmlns:ns2="http://www.vodafone.pt/myvdf/core/service/prl">
			 <return>
				<assetId>1-3IT9</assetId>
				<accountId>300031330</accountId>
				<accountStatus>ACTIVE</accountStatus>
				<hasAccountPack>false</hasAccountPack>
				<deactivationDate xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:nil="true" />
			 </return>
		  </ns2:getAccountsResponse>
	   </S:Body>
	</S:Envelope>`
	
	expectedName = "getAccountsResponse"
	methodName = net.ExtractMethodName(xml)

	if expectedName != methodName {
		t.Errorf("Method name should be %s but %s",expectedName,methodName)
	}

	
}