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
}

// Echoppe ouvre le menu de la boutique du marchand : achat et vente d'objets.
func Echoppe(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n--- Boutique du marchand --- (Or : %d)\n", c.Or)
		fmt.Println("Jetez un oeil à mes marchandises !")
		for i, objet := range catalogue {
			fmt.Printf("%d. Acheter %s - %d or\n", i+1, objet.Nom, objet.Prix)
		}
		fmt.Println("v. Vendre un objet")
		fmt.Println("0. Quitter la boutique")
		fmt.Print("Votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "0":
			return
		case "v":
			vendre(c, lecteur)
		default:
			acheter(c, choix)
		}
	}
}

func acheter(c *character.Character, choix string) {
	num, err := strconv.Atoi(choix)
	if err != nil || num < 1 || num > len(catalogue) {
		fmt.Println("Choix invalide.")
		return
	}

	objet := catalogue[num-1]
	if c.Or < objet.Prix {
		fmt.Println("Vous n'avez pas assez d'or.")
		return
	}

	err = c.AddInventory(objet.Nom)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	c.Or -= objet.Prix
	outils.ClearScreen()
	ascii.AfficherASCII("BanqueASCII/echoppe_marchand.txt")
	fmt.Printf("Vous avez acheté %s pour %d or.\n", objet.Nom, objet.Prix)
}

func vendre(c *character.Character, lecteur *bufio.Reader) {
	if len(c.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Votre inventaire :", strings.Join(c.Inventaire, " - "))
	fmt.Print("Quel objet voulez-vous vendre ? Ecrivez le en toute lettre")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)

	prixVente := prixDeVente(nom)
	if prixVente == 0 {
		fmt.Println("Le marchand n'achète pas cet objet.")
		return
	}

	err := c.RemoveInventory(nom)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	c.Or += prixVente
	fmt.Printf("Vous avez vendu %s pour %d or.\n", nom, prixVente)
}

func prixDeVente(nom string) int {
	for _, objet := range catalogue {
		if objet.Nom == nom {
			return objet.Prix / 2 // le marchand rachète à moitié prix
		}
	}
	return 0
}
