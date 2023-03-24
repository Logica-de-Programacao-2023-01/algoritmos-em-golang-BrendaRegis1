package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).

	var x float64
	fmt.Print("Qual é o seu peso?")
	fmt.Scan(&x)

	var x1 float64
	fmt.Print("Qual é a sua altura?")
	fmt.Scan(&x1)

	IMC := x / (x1 * x1)
	fmt.Print("Seu IMC é", IMC)
}
