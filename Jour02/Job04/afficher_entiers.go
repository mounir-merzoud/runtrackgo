package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Demander à l'utilisateur de saisir les bornes a et b
	fmt.Print("Entrez la borne inférieure a : ")
	fmt.Scan(&a)

	fmt.Print("Entrez la borne supérieure b : ")
	fmt.Scan(&b)

	// Vérifier que a est inférieur ou égal à b pour éviter des erreurs d'affichage
	if a > b {
		fmt.Println("Erreur : la borne inférieure doit être inférieure ou égale à la borne supérieure.")
		return
	}

	// Utiliser une boucle for pour afficher tous les entiers de a à b inclus
	for i := a; i <= b; i++ {
		fmt.Println(i)
	}
}
