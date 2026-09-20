package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Menu affiche le menu principal du jeu et redirige vers les actions du personnage.
// Le joueur peut y revenir autant de fois qu'il le souhaite au cours de la partie.
func (c *Character) Menu() {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter sans sauvegarder")
		fmt.Print("Entrez le numéro de votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			c.DisplayInfo()

		case "2":
			c.AccessInventory()

		case "3":
			fmt.Println("À bientôt ! La bise !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
