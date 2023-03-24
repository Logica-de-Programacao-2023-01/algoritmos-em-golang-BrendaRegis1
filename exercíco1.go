package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
	var x int
	fmt.Print("Digite o primeiro número.")
	fmt.Scan(&x)

	var x1 int
	fmt.Print("Digite o segundo número.")
	fmt.Scan(&x1)

	var x2 int
	fmt.Print("Digite o terceiro número.")
	fmt.Scan(&x2)

	soma := x + x1 + x2
	fmt.Print("A soma é", soma)
}
