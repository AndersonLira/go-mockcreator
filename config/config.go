package config

import (
	"fmt"
	"encoding/json"
	"os"
	"sync"

	"github.com/andersonlira/goutils/io"
)

type Config struct {
	Port int `json:"port"`
	Context string
	URL string
	User string
	Password string
	PayloadFolder string

	ReturnDelay int `json:"returnDelay"`
	DelayMethods []string `json:"delayMethods"`
	ShowErrorServer bool `json:"showErrorServer"`
	WorkAsProxy bool `json:"workingAsProxy"`

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
	defaultPort int = 8088
)

//GetConfig returns config for application. It is a singleton object
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
				Port: defaultPort, 
				Context: context, 
				URL: url,
				User: user,
				Password: pass,
				PayloadFolder: pf,
			}
			readFileConf(configInstance)
    })

	return configInstance
}

func readFileConf(cfg *Config){
	s,_ := io.ReadFile("config.json")
	json.Unmarshal([]byte(s), cfg)
}

