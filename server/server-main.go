package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"./handlers"
)

var addr = flag.String("addr", "localhost", "La direccion para escuchar los clientes; por defecto es \"\".")
var port = flag.Int("port", 6660, "El puerto por defecto es 6660.")
var protocol = flag.String("protocol", "tcp", "El protocolo de la conneccion es")

func main() {
	flag.Parse()

	server, err := net.Listen(*protocol, fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		fmt.Printf("Error en servidor \n%s", err)
	}
	defer server.Close()

	fmt.Printf("Servidor iniciado en %s:%d.\nEsperando de conexiones.\n", *addr, *port)

	for {

		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		handlers.HandlerConnection(connection)
	}

}
