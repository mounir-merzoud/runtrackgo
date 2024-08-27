package main

import (
    "fmt"
)

func main() {
    var sum, input int

    // Demander à l'utilisateur de saisir cinq entiers
    for i := 0; i < 5; i++ {
        fmt.Printf("Entrez l'entier %d : ", i+1)
        fmt.Scan(&input)
        sum += input
    }

    // Calculer la moyenne
    moyenne := float64(sum) / 5

    // Afficher la moyenne
    fmt.Printf("La moyenne des cinq entiers est : %.2f\n", moyenne)
}
