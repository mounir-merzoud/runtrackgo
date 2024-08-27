package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Définir la plage des nombres et le nombre de tentatives
	const min = 0
	const max = 100
	const maxAttempts = 10

	// Générer un nombre aléatoire entre min et max
	target := rand.Intn(max-min+1) + min

	fmt.Println("Bienvenue au jeu 'Nombre mystère' !")
	fmt.Printf("Devinez le nombre entre %d et %d.\n", min, max)
	fmt.Printf("Vous avez %d tentatives pour trouver le nombre.\n", maxAttempts)

	var guess int
	for attempts := 0; attempts < maxAttempts; attempts++ {
		fmt.Print("Entrez votre nombre : ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'entrée. Veuillez entrer un nombre valide.")
			// Réessayer sans augmenter le compteur de tentatives
			attempts--
			continue
		}

		if guess < min || guess > max {
			fmt.Printf("Veuillez entrer un nombre entre %d et %d.\n", min, max)
			attempts-- // Réessayer sans augmenter le compteur de tentatives
			continue
		}

		if guess < target {
			fmt.Println("Trop petit ! Essayez encore.")
		} else if guess > target {
			fmt.Println("Trop grand ! Essayez encore.")
		} else {
			fmt.Println("Félicitations ! Vous avez trouvé le nombre !")
			break
		}

		if attempts == maxAttempts-1 {
			fmt.Printf("Désolé, vous avez utilisé toutes vos tentatives. Le nombre mystère était %d.\n", target)
		}
	}
}
