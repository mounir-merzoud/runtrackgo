package main

import (
	"fmt"
	"strings"
)

// Fonction pour comparer deux chaînes et retourner 0 si égales, 1 sinon
func compareStrings(s1, s2 string) int {
	if strings.Compare(s1, s2) == 0 {
		return 0
	}
	return 1
}

func main() {
	var str1, str2 string

	// Demander à l'utilisateur de saisir les deux chaînes
	fmt.Print("Entrez la première chaîne : ")
	fmt.Scanln(&str1)
	fmt.Print("Entrez la deuxième chaîne : ")
	fmt.Scanln(&str2)

	// Comparer les chaînes et obtenir le résultat
	result := compareStrings(str1, str2)

	// Afficher le résultat
	fmt.Println("Résultat de la comparaison :", result)
}
