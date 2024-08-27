package main

import (
	"fmt"
)

func main() {
	var n int

	// Demander à l'utilisateur de saisir un nombre entier n
	fmt.Print("Entrez un entier n pour calculer la somme des n premiers termes de la série harmonique : ")
	fmt.Scan(&n)

	// Vérifier que n est un nombre positif
	if n <= 0 {
		fmt.Println("Erreur : n doit être un entier positif.")
		return
	}

	// Calculer la somme des n premiers termes de la série harmonique
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}

	// Afficher le résultat
	fmt.Printf("La somme des %d premiers termes de la série harmonique est %.6f\n", n, sum)
}
