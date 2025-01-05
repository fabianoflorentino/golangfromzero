package main

import (
	"github.com/fabianoflorentino/golangdozero/pkg/fundamentos_da_linguagem"
)

func main() {
	println("Aprenda Golang do Zero!")

	println("\n5. Pacotes Exeternos\n")
	println("fabianoflorentino@outlook.com: ", fundamentos_da_linguagem.PacoteExterno("fabianoflorentino@outlook.com"))
	println("invalido: ", fundamentos_da_linguagem.PacoteExterno("invalido"))

	println("\n6. Variáveis")
	println(fundamentos_da_linguagem.Variaveis())
}
