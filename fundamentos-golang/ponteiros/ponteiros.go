package main

import "fmt"

func main() {
	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	variavel++
	fmt.Println(variavel, variavel2)

	// ponteiro é uma referência de memória
	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // & aponta para o endereço de memória da variável em questão
	fmt.Println(variavel3, *ponteiro) //desreferenciação
}