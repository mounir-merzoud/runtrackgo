package main

import (
	"fmt"
	"strings"
)

func main() {
	// Définir la chaîne à transformer
	original := "vive la plateforme !"

	// Convertir la chaîne en majuscules
	uppercased := strings.ToUpper(original)

	// Afficher le résultat
	fmt.Println(uppercased)
}
