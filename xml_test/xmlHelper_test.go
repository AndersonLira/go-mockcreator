package xml_test

import (
	"testing"
	
	"github.com/andersonlira/go-mockcreator/xml"
	"github.com/andersonlira/goutils/str"
)

var (
	xmlA = 	 `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:prl="http://www.andersonlira.com/service/prl">
	<soapenv:Header/>
	<soapenv:Body>
	   <prl:getUser>
		  <identifier>contato@andersonlira.com</identifier>
	   </prl:getUser>
	</soapenv:Body>
 	</soapenv:Envelope>	
	`
	bodyA = `
		<prl:getUser>
			<identifier>contato@andersonlira.com</identifier>
		</prl:getUser>
	`

	xmlB = `<?xml version="1.0" encoding="UTF-8"?>
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

	bodyB = `
		<ns2:getAccountsResponse xmlns:ns2="http://www.vodafone.pt/myvdf/core/service/prl">
			<return>
				<assetId>1-3IT9</assetId>
				<accountId>300031330</accountId>
				<accountStatus>ACTIVE</accountStatus>
				<hasAccountPack>false</hasAccountPack>
				<deactivationDate xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:nil="true" />
			</return>
		</ns2:getAccountsResponse>

	`
)

func TestExtractXmlMethodName(t *testing.T) {
	expectedName := "getUser"
	methodName := xml.ExtractXmlMethodName(xmlA)

	if expectedName != methodName {
		t.Errorf("Method name should be %s but %s",expectedName,methodName)
	}


	
	expectedName = "getAccountsResponse"
	methodName = xml.ExtractXmlMethodName(xmlB)

	if expectedName != methodName {
		t.Errorf("Method name should be %s but %s",expectedName,methodName)
	}

	
}

func TestExtractXmlBody(t *testing.T){
	extractedBodyA := xml.ExtractXmlBody(xmlA)

	if str.Compact(bodyA) != str.Compact(extractedBodyA) {
		t.Errorf("Body should be %s but %s",bodyA,extractedBodyA)
	}

	extractedBodyB := xml.ExtractXmlBody(xmlB)

	if str.Compact(bodyB) != str.Compact(extractedBodyB) {
		t.Errorf("Body should be %s but %s",bodyB, extractedBodyB)
	}

}

func TestNameSugested(t *testing.T){
	sugestion := xml.NameSugested(xmlA)
	expectedName := "getUser4029099683.xml"


	if sugestion != expectedName {
		t.Errorf("Name should be %s but %s",expectedName, sugestion)
	}

}
