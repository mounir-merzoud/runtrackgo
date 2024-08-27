package main

import (
	"fmt"
)

func main() {
	var n int

	// Demander à l'utilisateur de saisir un entier n
	fmt.Print("Entrez un entier positif n : ")
	fmt.Scan(&n)

	// Itérer de 1 à n
	for i := 1; i <= n; i++ {
		// Vérifier si i est un multiple de 3 ou de 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d est un multiple de 3 et de 5\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d est un multiple de 3\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d est un multiple de 5\n", i)
		}
	}
}
