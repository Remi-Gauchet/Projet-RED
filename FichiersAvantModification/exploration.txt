package exploration

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/ascii"
	"scarlet/audio"
	"scarlet/character"
	"scarlet/forge"
	"scarlet/marchand"
	"scarlet/outils"
)

// Explorer lance la boucle principale d'exploration du personnage.
func Explorer(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)
	localisation := "entree_village"

	for {
		switch localisation {

		case "entree_village":
			outils.ClearScreen()
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
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/place_marchande.txt")
			fmt.Println("\nVous êtes sur la place marchande.")
			fmt.Println("1. Retourner à l'entrée du village")
			fmt.Println("2. Entrer dans l'échoppe du marchand")
			fmt.Println("3. Rentrer dans la forge")
			fmt.Println("4. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "entree_village"
			case "2":
				localisation = "echoppe_marchand"
			case "3":
				localisation = "forge"
			case "4":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "echoppe_marchand":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/echoppe.txt")
			fmt.Println("\nVous voici dans l'échoppe du marchand.")
			audio.PlaySound("BanqueSon/NPC/vendeur.ogg")
			fmt.Println("1. Retourner sur la place marchande")
			fmt.Println("2. Commercer avec le marchand")
			fmt.Println("3. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				audio.PlaySoundBlocking("BanqueSon/NPC/vendeurbye.ogg")
				localisation = "place_marchande"
			case "2":
				marchand.Echoppe(c)
			case "3":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "forge":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/forgeron.txt")
			audio.PlaySound("BanqueSon/NPC/forgeron.ogg")
			fmt.Println("\nVous êtes dans la forge.")
			fmt.Println("Le forgeron vous dit : Que puis-je faire pour vous ?")
			fmt.Println("1. Fabriquer un objet")
			fmt.Println("2. Sortir de la forge")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				forge.Fabriquer(c)
			case "2":
				audio.PlaySoundBlocking("BanqueSon/NPC/forgeronbye.ogg")
				localisation = "place_marchande"
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
