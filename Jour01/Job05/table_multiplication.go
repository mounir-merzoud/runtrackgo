package main

import (
    "fmt"
)

func main() {
    var nombre int

    // Demander à l'utilisateur de saisir un entier
    fmt.Print("Entrez un nombre pour afficher sa table de multiplication : ")
    fmt.Scan(&nombre)

    // Afficher la table de multiplication de l'entier saisi
    fmt.Printf("Table de multiplication de %d :\n", nombre)
    for i := 1; i <= 10; i++ {
        produit := nombre * i
        fmt.Printf("%d x %d = %d\n", nombre, i, produit)
    }
}
