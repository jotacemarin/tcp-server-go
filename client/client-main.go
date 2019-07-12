package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:6660")
	if err != nil {
		fmt.Printf("Error de conexion: \n%s", err)
	}

	defer connection.Close()
	fmt.Println("Conectado con el servidor")

	for {
		scanner := bufio.NewScanner(connection)

		for {

			ok := scanner.Scan()
			if !ok {
				fmt.Println("Reached EOF on server connection.")
				os.Exit(1)
				break
			}

			text := scanner.Text()
			switch {
			case text == "Hola.":
				sendFile(connection)
				return
			}

		}
	}
}

func sendMessage(message string, connection net.Conn) {
	connection.Write([]byte(message))
}

func sendFile(connection net.Conn) {
	file, err := os.Open("./client/dummyfile.dat")
	if err != nil {
		fmt.Printf("Error en lectura de archivo %s", err)
		os.Exit(1)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error obteniendo informacion del archivo %s", err)
	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)

	sendMessage(fileSize, connection)
	sendMessage(fileName, connection)

	sendBuffer := make([]byte, 1024)

	fmt.Println("Iniciando envio del archivo")

	for {
		_, err := file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}

	fmt.Println("Archivo enviado al servidor")
	return
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
