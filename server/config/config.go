package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct for model Configurations
type Config struct {
	Addr              string
	DestinationFolder string
	MessageConnection string
}

// Configurations for server
var Configurations Config

// LoadConfig : func
func init() {
	configurationFile, err := os.Open("/opt/tcp-server-go/conf.json")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer configurationFile.Close()
	jsonParser := json.NewDecoder(configurationFile)
	jsonParser.Decode(&Configurations)
}
