package main

import "fmt"

func main() {

	var x, y int
	var operacao string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&x)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&y)

	fmt.Println("Escolha a operação desejada (+, -, *, /): ")
    fmt.Scanln(&operacao)

	if operacao == "+" {
		fmt.Printf("Resultado: %d\n", x+y)
	}  else if operacao == "-" {
		fmt.Printf("Resultado: %d\n", x-y)
	} else if operacao == "*" {
		fmt.Printf("Resultado: %d\n", x*y)
	} else if operacao == "/" {
		fmt.Printf("Resultado: %d\n", x/y)
	} 







}
