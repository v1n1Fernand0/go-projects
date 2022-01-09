package main

import "net"

type sala struct {
	nome    string
	membros map[net.Addr]*cliente
}

func (s *sala) transmissao(sender *cliente, msg string) {
	for addr, m := range s.membros {
		if sender.con.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
