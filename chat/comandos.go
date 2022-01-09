package main

type comandoId int

const (
	CMD_APELIDO comandoId = iota
	CMD_ENTRAR
	CMD_SALAS
	CMD_MSG
	CMD_SAIR
)

type comando struct {
	id      comandoId
	cliente *cliente
	args    []string
}
