package main

import (
    "fmt"
)

func main() {
    var annee int

    // Demander à l'utilisateur de saisir une année
    fmt.Print("Entrez une année : ")
    fmt.Scan(&annee)

    // Déterminer si l'année est bissextile ou non
    if (annee%4 == 0 && annee%100 != 0) || (annee%400 == 0) {
        fmt.Printf("%d est une année bissextile.\n", annee)
    } else {
        fmt.Printf("%d n'est pas une année bissextile.\n", annee)
    }
}
