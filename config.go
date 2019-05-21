package main

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	Port int
	Context string
}

func (c Config) GetPort() string {
	return fmt.Sprintf(":%d",c.Port)
}

func (c Config) GetContext() string {
	return fmt.Sprintf("/%s",c.Context)
}


var (
	configInstance *Config
	once sync.Once
	port int = 8088
)

func GetConfig() *Config {
    once.Do(func() {
		context := os.Getenv("SERVER_CONTEXT")
		configInstance = &Config{Port: port, Context: context}
    })

	return configInstance
}

