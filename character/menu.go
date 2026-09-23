package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/outils"
)

func (c *Character) Menu() {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		outils.ClearScreen()

		fmt.Println("\n\n\n\n\n\n\n\n\n")

		menuLignes := []string{
			"--- Menu Principal ---",
			"",
			"1. Afficher les informations du personnage",
			"2. Accéder à l'inventaire",
			"3. Équipement",
			"4. Sortir du menu",
		}

		outils.Encadrer(menuLignes)

		// Exactly 74 spaces before the prompt
		fmt.Print("                                                                          Entrez le numéro de votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			outils.ClearScreen()
			c.DisplayInfo()
			attendreEntree(lecteur)

		case "2":
			outils.ClearScreen()
			c.AccessInventory()

		case "3":
			outils.ClearScreen()
			c.GererEquipement()

		case "4":
			fmt.Println("Fermeture du menu")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func attendreEntree(lecteur *bufio.Reader) {
	fmt.Println("\nAppuyez sur Entrée pour continuer...")
	lecteur.ReadString('\n')
}
