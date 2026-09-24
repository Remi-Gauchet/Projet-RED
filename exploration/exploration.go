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
			if c.QueteMagieNoire > 2 {
				lignes = append(lignes, "3. Se guider à l'aide du cristal traqueur. Chercher la source de magie noire.")
			}
			outils.Encadrer(lignes)

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
			case "3":
				if c.QueteMagieNoire > 0 {
					localisation = "grotte"
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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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

			outils.Encadrer(lignes)

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
					outils.Encadrer(lignesQuete)

					choixQuete := lireChoix(lecteur)
					switch choixQuete {
					case "1":
						outils.ClearScreen()
						ascii.AfficherASCII("BanqueASCII/duc.txt")
						outils.Encadrer([]string{"\"Prenez cette carte, vous en aurez besoin.\""})
						c.QueteMagieNoire++
						outils.AttendreEntree()
						audio.PlaySoundBlocking("BanqueSon/NPC/ducbye.ogg")
						localisation = "haute_ville"

					case "2":
						outils.ClearScreen()
						ascii.AfficherASCII("BanqueASCII/duc.txt")
						outils.Encadrer([]string{"\"Soit. Mais vous feriez bien de rentrer dans mes bonnes grâces...\""})
						outils.AttendreEntree()
						localisation = "haute_ville"

					default:
						fmt.Println("Choix invalide.")
					}
				} else {
					outils.Encadrer([]string{"\"Avez-vous des nouvelles de la horde de gobelins ?\""})
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

			lignes := []string{
				"L'interieur de la cathédrale est grandiose !",
				"La lumière ruisselle à travers les vitraux et retombe sur un autel immaculé.",
				"Dans l'allée, le prêtre passe le balai sur le sol recouvert de pétales de roses.",
				"1. Sortir de la cathédrale.",
				"2. S'adresser au prêtre.",
				"0. Ouvrir le menu.",
			}

			outils.Encadrer(lignes)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
				localisation = "haute_ville"
			case "0":
				c.Menu()
			default:
				fmt.Println("Choix invalide.")
			case "2":
				if c.QueteMagieNoire == 2 {
					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/eglise.txt")
					audio.PlaySound("BanqueSon/NPC/bonjourpretre.ogg")
					lignes := []string{
						"\"Je sens une ombre émaner de vous, mon enfant... Expliquez vous.\"",
						"Vous montrez le cristal noir au prêtre.",
						"\"Par la lumière, un cristal de détonation ! Lâchez ça sur le champs !\"",
						"Vous lui expliquez tout...",
					}
					outils.Encadrer(lignes)
					outils.AttendreEntree()

					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/eglise.txt")
					c.QueteMagieNoire = 3
					lignes = []string{
						"\"Je vois... C'était donc ça, l'origine de ces séismes. Mais il doit y en avoir d'autres. Laissez moi lancer un sort.\"",
						"D'un geste de la main, le prêtre illumine le cristal. Il devient translucide, sauf sa pointe qui reste obscure en fonction de comment vous le tenez.",
						"\"Voilà, un traqueur de magie noire. J'ai appris ça quand j'officiais parmi les inquisiteurs du bucher sacré...\"",
						"\"L'extrêmité de ce cristal pointe désormais vers l'emplacement du rituel où il a été conçu.\"",
						"Appuyez sur entrée pour sortir de la cathédrale",
					}
					outils.Encadrer(lignes)
					outils.AttendreEntree()

					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/eglise.txt")
					lignes = []string{
						"\"Attendez ! Je pense savoir qui est derrière tout ça. Cette magie noire me rappelle une vieille ennemie.\"",
						"\"La sorcière Morgane. L'une des mage noir les plus fourbes... Nous courons un grave danger !\"",
						"\"J'envoie immédiatement une missive à l'ordre des paladins. Et, mon enfant... Si c'est vraiment elle, je vous déconseille de suivre ce cristal.\"",
					}
					outils.Encadrer(lignes)
					outils.AttendreEntree()
					audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
					localisation = "haute_ville"

				} else {
					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/eglise.txt")
					audio.PlaySound("BanqueSon/NPC/bonjourpretre.ogg")
					lignes := []string{
						"Que la lumière soit toujours avec vous, mon enfant. Qu'elle vous accompagne dans votre voyage",
					}
					outils.Encadrer(lignes)
					outils.AttendreEntree()
					localisation = "cathedrale"
				}
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
				audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
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
			outils.Encadrer(lignes)

			choix := lireChoix(lecteur)
			switch choix {
			case "1":
				localisation = "ferme"
			case "2":
				mort, err := combat.Lancer("RoiGobelin", c)
				if err != nil {
					fmt.Println("Erreur combat :", err)
				}

				if mort {
					localisation = "cathedrale"
				} else {
					localisation = "champs_2"
					audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
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
			if c.QueteMagieNoire < 2 {
				audio.PlaySound("BanqueSon/NPC/hurlement.ogg")
				lignesQuete := []string{
					"La ferme est attaquée ! Les paysans se battent avec leurs fourches... Ils ont besoin d'aide !",
					"Un gobelin plus grand que les autres donne des ordres. Ce doit être leur chef !",
				}
				outils.Encadrer(lignesQuete)
				outils.AttendreEntree()
				mort, err := combat.Lancer("roigobelin", c)
				if err != nil {
					fmt.Println("Erreur combat :", err)
				}
				if mort {
					localisation = "cathedrale"
				} else {
					c.QueteMagieNoire++
					outils.ClearScreen()
					ascii.AfficherASCII("BanqueASCII/ferme.txt")
					audio.PlaySound("BanqueSon/NPC/nainbonjour.ogg")
					lignesQuete := []string{
						"Les gobelins se dispersent ! La campagne devrait avoir un peu de répit. Un vieux fermier vous approche en s'essuyant le front :",
						"\"Vindiou ! Ya d'ces bestiaux d'nos jours ! Merci pour l'coup d'main, pour sûr !\"",
						"Sur le corps de la bête, vous ramassez un cristal noir, chaud au toucher, qui pulse une énergie malfaisante...",
						"Le paysan et sa femmme vous serrent la main. Et vous décidez de poursuivre votre route.",
					}
					outils.Encadrer(lignesQuete)
					outils.AttendreEntree()
					localisation = "champs_2"
				}
			} else {
				audio.PlaySound("BanqueSon/NPC/nainbonjour.ogg")
				lignesQuete := []string{
					"Dès qu'ils vous voient arriver, le paysan et sa femme se précipitent pour vous saluer chaleureusement.",
					"1. Reprendre la route",
					"0. Ouvrir le menu.",
				}
				outils.Encadrer(lignesQuete)
				choix := lireChoix(lecteur)
				switch choix {
				case "1":
					localisation = "champs_2"
				case "0":
					c.Menu()
				default:
					fmt.Println("Choix invalide.")
				}
			}

		case "grotte":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/grotte.txt")
			audio.PlayMusic("BanqueSon/musique/grotte.ogg")

			lignes := []string{
				"Votre cristal traqueur vous a mené à une grotte, bien cachée derrière la végétation.",
				"Le traqueur s'affole : vous devez vous repérer par vous-même. À propos, il y a des traces de pas fraiches...",
				"Hein ? Qu'est ce que c'était ? Quelque chose se déplace dans les ombre !",
				"Ça vient vers vous !",
			}
			outils.Encadrer(lignes)
			mort, err := combat.Lancer("demon", c)
			if err != nil {
				fmt.Println("Erreur combat :", err)
			}

			if mort {
				localisation = "cathedrale"
			} else {
				outils.ClearScreen()
				ascii.AfficherASCII("BanqueASCII/grotte.txt")
				audio.PlayMusic("BanqueSon/musique/grotte.ogg")
				lignesQuete := []string{
					"C'était un démon ! Qu'est ce qu'une créature des Tréfonds fait si proche de la surface ?",
					"1. Suivre les traces de pas, et vous enfoncer plus profondément dans la grotte.",
					"2. Retourner dans les champs.",
					"0. Ouvrir le menu.",
				}
				outils.Encadrer(lignesQuete)

				choix := lireChoix(lecteur)
				switch choix {
				case "1":
					localisation = "trefonds"
				case "2":
					localisation = "champs_2"
				case "0":
					c.Menu()
				default:
					fmt.Println("Choix invalide.")
				}
			}

		case "trefonds":
			outils.ClearScreen()
			ascii.AfficherASCII("BanqueASCII/trefonds.txt")

			lignes := []string{
				"Vous remarquez des cristaux fracturés sur le sol, et des roches vitrifiées. Une grande explosion a eu lieu ici.",
				"Un courant d'air ?  Le passage s'élargit vers un gouffre, et vous arrivez au bout du chemin.",
				"Non... ce sont les Tréfonds ! Ce qu'il restes de la civilisation naine à son apogée, avant son ravage par les démons.",
				"Vous essayez de passer inaperçu, mais une patrouille vous repère !",
			}
			outils.Encadrer(lignes)
			mort, err := combat.Lancer("demon", c)
			if err != nil {
				fmt.Println("Erreur combat :", err)
			}

			if mort {
				localisation = "cathedrale"
			} else {
				outils.ClearScreen()
				ascii.AfficherASCII("BanqueASCII/trefonds.txt")
				audio.PlayMusic("BanqueSon/musique/grotte.ogg")
				c.QueteMagieNoire = 4
				lignes := []string{
					"Vite, avant que d'autres n'arrivent !",
					"1. Faire demi tour et courir vers l'entrée de la grotte.",
					"0. Menu.",
				}
				outils.Encadrer(lignes)
				choix := lireChoix(lecteur)
				switch choix {
				case "1":
					localisation = "morgane"
				case "0":
					c.Menu()
				default:
					fmt.Println("Choix invalide.")
				}

			}

		}
	}
}

// lireChoix lit une ligne tapée par le joueur et la nettoie
func lireChoix(lecteur *bufio.Reader) string {
	fmt.Print("\nVotre choix : ")
	saisie, _ := lecteur.ReadString('\n')
	return strings.TrimSpace(saisie)
}
