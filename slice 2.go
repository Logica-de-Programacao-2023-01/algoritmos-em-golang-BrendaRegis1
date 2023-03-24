package main

import "fmt"

func main() {
	//Crie um Slice de strings vazio, adicione os nomes "João", "Maria" e "José" a ele e imprima o Slice.
	nomes := []string{}
	fmt.Println(nomes)
	nomes = append(nomes, "João, Maria, José")
	fmt.Println(nomes)
}
