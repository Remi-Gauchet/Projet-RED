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
			fmt.Println("\nVous êtes à l'entrée de la ville.")
			fmt.Println("1. S'engager dans les étroites rues de la ville.")
			fmt.Println("2. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
			case "2":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "place_marchande":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/place_marchande.txt")
			fmt.Println("\nVous êtes sur la place marchande.")
			fmt.Println("1. Retourner dans les rues.")
			fmt.Println("2. Entrer dans l'échoppe du marchand")
			fmt.Println("3. Rentrer dans la forge")
			fmt.Println("4. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
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

		case "rue_principale":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/rue_principale.txt")
			fmt.Println("\nVous arpentez les rues de la ville. Certaines sont gorgées de monde. Et d'autres, de véritables coupe gorge.")
			fmt.Println("Vous passez à côté d'une enseigne \"Auberge du Cul Tourné\".")
			fmt.Println("1. Se rendre à l'entrée de la ville.")
			fmt.Println("2. Se rendre sur la place marchande.")
			fmt.Println("3. Gravir les marches vers la cathédrale.")
			fmt.Println("4. Entrer dans l'auberge.")
			fmt.Println("0. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "entree_village"
			case "2":
				localisation = "place_marchande"
			case "3":
				localisation = "haute_ville"
			case "4":
				localisation = "auberge"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "auberge":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/auberge.txt")
			fmt.Println("\nÀ peine la porte ouverte, une odeur d'alcool vous submerge.")
			fmt.Println("Une naine chaleureuse s'avance vers vous avec des pintes de bière :")
			fmt.Println("\"Bienvenue à l'Auberge du Cul Tourné ! Vous voulez dormir ici ? C'est 5 pièces la chambre.\"")
			fmt.Println("1. Sortir de l'auberge.")
			fmt.Println("2. Payer 5 pièces d'or et dormir à l'Auberge du Cul Tourné.")
			fmt.Println("0. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
			case "2":
				err := c.SeReposer()
				if err != nil {
					fmt.Println("Erreur :", err)
				}
				fmt.Println("\nAppuyez sur Entrée pour continuer...")
				lecteur.ReadString('\n')
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "haute_ville":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/haute_ville.txt")
			fmt.Println("\nLa montée était rude... Mais vous voilà dans les beaux quartiers.")
			fmt.Println("Autour d'une fontaine se font face l'hôtel de ville et la cathédrale.")
			fmt.Println("1. Redescendre dans les rues.")
			fmt.Println("2. Entrer dans l'hôtel de ville.")
			fmt.Println("3. Pénétrer dans la cathédrale.")
			fmt.Println("0. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
			case "2":
				localisation = "hotel_de_ville"
			case "3":
				localisation = "cathedrale"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "hotel_de_ville":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/maire.txt")
			fmt.Println("\nL'interieur est décoré de boiseries sculptées et de quelques statues de marbre. Vous profitez un instant du feu de cheminée.")
			fmt.Println("Par les fenêtres, vous entrapercez une vue magnifique sur le reste de la ville !")
			fmt.Println("Le maire, un homme de petite stature mais richement habillé, se dirige vers vous :")
			fmt.Println("\"J'espère que appréciez votre séjour parmi nous.\"")
			fmt.Println("1. Le saluer et s'en aller.")
			fmt.Println("2. Parler au maire.")
			fmt.Println("0. Ouvrir le menu")

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "haute_ville"
			case "2":

			case "0":
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
