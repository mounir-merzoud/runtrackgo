package main

import (
    "fmt"
)

func main() {
    var a, b, c int

    // Demander à l'utilisateur de saisir trois entiers
    fmt.Print("Entrez le premier entier : ")
    fmt.Scan(&a)

    fmt.Print("Entrez le deuxième entier : ")
    fmt.Scan(&b)

    fmt.Print("Entrez le troisième entier : ")
    fmt.Scan(&c)

    // Déterminer le plus grand entier
    var max int
    if a > b && a > c {
        max = a
    } else if b > a && b > c {
        max = b
    } else {
        max = c
    }

    // Afficher le plus grand entier
    fmt.Printf("Le plus grand des trois entiers est : %d\n", max)
}
