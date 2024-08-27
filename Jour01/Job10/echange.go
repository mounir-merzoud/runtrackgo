package main

import (
    "fmt"
)

func main() {
    var n, m int

    // Demander à l'utilisateur de saisir deux entiers
    fmt.Print("Entrez le premier entier (n) : ")
    fmt.Scan(&n)

    fmt.Print("Entrez le deuxième entier (m) : ")
    fmt.Scan(&m)

    // Afficher les valeurs avant l'échange
    fmt.Printf("Avant l'échange: n = %d, m = %d\n", n, m)

    // Échanger les valeurs
    n, m = m, n

    // Afficher les valeurs après l'échange
    fmt.Printf("Après l'échange: n = %d, m = %d\n", n, m)
}
