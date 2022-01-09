package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrencia != paralilismo
	go escrever("Olá mundo!") //goroutines
	escrever("Tchau mundo!")
}
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
