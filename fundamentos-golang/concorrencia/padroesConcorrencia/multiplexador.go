package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo 1"), escrever("Olá mundo 2"))
	fmt.Println(<-canal)
}

func multiplexar(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		select {
		case mensagem := <-ch1:
			ch <- mensagem
		case mensagem := <-ch2:
			ch <- mensagem
		}
	}()

	return ch
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	return canal
}
