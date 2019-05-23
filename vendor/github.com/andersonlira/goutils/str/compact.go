package str

import (
	"strings"
)

//Compact remove from string spaces, tabs and end of lines
func Compact(s string) (r string) {
 	r = strings.Replace(s, " ", "", -1)
 	r = strings.Replace(r, "\t", "", -1)
 	r = strings.Replace(r, "\n", "", -1)
	return
}