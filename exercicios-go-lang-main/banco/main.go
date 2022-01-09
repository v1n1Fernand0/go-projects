package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaCorrente{}
	contaDoDenis.Depositar(300)
	PagarBoleto(&contaDoDenis, 150)
	fmt.Println(contaDoDenis)
}
