package handlers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// HandlerConnection : administrador de la conexion
func HandlerConnection(connection net.Conn) {
	remoteAddr := connection.RemoteAddr().String()
	log.Printf("Cliente conectado desde %s\n", remoteAddr)

	sendMessage("\r\nHola.\r\n", connection)

	scanner := bufio.NewScanner(connection)
	handleMessage(scanner.Text(), connection)

	log.Println("Cliente en " + remoteAddr + " se ha desconectado.")
}

// manejador de los mensajes
func handleMessage(message string, connection net.Conn) {
	folderFile := fmt.Sprint("/tmp/tcp-server-go/" + time.Now().Format("2006-01-02") + "/")

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := fmt.Sprintf("%s %s", time.Now().Format(time.StampMilli), strings.Trim(string(bufferFileName), ":"))

	if _, err := os.Stat(folderFile); os.IsNotExist(err) {
		os.MkdirAll(folderFile, os.ModePerm)
	}

	newFile, err := os.Create(folderFile + fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer newFile.Close()

	var receivedBytes int64

	for {

		if (fileSize - receivedBytes) < 1024 {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+1024)-fileSize))
			break
		}
		io.CopyN(newFile, connection, 1024)
		receivedBytes += 1024

	}

	log.Println("Archivo recivido!")
}

// envia mensaje hacia el cliente
func sendMessage(message string, connection net.Conn) {
	connection.Write([]byte(message))
}
