package main

import (
	"log"
	"net"
)

func main() {
	s := novoServidor()
	go s.executar()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("erro ao iniciar o servidor: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("servidor iniciado na porta :8888")

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Printf("não foi possível estabelecer a conexão: %s", err.Error())
			continue
		}

		go s.novoCliente(con)
	}
}
