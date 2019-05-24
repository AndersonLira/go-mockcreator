package config

import(
	"strings"
)

func (c *Config) IsDelayedMethod(name string) bool{
	for _, s := range c.DelayMethods {
		if strings.HasPrefix(name, s){
			return true
		}
	}
	return false;
}




