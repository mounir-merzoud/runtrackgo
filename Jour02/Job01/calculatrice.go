package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	// Demander à l'utilisateur de saisir deux nombres
	fmt.Print("Entrez le premier nombre: ")
	fmt.Scan(&num1)

	fmt.Print("Entrez le deuxième nombre: ")
	fmt.Scan(&num2)

	// Demander à l'utilisateur de choisir une opération
	fmt.Print("Entrez l'opération (+, -, *, /): ")
	fmt.Scan(&operation)

	// Effectuer l'opération en fonction de l'entrée de l'utilisateur
	switch operation {
	case "+":
		fmt.Printf("Résultat : %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Résultat : %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Résultat : %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Résultat : %.2f\n", num1/num2)
		} else {
			fmt.Println("Erreur : Division par zéro")
		}
	default:
		fmt.Println("Opération non reconnue")
	}
}
