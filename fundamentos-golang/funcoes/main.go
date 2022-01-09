package main

import "fmt"

//return nomeado
func calculoMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// função variática
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// função variática + tipo fixo
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

// função recursiva
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	//retorno nomeado
	fmt.Println(calculoMatematico(1, 2))

	//variáticas
	fmt.Println(soma(1, 2, 3))

	// variáticas + tipo fixo
	escrever("O valor é:", 1, 2, 3)

	// funções anônimas
	func() {
		fmt.Println("Olá")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Olá")

	retorno := func(texto string) string {
		return fmt.Sprintf("texto recebido %s", texto)
	}("Olá")

	fmt.Println(retorno)

	// recursiva
	// 1 1 2 3 5 8 13
	posicao := uint(10)
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(posicao))

}
