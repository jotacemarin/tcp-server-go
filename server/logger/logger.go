package logger

import (
	"log"
	"os"
)

// Logger to write logs in file
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("tcp-server-go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	Logger = log.New(file, "TCP-SERVER-GO: ", log.LstdFlags)
}
