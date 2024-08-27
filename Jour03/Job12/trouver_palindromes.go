package main

import (
	"fmt"
	"strings"
)

// Fonction pour vérifier si une chaîne est un palindrome
func estPalindrome(s string) bool {
	// Convertir la chaîne en minuscules pour une comparaison insensible à la casse
	s = strings.ToLower(s)
	// Inverser la chaîne
	// Conversion en rune pour gérer correctement les caractères Unicode
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Définir le tableau de chaînes
	tab := []string{"radar", "hello", "lvel", "stats", "world"}

	fmt.Println("Palindromes trouvés :")

	// Parcourir le tableau et afficher les palindromes
	for _, s := range tab {
		if estPalindrome(s) {
			fmt.Println(s)
		}
	}
}
