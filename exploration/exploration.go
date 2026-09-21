package exploration

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/ascii"
	"scarlet/character"
)

// Explorer lance la boucle principale d'exploration du personnage.
func Explorer(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)
	localisation := "entree_village"

	for {
		switch localisation {

		case "entree_village":
			ascii.AfficherASCII("BanqueASCII/entree_village.txt")
			fmt.Println("\nVous êtes à l'entrée du village.")
			fmt.Println("1. Se rendre à la place marchande")
			fmt.Println("2. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "place_marchande"
			case "2":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "place_marchande":
			ascii.AfficherASCII("BanqueASCII/place_marchande.txt")
			fmt.Println("\nVous êtes sur la place marchande.")
			fmt.Println("1. Retourner à l'entrée du village")
			fmt.Println("2. Entrer dans l'échoppe du marchand")
			fmt.Println("3. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "entree_village"
			case "2":
				localisation = "echoppe_marchand"
			case "3":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "echoppe_marchand":
			ascii.AfficherASCII("BanqueASCII/echoppe_marchand.txt")
			fmt.Println("\nVous êtes dans l'échoppe du marchand.")
			fmt.Println("1. Retourner à la place marchande")
			fmt.Println("2. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "place_marchande"
			case "2":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}
		}
	}
}

// lireChoix lit une ligne tapée par le joueur et la nettoie
func lireChoix(lecteur *bufio.Reader) string {
	fmt.Print("Votre choix : ")
	saisie, _ := lecteur.ReadString('\n')
	return strings.TrimSpace(saisie)
}
