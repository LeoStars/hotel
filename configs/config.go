package configs

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	SSLMode string `json:"ssl_mode"`
	Host string `json:"host"`
	DBName string `json:"db_name"`
	User string `json:"user"`
	Password string `json:"password"`
}


func NewConfig(conf string) (*Config, error) {
	config := new(Config)
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return config, nil
}