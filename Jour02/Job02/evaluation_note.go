package main

import (
	"fmt"
)

func main() {
	var note float64

	// Demander à l'utilisateur de saisir une note
	fmt.Print("Entrez une note (comprise entre 0 et 20) : ")
	fmt.Scan(&note)

	// Vérifier si la note est dans l'intervalle valide
	if note >= 0 && note <= 20 {
		// Évaluer la note
		if note >= 10 {
			fmt.Println("Validé")
		} else {
			fmt.Println("Non validé")
		}
	} else {
		fmt.Println("La note doit être comprise entre 0 et 20")
	}
}
