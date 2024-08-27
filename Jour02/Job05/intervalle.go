package main

import (
	"fmt"
)

func main() {
	var a, b, x int

	// Demander à l'utilisateur de saisir les bornes a et b
	fmt.Print("Entrez la borne inférieure a (doit être inférieur à b) : ")
	fmt.Scan(&a)

	fmt.Print("Entrez la borne supérieure b (doit être supérieur à a) : ")
	fmt.Scan(&b)

	// Vérifier que a est inférieur à b
	if a >= b {
		fmt.Println("Erreur : la borne inférieure a doit être strictement inférieure à la borne supérieure b.")
		return
	}

	// Demander à l'utilisateur de saisir l'entier x
	fmt.Print("Entrez un entier x : ")
	fmt.Scan(&x)

	// Vérifier si x est dans l'intervalle [a, b]
	if x >= a && x <= b {
		fmt.Println("GAGNE")
	} else {
		fmt.Println("PERDU")
	}
}
