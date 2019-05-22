package net

import (
	"regexp"
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

func ExtractBody(xml string) string {
	panic("Not implemented yet")
}
