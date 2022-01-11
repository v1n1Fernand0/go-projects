package contas

import cliente "banco/clientes"

type ContaPoupanca struct {
	Titular                    cliente.Titular
	NumeroConta, NumeroAgencia int
	saldo                      float64
}

func (conta *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor <= conta.saldo && valor > 0
	if podeSacar {
		conta.saldo -= valor
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente!"
}

func (conta *ContaPoupanca) Depositar(valor float64) string {

	podeDepositar := valor > 0
	if podeDepositar {
		conta.saldo += valor
		return "Valor depositado com sucesso!"
	} else {
		return "Depósito não realizado, informe um valor positivo!"
	}
}

func (conta *ContaPoupanca) Transferir(contaDestino ContaPoupanca, valor float64) string {
	podeTransferir := valor > 0 && valor <= conta.saldo
	if podeTransferir {
		conta.saldo -= valor
		contaDestino.Depositar(valor)
		return "Transferência realizada com sucesso!"
	}

	return "Ocorreu um erro durante a transferência! Verifique se você tem saldo para esta operação."
}

func (Conta *ContaPoupanca) ObterSaldo() float64 {
	return Conta.saldo
}
