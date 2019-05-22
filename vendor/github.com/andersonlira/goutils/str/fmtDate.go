package str

import (
	"strings"
	"time"
)
//FmtDate formats date with basic java pattern instead golang pattern
func FmtDate(format string, t time.Time) string{
    //2006-01-02T15:04:05Z07:00
	newFormat := strings.ReplaceAll(format,"dd","02")
	newFormat = strings.ReplaceAll(newFormat,"MMM","Jan")
	newFormat = strings.ReplaceAll(newFormat,"MM","01")
	newFormat = strings.ReplaceAll(newFormat,"yyyy","2006")
	newFormat = strings.ReplaceAll(newFormat,"yy","06")
	newFormat = strings.ReplaceAll(newFormat,"HH","15")
	newFormat = strings.ReplaceAll(newFormat,"mm","04")
	newFormat = strings.ReplaceAll(newFormat,"ss","05")

	
	return t.Format(newFormat)
}