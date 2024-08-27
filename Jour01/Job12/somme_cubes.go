package main

import (
    "fmt"
)

func main() {
    var N int
    var sum int

    // Demander à l'utilisateur de saisir un entier N
    fmt.Print("Entrez un entier N : ")
    fmt.Scan(&N)

    // Calculer la somme des cubes de 5^3 à N^3
    for i := 5; i <= N; i++ {
        sum += i * i * i
    }

    // Afficher la somme des cubes
    fmt.Printf("La somme des cubes de 5^3 à %d^3 est : %d\n", N, sum)
}
