package main

import "fmt"

func main() {
	
var x, y int
var operacao string

fmt.Println("Digite o primeiro número: ")
fmt.Scan(&x)
fmt.Println("Digite o segundo número: ")
fmt.Scan(&y)

fmt.Println("Digite a operação desejada (+, -, *, /): ")
fmt.Scan(&operacao)

if operacao == "+" {
	fmt.Println("Resultado:", x+y)
} else if operacao == "-" {
	fmt.Println("Resultado:", x-y)
} else if operacao == "*" {
	fmt.Println("Resultado:", x*y)
} else if operacao == "/" {
	if y != 0 {
		fmt.Println("Resultado:", x/y)
	} else {
		fmt.Println("Erro: Divisão por zero não é permitida.")
	}
}
}