package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Println("o usuário que invocou esse método foi", u.nome)
}

func (u usuario) maiorIdade() bool {
	maior := u.idade >= 18
	fmt.Println("Usuário maior de idade?:", maior)
	return maior
}

func main() {
	usuario := usuario{"vinicius", 20}
	fmt.Println(usuario)
	usuario.salvar()
	usuario.maiorIdade()
}
