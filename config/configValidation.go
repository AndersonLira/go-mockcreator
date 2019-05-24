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

func (c *Config) ShouldClearCache(name string ) ([]string, bool){
	list, ok := c.ClearCache[name]
	return list,ok
}


