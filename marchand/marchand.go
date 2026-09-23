package marchand

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"scarlet/ascii"
	"scarlet/character"
	"scarlet/outils"
)

// Objet représente un article vendu par le marchand
type Objet struct {
	Nom  string
	Prix int
}

var catalogue = []Objet{
	{Nom: "Potion de soins", Prix: 3},
	{Nom: "Potion Poison", Prix: 6},
	{Nom: "Fourrure de Loup", Prix: 4},
	{Nom: "Peau de Troll", Prix: 7},
	{Nom: "Cuir de Sanglier", Prix: 3},
	{Nom: "Plume de Corbeau", Prix: 1},
	{Nom: "Livre de Sort : Boule de Feu", Prix: 25},
	{Nom: "Livre de Sort : Glace", Prix: 30},
	{Nom: "Livre de Sort : Lumière Divine", Prix: 35},
}

// Echoppe ouvre le menu de la boutique du marchand : achat et vente d'objets.
func Echoppe(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		outils.ClearScreen()
		ascii.AfficherASCII("BanqueASCII/echoppe.txt")

		var lignes []string
		lignes = append(lignes, fmt.Sprintf("--- Boutique du marchand --- (Or : %d)", c.Or))
		lignes = append(lignes, "Jetez un oeil à mes marchandises !")
		lignes = append(lignes, "")

		// Articles du catalogue (numérotés de 1 à len(catalogue))
		for i, objet := range catalogue {
			lignes = append(lignes, fmt.Sprintf("%d. Acheter %s - %d or", i+1, objet.Nom, objet.Prix))
		}

		lignes = append(lignes, "")

		// Options secondaires
		numVendre := len(catalogue) + 1
		numUpgrade := len(catalogue) + 2
		numQuitter := len(catalogue) + 3

		lignes = append(lignes, fmt.Sprintf("%d. Vendre un objet", numVendre))
		lignes = append(lignes, fmt.Sprintf("%d. Améliorer la capacité de l'inventaire (%d or) [%d/%d]", numUpgrade, character.CoutUpgrade, character.MaxUpgrades-c.UpgradesRestantes, character.MaxUpgrades))
		lignes = append(lignes, fmt.Sprintf("%d. Quitter la boutique", numQuitter))

		// Encadrer les articles et options avec une largeur de 90
		outils.Encadrer(lignes)

		fmt.Print("                                                                    Votre choix : ")
		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		valeurChoix, err := strconv.Atoi(choix)
		if err != nil {
			fmt.Println("Choix invalide.")
			attendreEntree(lecteur)
			continue
		}

		switch valeurChoix {
		case numQuitter:
			return
		case numVendre:
			vendre(c, lecteur)
		case numUpgrade:
			err := c.UpgradeInventorySlot()
			if err != nil {
				fmt.Println("Erreur :", err)
			}
			attendreEntree(lecteur)
		default:
			if valeurChoix >= 1 && valeurChoix <= len(catalogue) {
				acheter(c, valeurChoix-1)
				attendreEntree(lecteur)
			} else {
				fmt.Println("Choix invalide.")
				attendreEntree(lecteur)
			}
		}
	}
}

func acheter(c *character.Character, index int) {
	objet := catalogue[index]
	if c.Or < objet.Prix {
		fmt.Println("Vous n'avez pas assez d'or.")
		return
	}

	err := c.AddInventory(objet.Nom)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	c.Or -= objet.Prix
	fmt.Printf("\nVous avez acheté %s pour %d or.\n", objet.Nom, objet.Prix)
}

func vendre(c *character.Character, lecteur *bufio.Reader) {
	outils.ClearScreen()
	ascii.AfficherASCII("BanqueASCII/echoppe.txt")

	if len(c.Inventaire) == 0 {
		lignesVide := []string{"Votre inventaire est vide."}
		outils.Encadrer(lignesVide)
		attendreEntree(lecteur)
		return
	}

	lignesVente := []string{
		"--- Vente d'objets ---",
		"Votre inventaire : " + strings.Join(c.Inventaire, " - "),
	}
	outils.Encadrer(lignesVente)

	fmt.Print("\n                                                                    Quel objet voulez-vous vendre ? ")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)

	prixVente := prixDeVente(nom)
	if prixVente == 0 {
		fmt.Println("Le marchand n'achète pas cet objet.")
		attendreEntree(lecteur)
		return
	}

	err := c.RemoveInventory(nom)
	if err != nil {
		fmt.Println("Erreur :", err)
		attendreEntree(lecteur)
		return
	}

	c.Or += prixVente
	fmt.Printf("Vous avez vendu %s pour %d or.\n", nom, prixVente)
	attendreEntree(lecteur)
}

func prixDeVente(nom string) int {
	for _, objet := range catalogue {
		if strings.EqualFold(objet.Nom, nom) {
			return objet.Prix / 2 // le marchand rachète à moitié prix
		}
	}
	return 0
}

func attendreEntree(lecteur *bufio.Reader) {
	fmt.Print("\n                                                                    Appuyez sur Entrée pour continuer...")
	lecteur.ReadString('\n')
}
