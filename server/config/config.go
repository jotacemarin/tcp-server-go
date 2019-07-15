package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Addr              string
	DestinationFolder string
	MessageConnection string
}

// Configurations for server
var Configurations Config

// LoadConfig : func
func init() (Config, errConfig error) {
	configurationFile, err := os.Open("/opt/tcp-server-go/conf.json")
	if err != nil {
		return nil, err
	}
	defer configurationFile.Close()
	jsonParser := json.NewDecoder(configurationFile)
	jsonParser.Decode(&Configurations)
	return configurationFile, nil
}
