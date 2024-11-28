package configuration

import (
	"os"

	"bstz.it/rest-api/utils"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Database struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		SslMode      string `yaml:"ssl_mode"`
		DatabaseName string `yaml:"database_name"`
	} `yaml:"database"`
}

var Config *AppConfig

func LoadConfig() {
	if Config == nil {
		Config = &AppConfig{}
		data, err := os.ReadFile("config.yaml")
		utils.PanicOnError(err, "Error reading config file")

		err = yaml.Unmarshal(data, &Config)
		utils.PanicOnError(err, "Error parsing config file")
	}
}
