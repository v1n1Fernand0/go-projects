package main

import (
	"log"
	"os"

	app "github.com/v1n1Fernand0/fundamentos-golang/fundamentos-golang/app/application"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
