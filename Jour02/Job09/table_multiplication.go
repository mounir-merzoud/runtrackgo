package main

import (
	"fmt"
)

func main() {
	const size = 10

	// Affichage de l'en-tête de la table
	fmt.Printf("    ")
	for i := 1; i <= size; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	// Affichage de la ligne de séparation
	fmt.Println("  " + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----")

	// Affichage de la table de multiplication
	for i := 1; i <= size; i++ {
		// Affichage du numéro de ligne
		fmt.Printf("%2d|", i)
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
