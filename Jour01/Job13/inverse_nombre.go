package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Demander à l'utilisateur d'entrer un nombre
    var nombre int
    fmt.Print("Entrez un nombre entier : ")
    fmt.Scan(&nombre)

    // Convertir le nombre en chaîne de caractères pour faciliter l'inversion
    strNombre := strconv.Itoa(nombre)

    // Inverser la chaîne de caractères
    inverseStr := ""
    for i := len(strNombre) - 1; i >= 0; i-- {
        inverseStr += string(strNombre[i])
    }

    // Convertir la chaîne de caractères inversée en entier
    inverse, err := strconv.Atoi(inverseStr)
    if err != nil {
        fmt.Println("Erreur lors de la conversion du nombre inversé :", err)
        return
    }

    // Afficher le nombre inversé
    fmt.Printf("Le nombre inversé est : %d\n", inverse)
}
