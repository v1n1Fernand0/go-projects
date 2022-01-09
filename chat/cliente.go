package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type cliente struct {
	con      net.Conn
	apelido  string
	sala     *sala
	comandos chan<- comando
}

func (c *cliente) lerEntrada() {
	for {
		msg, err := bufio.NewReader(c.con).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/apelido":
			c.comandos <- comando{
				id:      CMD_APELIDO,
				cliente: c,
				args:    args,
			}
		case "/entrar":
			c.comandos <- comando{
				id:      CMD_ENTRAR,
				cliente: c,
				args:    args,
			}
		case "/salas":
			c.comandos <- comando{
				id:      CMD_SALAS,
				cliente: c,
			}
		case "/msg":
			c.comandos <- comando{
				id:      CMD_MSG,
				cliente: c,
				args:    args,
			}
		case "/sair":
			c.comandos <- comando{
				id:      CMD_SAIR,
				cliente: c,
			}
		default:
			c.err(fmt.Errorf("erro desconhecido: %s", cmd))
		}
	}
}

func (c *cliente) err(err error) {
	c.con.Write([]byte("Erro:" + err.Error() + "\n"))
}
func (c *cliente) msg(msg string) {
	c.con.Write([]byte("> " + msg + "\n"))
}
