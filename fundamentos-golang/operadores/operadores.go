package main

import "fmt"

func main() {
	//aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	var soma2 = numero1 + int16(numero2)
	fmt.Println(soma2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Print(verdadeiro && falso)
	fmt.Print(verdadeiro || falso)
	fmt.Print(!verdadeiro)

	//UNÁRIOS
	numero:= 10
	numero++
	numero+=15
	numero--
	numero -= 10
	numero /= 10
	numero *= 10
	fmt.Println(numero)

}
