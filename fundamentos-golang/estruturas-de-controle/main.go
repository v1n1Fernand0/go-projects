package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}
func diaDaSemana2(numero int) string {
	//fallthrough serve para pular para a próxima condição
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}
func diaDaSemana3(numero int) string {
	//fallthrough serve para pular para a próxima condição
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
		fallthrough
	default:
		diaDaSemana = "valor inválido"
	}

	return diaDaSemana
}
func main() {
	numero := 10

	//IF
	//if comun
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	} else if numero < 0 {
		fmt.Println("Menor que 0")
	}

	//Switch
	var dia = diaDaSemana(3)
	fmt.Println(dia)

	var dia2 = diaDaSemana2(3)
	fmt.Println(dia2)
}
