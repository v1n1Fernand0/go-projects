package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	salas    map[string]*sala
	comandos chan comando
}

func novoServidor() *server {
	return &server{
		salas:    make(map[string]*sala),
		comandos: make(chan comando),
	}
}

func (s *server) executar() {
	for cmd := range s.comandos {
		switch cmd.id {
		case CMD_APELIDO:
			s.apelido(cmd.cliente, cmd.args)
		case CMD_ENTRAR:
			s.entrar(cmd.cliente, cmd.args)
		case CMD_MSG:
			s.msg(cmd.cliente, cmd.args)
		case CMD_SALAS:
			s.listarSalas(cmd.cliente, cmd.args)
		case CMD_SAIR:
			s.sair(cmd.cliente, cmd.args)
		}
	}
}

func (s *server) novoCliente(conexao net.Conn) {
	log.Printf("novo cliente conectado: %s", conexao.RemoteAddr().String())

	c := &cliente{
		con:      conexao,
		apelido:  "anônimo",
		comandos: s.comandos,
	}

	c.lerEntrada()
}

func (s *server) apelido(c *cliente, args []string) {
	if len(args) < 2 {
		c.msg("um apelido é necessário. use o comando /apelido APELIDO")
		return
	}

	c.apelido = args[1]
	c.msg(fmt.Sprintf("ola, eu vou chama-lo de %s", c.apelido))
}

func (s *server) entrar(c *cliente, args []string) {
	if len(args) < 2 {
		c.msg("o nome da sala é necessário. use o comando /entrar NOME_SALA")
		return
	}

	nomeSala := args[1]

	r, ok := s.salas[nomeSala]
	if !ok {
		r = &sala{
			nome:    nomeSala,
			membros: make(map[net.Addr]*cliente),
		}
		s.salas[nomeSala] = r
	}

	r.membros[c.con.RemoteAddr()] = c

	s.sairSalaAtual(c)
	c.sala = r
	r.transmissao(c, fmt.Sprintf("%s entrou na sala ", c.apelido))
	c.msg(fmt.Sprintf("bem vindo a sala %s", r.nome))
}

func (s *server) msg(c *cliente, args []string) {
	if len(args) < 2 {
		c.msg("uma mensagem é necessária, use: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.sala.transmissao(c, c.apelido+": "+msg)
}

func (s *server) listarSalas(c *cliente, args []string) {
	var salas []string
	for name := range s.salas {
		salas = append(salas, name)
	}

	c.msg(fmt.Sprintf("salas disponíveis: %s", strings.Join(salas, ", ")))
}

func (s *server) sair(c *cliente, args []string) {
	log.Printf("cliente desconectado: %s", c.con.RemoteAddr().String())
	s.sairSalaAtual(c)
	c.msg("já vai? que triste! :(")
	c.con.Close()
}

func (s *server) sairSalaAtual(c *cliente) {
	if c.sala != nil {
		oldRoom := s.salas[c.sala.nome]
		delete(s.salas[c.sala.nome].membros, c.con.RemoteAddr())
		oldRoom.transmissao(c, fmt.Sprintf("%s saiu da sala", c.apelido))
	}
}
