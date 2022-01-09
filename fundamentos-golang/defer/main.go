package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}
func funcao2() {
	fmt.Println("Função 2")
}

func alunosAprovados(n1, n2 float32) bool {

	defer fmt.Println("Entrando na função para verificar se o aluno foi aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {

		return true
	} else {
		return false
	}
}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunosAprovados(7, 8))
}
