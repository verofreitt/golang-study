package main

import "fmt"

func main() {
	var nome string = "Variavel 1"
	nome2 := "Variavel 2"
	fmt.Println(nome)
	fmt.Println(nome2)

	var (
		nome3 string = "Variavel 3"
		nome4 string = "Variavel 4"
	)
	fmt.Println(nome3, nome4)

	nome5, nome6 := "Variavel 5", "Variavel 6"
	fmt.Println(nome6, nome5)

	const constante string = "Constante"
	fmt.Println(constante)

	nome5, nome6 = nome6, nome5
	fmt.Println(nome5, nome6)
}
