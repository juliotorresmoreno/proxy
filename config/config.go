package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Address   string `json:"address"`
	Remote    string `json:"remote"`
	Debug     bool   `json:"debug"`
	Mode      string `json:"mode"`
	Logging   bool   `json:"logging"`
	Logs      string `json:"logs"`
	Tunneling struct {
		BufferSize int `json:"buffersize"`
	} `json:"tunneling"`
}

func GetConfig() Config {
	r := Config{
		Address: ":8080",
	}
	c, err := ioutil.ReadFile("./config.conf")
	if err != nil {
		return r
	}
	json.Unmarshal(c, &r)
	return r
}
