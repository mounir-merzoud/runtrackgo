package main

import (
	"fmt"
	"strings"
)

// Fonction pour supprimer les voyelles d'une chaîne
func removeVowels(s string) string {
	// Définir les voyelles
	vowels := "aeiouAEIOU"
	var result strings.Builder

	// Parcourir chaque caractère de la chaîne
	for _, char := range s {
		// Vérifier si le caractère est une voyelle
		if !strings.ContainsRune(vowels, char) {
			// Ajouter le caractère à la chaîne résultante
			result.WriteRune(char)
		}
	}

	// Retourner la chaîne résultante
	return result.String()
}

func main() {
	// Définir la chaîne d'origine
	original := "vive la plateforme !"

	// Supprimer les voyelles de la chaîne
	noVowels := removeVowels(original)

	// Afficher le résultat
	fmt.Println(noVowels)
}
