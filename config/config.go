package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

var DefaultAvatar = "/front/default-avatar.jpg"

var Conf = NewConfig("./config/config.json")

type Config struct {
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	MySQL struct {
		GormDNS string `json:"gorm_dns"`
	}
}

func NewConfig(filePath string) *Config {
	var config Config
	data, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatal(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		logrus.Fatal(err)

	}
	return &config
}
