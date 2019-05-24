package config

import(
	"strings"
)

func IsDelayedMethod(name string) bool{
	cfg := GetConfig()
	for _, s := range cfg.DelayMethods {
		if strings.HasPrefix(s,name){
			return true
		}
	}
	return false;
}




