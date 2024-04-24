package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("vevsdev@gmail.com")
	fmt.Println(err)
}
