package main

import (
	"fmt"
	"time"
)

func main() {

	// for
	i := 0
	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Carlos"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indece, nome := range nomes {
		fmt.Println(indece, nome)
	}

	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//loop infinito
	for {
		fmt.Println("Rodando")
		time.Sleep(time.Second)
	}
}
