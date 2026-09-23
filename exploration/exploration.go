package exploration

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/ascii"
	"scarlet/audio"
	"scarlet/character"
	"scarlet/combat"
	"scarlet/forge"
	"scarlet/marchand"
	"scarlet/outils"
)

const largeurCadre = 65

// Explorer lance la boucle principale d'exploration du personnage.
func Explorer(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)
	localisation := "entree_village"

	for {
		switch localisation {

		case "entree_village":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/entree_village.txt")

			lignes := []string{
				"Vous êtes à l'entrée de la ville.",
				"",
				"1. S'engager dans les étroites rues de la ville.",
				"0. Ouvrir le menu",
			}
			if c.QueteMagieNoire > 0 {
				lignes = append(lignes, "2. Prendre la route vers les champs.")
			}

			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
			case "0":
				c.Menu()
			case "2":
				if c.QueteMagieNoire > 0 {
					localisation = "champs"
				}
			default:
				fmt.Println("Choix invalide.")
			}

		case "place_marchande":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/place_marchande.txt")

			lignes := []string{
				"Vous êtes sur la place marchande.",
				"",
				"1. Retourner dans les rues.",
				"2. Entrer dans l'échoppe du marchand",
				"3. Rentrer dans la forge",
				"4. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

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
			audio.PlaySound("BanqueSon/NPC/vendeur.ogg")

			lignes := []string{
				"Vous voici dans l'échoppe du marchand.",
				"",
				"1. Retourner sur la place marchande",
				"2. Commercer avec le marchand",
				"3. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

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

			lignes := []string{
				"Vous êtes dans la forge.",
				"Le forgeron vous dit : Que puis-je faire pour vous ?",
				"",
				"1. Fabriquer un objet",
				"2. Sortir de la forge",
			}

			outils.Encadrer(lignes, largeurCadre)

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

			lignes := []string{
				"Vous arpentez les rues de la ville.",
				"Vous passez à côté d'une enseigne \"Auberge du Cul Tourné\".",
				"",
				"1. Se rendre à l'entrée de la ville.",
				"2. Se rendre sur la place marchande.",
				"3. Gravir les marches vers la cathédrale.",
				"4. Entrer dans l'auberge.",
				"0. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

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
			audio.PlayMusic("BanqueSon/musique/auberge.ogg")
			audio.PlaySound("BanqueSon/NPC/aubergistebonjour.ogg")

			lignes := []string{
				"Une naine chaleureuse s'avance vers vous :",
				"\"Bienvenue à l'Auberge ! C'est 5 pièces la chambre.\"",
				"",
				"1. Sortir de l'auberge.",
				"2. Payer 5 pièces d'or et dormir.",
				"0. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				audio.PlaySoundBlocking("BanqueSon/NPC/aubergistebye.ogg")
				audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
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

			lignes := []string{
				"La montée était rude... Vous voilà dans les beaux quartiers.",
				"Autour d'une fontaine se font face le manoir ducal et la cathédrale.",
				"",
				"1. Redescendre dans les rues.",
				"2. Entrer dans le manoir.",
				"3. Pénétrer dans la cathédrale.",
				"0. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "rue_principale"
			case "2":
				localisation = "manoir"
			case "3":
				localisation = "cathedrale"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "manoir":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/duc.txt")
			audio.PlaySound("BanqueSon/NPC/ducbonjour.ogg")

			lignes := []string{
				"Le duc se dirige vers vous :",
				"\"J'espère que vous appréciez votre séjour parmi nous.\"",
				"",
				"1. Le saluer et s'en aller.",
				"2. Parler au duc.",
				"0. Ouvrir le menu",
			}

			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				audio.PlaySoundBlocking("BanqueSon/NPC/ducbye.ogg")
				localisation = "haute_ville"

			case "2":
				if c.QueteMagieNoire == 0 {
					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/duc.txt")

					lignesQuete := []string{
						"\"Un tremblement de terre a ébranlé la ville...\"",
						"\"Une horde de gobelins ravage mes campagnes.\"",
						"\"Tuez leur chef et vous serez grassement récompensé.\"",
						"",
						"1. Accepter.",
						"2. Refuser.",
					}
					outils.Encadrer(lignesQuete, largeurCadre)

					choixQuete := lireChoix(lecteur)
					switch choixQuete {
					case "1":
						outils.ClearScreen()
						ascii.AfficherASCII("BanqueASCII/duc.txt")
						outils.Encadrer([]string{"\"Prenez cette carte, vous en aurez besoin.\""}, largeurCadre)
						c.QueteMagieNoire++
						outils.AttendreEntree()
						audio.PlaySoundBlocking("BanqueSon/NPC/ducbye.ogg")
						localisation = "haute_ville"

					case "2":
						outils.ClearScreen()
						ascii.AfficherASCII("BanqueASCII/duc.txt")
						outils.Encadrer([]string{"\"Soit. Mais vous feriez bien de rentrer dans mes bonnes grâces...\""}, largeurCadre)
						outils.AttendreEntree()
						localisation = "haute_ville"

					default:
						fmt.Println("Choix invalide.")
					}
				} else {
					outils.Encadrer([]string{"\"Avez-vous des nouvelles de la horde de gobelins ?\""}, largeurCadre)
					outils.AttendreEntree()
				}

			case "0":
				c.Menu()

			default:
				fmt.Println("Choix invalide.")
			}

		case "cathedrale":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/eglise.txt")
			audio.PlayMusic("BanqueSon/musique/cathedrale.ogg")
			audio.PlaySound("BanqueSon/NPC/bonjourpretre.ogg")

			lignes := []string{
				"Le prêtre vous adresse une bénédiction :",
				"\"Que la lumière vous guide, mon enfant.\"",
				"",
				"1. Sortir de la cathédrale.",
				"0. Ouvrir le menu.",
			}

			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
				localisation = "haute_ville"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}

		case "champs":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/carte.txt")
			fmt.Println("\nVous vous guidez sur les routes de campagne avec la carte que vous a donné le duc.")
			fmt.Println("Un gobelin sort d'une fougère et vous fonce dessus !")
			outils.AttendreEntree()
			mort, err := combat.Lancer("gobelin", c)
			if err != nil {
				fmt.Println("Erreur combat :", err)
			}

			if mort {
				localisation = "cathedrale"
			} else {
				localisation = "champs_2"
			}

		case "champs_2":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/champs_2.txt")
			lignes := []string{
				"Vous voici à un embranchement. Vous entendez des bruits dans les hautes herbes...",
				"1. Vous diriger vers la ferme.",
				"2. Inspecter les hautes herbes.",
				"3. Retourner en ville.",
				"0. Ouvrir le menu.",
			}
			outils.Encadrer(lignes, largeurCadre)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "ferme"
			case "2":
				mort, err := combat.Lancer("gobelin", c)
				if err != nil {
					fmt.Println("Erreur combat :", err)
				}

				if mort {
					localisation = "cathedrale"
				} else {
					localisation = "champs_2"
				}
			case "3":
				localisation = "entree_village"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			}
		case "ferme":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/ferme.txt")

		}
	}
}

// lireChoix lit une ligne tapée par le joueur et la nettoie
func lireChoix(lecteur *bufio.Reader) string {
	fmt.Print("\nVotre choix : ")
	saisie, _ := lecteur.ReadString('\n')
	return strings.TrimSpace(saisie)
}
