package main

import (
	"fmt"
)

func main() {
	var limit int

	// Demander à l'utilisateur de saisir une limite
	fmt.Print("Entrez la limite jusqu'à laquelle générer la suite de Fibonacci : ")
	fmt.Scan(&limit)

	// Vérifier que la limite est positive
	if limit < 0 {
		fmt.Println("Erreur : la limite doit être un entier positif.")
		return
	}

	// Initialiser les deux premiers termes de la suite de Fibonacci
	a, b := 0, 1

	// Afficher les termes de la suite de Fibonacci jusqu'à la limite
	fmt.Print("Suite de Fibonacci jusqu'à ", limit, ": ")
	for a <= limit {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
