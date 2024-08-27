package main

import (
	"fmt"
)

func main() {
	// Demander à l'utilisateur de saisir une chaîne de caractères
	var string1 string
	fmt.Print("Veuillez entrer une chaîne de caractères : ")
	fmt.Scanln(&string1)

	// Définir la deuxième chaîne
	string2 := "Bonjour"

	// Trier les chaînes lexicographiquement
	if string1 < string2 {
		fmt.Printf("Chaîne 1 : %s\n", string1)
		fmt.Printf("Chaîne 2 : %s\n", string2)
	} else {
		fmt.Printf("Chaîne 1 : %s\n", string2)
		fmt.Printf("Chaîne 2 : %s\n", string1)
	}
}
