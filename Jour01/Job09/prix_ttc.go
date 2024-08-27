package main

import (
    "fmt"
)

func main() {
    var prixHT float64
    var nombreKilos float64
    var tauxTVA float64

    // Demander à l'utilisateur de saisir le prix HT d'un kilo de carottes
    fmt.Print("Entrez le prix HT d'un kilo de carottes (en euros) : ")
    fmt.Scan(&prixHT)

    // Demander à l'utilisateur de saisir le nombre de kilos de carottes
    fmt.Print("Entrez le nombre de kilos de carottes : ")
    fmt.Scan(&nombreKilos)

    // Demander à l'utilisateur de saisir le taux de TVA en pourcentage
    fmt.Print("Entrez le taux de TVA (en pourcentage, ex : 15) : ")
    fmt.Scan(&tauxTVA)

    // Calculer le prix TTC
    prixTTC := prixHT * nombreKilos * (1 + tauxTVA/100)

    // Afficher le prix TTC
    fmt.Printf("Le prix TTC de %.2f kilos de carottes est : %.2f euros\n", nombreKilos, prixTTC)
}
