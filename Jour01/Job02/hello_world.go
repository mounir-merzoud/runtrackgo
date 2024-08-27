package main

import (
    "fmt"
)

func main() {
    var N int

    // Demande à l'utilisateur de saisir la valeur de N
    fmt.Print("Entrez le nombre de fois à afficher 'Hello World': ")
    fmt.Scan(&N)

    // Boucle pour afficher "Hello World" N fois
    for N != 0 {
        fmt.Println("Hello World")
        N-- // Décrémenter N
    }
}
