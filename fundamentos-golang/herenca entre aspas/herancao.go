package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint16
	altura float32
}

type estudante struct {
	pessoa    // Mais próximo que o Go chega da herança, essa declaração copia os campos da struct mencionada.
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{nome: "vinicius", idade: 28, altura: 1.75}
	fmt.Println(p1);

	e1 := estudante{curso: "Análise de Sistemas", faculdade: "Estácio"}
	e1.nome = "Fernando"

	fmt.Println(e1)
}