package main

import (
	"fmt"
)

func main() {
	const taille = 10
	var tableau [taille]int
	compteur := 0

	// Demander à l'utilisateur de saisir 10 entiers
	fmt.Println("Veuillez entrer 10 entiers :")
	for i := 0; i < taille; i++ {
		fmt.Printf("Entier %d : ", i+1)
		fmt.Scan(&tableau[i])
	}

	// Compter les entiers supérieurs ou égaux à 5
	for _, valeur := range tableau {
		if valeur >= 5 {
			compteur++
		}
	}

	// Afficher le nombre d'entiers supérieurs ou égaux à 5
	fmt.Printf("Nombre d'entiers supérieurs ou égaux à 5 : %d\n", compteur)
}
