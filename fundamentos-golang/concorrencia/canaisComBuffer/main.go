package main

import "fmt"

func main() {
	canal := make(chan string, 2) // espera receber 2 argumentos
	canal <- "Olá mundo"
	canal <- "Olá mundo2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
