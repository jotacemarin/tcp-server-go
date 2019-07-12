package handlers

import (
	"bufio"
	"fmt"
	"net"
)

// HandlerConnection : administrador de la conexion
func HandlerConnection(connection net.Conn) {
	remoteAddr := connection.RemoteAddr().String()
	fmt.Printf("Cliente conectado desde %s\n", remoteAddr)

	sendMessage("\r\nHola.\r\n", connection)

	scanner := bufio.NewScanner(connection)

	for {

		ok := scanner.Scan()
		if !ok {
			break
		}

		handleMessage(scanner.Text(), connection)

	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

// manejador de los mensajes
func handleMessage(message string, conn net.Conn) {
	fmt.Println("> " + message)
}

// envia mensaje hacia el cliente
func sendMessage(message string, connection net.Conn) {
	connection.Write([]byte(message))
}
