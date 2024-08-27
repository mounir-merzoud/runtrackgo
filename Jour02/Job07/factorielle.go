package main

import (
	"fmt"
)

// Fonction pour calculer la factorielle d'un nombre entier
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var number int

	// Demander à l'utilisateur de saisir un nombre entier
	fmt.Print("Entrez un nombre entier pour calculer sa factorielle : ")
	fmt.Scan(&number)

	// Vérifier que le nombre est non négatif
	if number < 0 {
		fmt.Println("Erreur : la factorielle n'est pas définie pour les nombres négatifs.")
		return
	}

	// Calculer et afficher la factorielle
	fact := factorial(number)
	fmt.Printf("La factorielle de %d est %d\n", number, fact)
}
