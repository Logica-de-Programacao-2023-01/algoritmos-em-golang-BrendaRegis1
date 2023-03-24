package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.]

	var x int
	fmt.Println("Digite um número inteiro.")
	fmt.Scan(&x)

	dobro := 2 * x
	fmt.Println("O dobro de ", x, " é", dobro)

	triplo := 3 * x
	fmt.Print("O triplo de ", x, "é", triplo)

	quadruplo := 4 * x
	fmt.Println("O quadruplo de", x, "é", quadruplo)
}
