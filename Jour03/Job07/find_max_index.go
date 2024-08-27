package main

import (
	"fmt"
)

func main() {
	const taille = 10
	var tableau [taille]int
	var maxIndice int

	// Demander à l'utilisateur de saisir 10 entiers
	fmt.Println("Veuillez entrer 10 entiers :")
	for i := 0; i < taille; i++ {
		fmt.Printf("Entier %d : ", i+1)
		fmt.Scan(&tableau[i])
	}

	// Trouver l'indice du plus grand élément
	max := tableau[0]
	maxIndice = 0
	for i := 1; i < taille; i++ {
		if tableau[i] > max {
			max = tableau[i]
			maxIndice = i
		}
	}

	// Afficher l'indice du plus grand élément
	fmt.Printf("L'indice du plus grand élément est : %d\n", maxIndice)
}
