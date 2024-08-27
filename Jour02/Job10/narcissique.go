package main

import (
	"fmt"
	"math"
)

// Fonction pour vérifier si un nombre est narcissique
func isNarcissistic(number int) bool {
	// Convertir le nombre en chaîne de caractères pour extraire les chiffres
	strNumber := fmt.Sprintf("%d", number)
	length := len(strNumber)
	
	// Calculer la somme des puissances des chiffres
	sum := 0
	for _, digit := range strNumber {
		digitValue := int(digit - '0') // Convertir le caractère en chiffre
		sum += int(math.Pow(float64(digitValue), float64(length)))
	}
	
	// Vérifier si la somme est égale au nombre initial
	return sum == number
}

func main() {
	var number int

	// Demander à l'utilisateur de saisir un nombre
	fmt.Print("Entrez un nombre pour vérifier s'il est narcissique : ")
	fmt.Scan(&number)

	// Vérifier si le nombre est narcissique
	if isNarcissistic(number) {
		fmt.Printf("%d est un nombre narcissique.\n", number)
	} else {
		fmt.Printf("%d n'est pas un nombre narcissique.\n", number)
	}
}
