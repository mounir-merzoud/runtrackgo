package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	for {
		// Demander à l'utilisateur de saisir un nombre réel
		fmt.Print("Entrez un nombre réel (0 pour quitter) : ")
		fmt.Scan(&number)

		// Vérifier si le nombre est égal à 0 pour arrêter le programme
		if number == 0 {
			fmt.Println("Programme terminé.")
			break
		}

		// Vérifier si le nombre est négatif
		if number < 0 {
			fmt.Println("Erreur : la valeur ne peut pas être négative.")
			continue
		}

		// Calculer la racine carrée
		sqrt := math.Sqrt(number)
		fmt.Printf("La racine carrée de %.2f est %.2f\n", number, sqrt)
	}
}
