package xml

import (
	"fmt"
	"regexp"

	"github.com/andersonlira/goutils/str"
)

//ExtractXmlMethodName return method name from a xml string. 
//It works for prl and ns2 for a while and return 'MethodUnknown' if no match
func ExtractXmlMethodName(xml string) string {
	regex := *regexp.MustCompile(`(?s)<ns2:([a-zA-Z0-9]+)`)
    res := regex.FindAllStringSubmatch(xml, -1)
	if len(res) > 0 {
		return res[0][1]
	}
	regex = *regexp.MustCompile(`(?s)<prl:([a-zA-Z0-9]+)`)
    res = regex.FindAllStringSubmatch(xml, -1)
	if len(res) > 0 {
		return res[0][1]
	}
	return "MethodUnknown"
}

//ExtractXmlBody return body content from giving xml
func ExtractXmlBody(xml string) string {
	regex := *regexp.MustCompile(`(?s)<.*Body>(.*)</[a-zA-Z0-9]+:Body>`)
    res := regex.FindAllStringSubmatch(xml, -1)
	if len(res) > 0 {
		return res[0][1]
	}
	return xml
}

func NameSugested(xml string) string {
	return fmt.Sprintf("%s%d.xml",ExtractXmlMethodName(xml),str.Hash(ExtractXmlBody(xml)))
}
