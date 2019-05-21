package main

import (
	"fmt"
	"sync"
)

type Config struct {
    Port int
}

func (c Config) getPort() string {
	return fmt.Sprintf(":%d",c.Port)
}



var (
	configInstance *Config
	once sync.Once
	port int = 8088
)

func GetConfig() *Config {
    once.Do(func() {
		port = port + 1
		configInstance = &Config{Port: port}
    })

	return configInstance
}

