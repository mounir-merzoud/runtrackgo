package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64

    // Demander à l'utilisateur de saisir le premier nombre
    fmt.Print("Entrez le premier nombre : ")
    fmt.Scan(&num1)

    // Demander à l'utilisateur de saisir le deuxième nombre
    fmt.Print("Entrez le deuxième nombre : ")
    fmt.Scan(&num2)

    // Calculer la somme des deux nombres
    somme := num1 + num2

    // Afficher le résultat
    fmt.Printf("La somme de %.2f et %.2f est %.2f\n", num1, num2, somme)
}
