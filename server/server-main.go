package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"./handlers"
)

var addr = flag.String("addr", "localhost", "La direccion para escuchar los clientes; por defecto es \"\".")
var port = flag.Int("port", 6660, "El puerto por defecto es 6660.")
var protocol = flag.String("protocol", "tcp", "El protocolo de la conneccion es")

func main() {
	flag.Parse()

	server, err := net.Listen(*protocol, fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		log.Printf("Error en servidor \n%s", err)
	}
	defer server.Close()

	log.Printf("Servidor iniciado en %s:%d.\nEsperando de conexiones.\n", *addr, *port)

	for {

		connection, err := server.Accept()
		if err != nil {
			log.Printf("Error con la conexion: \n%s", err)
			return
		}

		handlers.HandlerConnection(connection)
	}

}
