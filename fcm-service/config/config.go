package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Config

func init() {
	file, err := os.Open("config/conf.json")
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	config = &v
}

// GetConfig :
func GetConfig() *Config {
	return config
}

// Config :
type Config struct {
	Port         int    `json:"port"`
	Env          string `json:"env"`
	GcmEndpoint  string `json:"gcm_endpoint"`
	GcmServerKey string `json:"gcm_server_key"`
}
