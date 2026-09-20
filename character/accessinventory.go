package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AccessInventory affiche le contenu de l'inventaire, puis propose un menu
// d'actions selon les objets présents (fermer l'inventaire, boire une potion...)
func (c *Character) AccessInventory() {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		if len(c.Inventaire) == 0 {
			fmt.Println("Inventaire vide.")
		} else {
			detail := strings.Join(c.Inventaire, " - ")
			fmt.Println(detail)
		}

		auMoinsUnePotion := false
		for _, objet := range c.Inventaire {
			if objet == "Potion" {
				auMoinsUnePotion = true
				break
			}
		}

		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Fermer l'inventaire")
		if auMoinsUnePotion {
			fmt.Println("2. Boire une potion")
		}
		fmt.Print("Selon votre choix, tapez le numéro puis entrée : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			return // sort de la fonction, referme l'inventaire

		case "2":
			if !auMoinsUnePotion {
				fmt.Println("Choix invalide.")
				fmt.Println()
				continue
			}
			err := c.TakePot()
			if err != nil {
				fmt.Println("Erreur :", err)
			}
			fmt.Println() // ligne vide pour aérer avant de réafficher le menu

		default:
			fmt.Println("Choix invalide.")
			fmt.Println()
		}
	}
}
