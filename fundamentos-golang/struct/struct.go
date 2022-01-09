package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint32
}
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {

	var u usuario
	u.nome = "Vinicius"
	u.idade = 21
	fmt.Println(u)

	endereco := endereco{"rua carla", 12}
	usuario2 := usuario{"vinicius", 28, endereco}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "vinicius"}
	fmt.Println(usuario3)
}
