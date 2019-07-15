package handlers

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// HandlerConnection : administrador de la conexion
func HandlerConnection(connection net.Conn) {
	remoteAddr := connection.RemoteAddr().String()
	fmt.Printf("Cliente conectado desde %s\n", remoteAddr)

	sendMessage("\r\nHola.\r\n", connection)

	scanner := bufio.NewScanner(connection)
	handleMessage(scanner.Text(), connection)

	fmt.Println("Cliente en " + remoteAddr + " se ha desconectado.")
}

// manejador de los mensajes
func handleMessage(message string, connection net.Conn) {

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	if _, err := os.Stat("/tmp/tcp-server-go/"); os.IsNotExist(err) {
		os.MkdirAll("/tmp/tcp-server-go/", os.ModePerm)
	}

	newFile, err := os.Create("/tmp/tcp-server-go/" + fileName)
	if err != nil {
		fmt.Println(err)
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

	fmt.Println("Archivo recivido!")
}

// envia mensaje hacia el cliente
func sendMessage(message string, connection net.Conn) {
	connection.Write([]byte(message))
}
