package main

import "fmt"

func main() { //Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles,
	// com pesos 2, 3 e 5, respectivamente.

	var x float64
	fmt.Println("Digite o primeiro número.")
	fmt.Scan(&x)

	var x1 float64
	fmt.Println("Digite o segundo número.")
	fmt.Scan(&x1)

	var x2 float64
	fmt.Println("Digite o terceiro número.")
	fmt.Scan(&x2)
	MP := (2 * x) + (3 * x1) + (5 * x2)
	fmt.Println("A média ponderada dos números é", MP)
}
