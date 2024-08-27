package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Définir le format attendu de l'heure
	const pattern = `^\d{2}h\d{2}$`

	// Créer une expression régulière avec le motif défini
	re := regexp.MustCompile(pattern)

	// Demander à l'utilisateur de saisir l'heure
	fmt.Print("Entrez l'heure au format XXhXX : ")
	var input string
	fmt.Scanln(&input)

	// Vérifier si l'entrée correspond au format attendu
	if re.MatchString(input) {
		fmt.Println("Le format est correct.")
	} else {
		fmt.Println("Le format est incorrect.")
	}
}
