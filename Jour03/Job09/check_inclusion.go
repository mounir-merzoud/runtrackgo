package main

import (
	"fmt"
	"strings"
)

// Fonction qui vérifie si la chaîne 'str1' est incluse dans la chaîne 'str2'
func estIncluse(str1, str2 string) bool {
	return strings.Contains(str2, str1)
}

func main() {
	// Demander à l'utilisateur d'entrer les deux chaînes de caractères
	var chaine1, chaine2 string

	fmt.Print("Veuillez entrer la première chaîne : ")
	fmt.Scanln(&chaine1)

	fmt.Print("Veuillez entrer la deuxième chaîne : ")
	fmt.Scanln(&chaine2)

	// Vérifier si la première chaîne est incluse dans la deuxième
	if estIncluse(chaine1, chaine2) {
		fmt.Println("La première chaîne est incluse dans la deuxième.")
	} else {
		fmt.Println("La première chaîne n'est pas incluse dans la deuxième.")
	}
}
