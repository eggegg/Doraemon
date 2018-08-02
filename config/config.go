package config

import (
	"io/ioutil"
	"encoding/json"

	"github.com/labstack/gommon/log"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`
	Mongodb struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"mongodb"`
	Runner struct{
		Duration int `json:"duration"`
	} `json:"runner"`
	LogLevel string `json:"loglevel"`
}


// FromFile returns a configuration parsed from the given file
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetLogLvl(debuglvl string) log.Lvl {
	//DEBUG INFO WARN ERROR OFF
	switch debuglvl {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OF":
		return log.OFF
	}

	return log.WARN
}