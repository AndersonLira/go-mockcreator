package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	Port int
	Context string
	URL string
	User string
	Password string
	PayloadFolder string
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
			url := os.Getenv("SERVICE_URL")

			user := os.Getenv("MC_USER")
			pass := os.Getenv("MC_PASS")
			if url == "" || user == "" || pass == "" {
				panic("Error. SERVICE_URL or MC_USER or MC_PASS environment variable is not setted")
			}

			pf := os.Getenv("MC_PAYLOAD_FOLDER")

			if pf == "" {
				pf = "payloads/"
			}

			configInstance = &Config{
				Port: port, 
				Context: context, 
				URL: url,
				User: user,
				Password: pass,
				PayloadFolder: pf,
			}
    })

	return configInstance
}

