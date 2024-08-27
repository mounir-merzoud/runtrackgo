package main

import (
	"fmt"
)

func main() {
	// Demander à l'utilisateur d'entrer une chaîne de caractères
	fmt.Print("Veuillez entrer une chaîne de caractères : ")
	var input string
	fmt.Scanln(&input)

	// Convertir la chaîne en un tableau de runes (caractères Unicode)
	tab := []rune(input)

	// Afficher les caractères stockés dans le tableau
	fmt.Println("Les caractères stockés dans le tableau sont :")
	for i, char := range tab {
		fmt.Printf("Index %d : %c\n", i, char)
	}

	// Afficher la longueur du tableau (pour référence)
	fmt.Printf("Longueur du tableau : %d\n", len(tab))
}
