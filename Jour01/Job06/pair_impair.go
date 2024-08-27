package main

import (
    "fmt"
)

func main() {
    var nombre int

    // Demander à l'utilisateur de saisir un nombre entier
    fmt.Print("Entrez un nombre entier : ")
    fmt.Scan(&nombre)

    // Vérifier si le nombre est pair ou impair
    if nombre%2 == 0 {
        fmt.Printf("%d est un nombre pair.\n", nombre)
    } else {
        fmt.Printf("%d est un nombre impair.\n", nombre)
    }
}
