package forge

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

// Recette représente un équipement fabricable, avec ses composants requis et son coût en or.
type Recette struct {
	Nom        string
	Composants map[string]int
	Prix       int
}

var recettes = []Recette{
	{
		Nom: "Chapeau de l'aventurier",
		Composants: map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		},
		Prix: 10,
	},
	{
		Nom: "Tunique de l'aventurier",
		Composants: map[string]int{
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		},
		Prix: 10,
	},
	{
		Nom: "Bottes de l'aventurier",
		Composants: map[string]int{
			"Fourrure de Loup": 1,
			"Cuir de Sanglier": 1,
		},
		Prix: 10,
	},
}

// Fabriquer ouvre le menu de fabrication d'objets chez le forgeron.
func Fabriquer(c *character.Character) {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		outils.ClearScreen()
		ascii.AfficherASCII("BanqueASCII/forgeron.txt")

		var lignes []string
		lignes = append(lignes, fmt.Sprintf("--- Atelier de Forge --- (Or : %d)", c.Or))
		lignes = append(lignes, "Que voulez-vous fabriquer ?")
		lignes = append(lignes, "")

		// Génération de la liste des recettes
		for i, r := range recettes {
			lignes = append(lignes, fmt.Sprintf("%d. %s (%d or)", i+1, r.Nom, r.Prix))
			for composant, quantite := range r.Composants {
				lignes = append(lignes, fmt.Sprintf("    - %d x %s", quantite, composant))
			}
			lignes = append(lignes, "")
		}

		lignes = append(lignes, "0. Retour")

		// Encadrer la liste des recettes
		outils.Encadrer(lignes, 90)

		fmt.Print("                                                                    Votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		if choix == "0" {
			return
		}

		num, err := strconv.Atoi(choix)
		if err != nil || num < 1 || num > len(recettes) {
			fmt.Println("Choix invalide.")
			attendreEntree(lecteur)
			continue
		}

		fabriquerObjet(c, recettes[num-1])
		attendreEntree(lecteur)
	}
}

func fabriquerObjet(c *character.Character, r Recette) {
	// Vérification de l'or
	if c.Or < r.Prix {
		fmt.Printf("\nVous n'avez pas assez d'or pour fabriquer %s (%d or nécessaires, vous en possédez %d).\n", r.Nom, r.Prix, c.Or)
		return
	}

	// Vérification des composants, un par un
	for composant, quantiteRequise := range r.Composants {
		quantitePossedee := compterOccurrences(c.Inventaire, composant)
		if quantitePossedee < quantiteRequise {
			fmt.Printf("\nIl vous manque des composants pour fabriquer %s : %s (besoin de %d, vous en avez %d).\n",
				r.Nom, composant, quantiteRequise, quantitePossedee)
			return
		}
	}

	// Vérification de la place dans l'inventaire (composants consommés, 1 objet ajouté)
	placeApresFabrication := len(c.Inventaire) - totalComposants(r) + 1
	if placeApresFabrication > c.InventaireMax {
		fmt.Println("\nVotre inventaire n'aura pas assez de place pour récupérer l'objet fabriqué.")
		return
	}

	// Tout est bon : on consomme les composants
	for composant, quantite := range r.Composants {
		for i := 0; i < quantite; i++ {
			c.RemoveInventory(composant)
		}
	}

	c.Or -= r.Prix
	c.AddInventory(r.Nom)

	fmt.Printf("\nVous avez fabriqué : %s !\n", r.Nom)
}

// compterOccurrences compte combien de fois un objet apparaît dans l'inventaire.
func compterOccurrences(inventaire []string, item string) int {
	count := 0
	for _, objet := range inventaire {
		if objet == item {
			count++
		}
	}
	return count
}

// totalComposants additionne les quantités de tous les composants d'une recette.
func totalComposants(r Recette) int {
	total := 0
	for _, quantite := range r.Composants {
		total += quantite
	}
	return total
}

func attendreEntree(lecteur *bufio.Reader) {
	fmt.Print("\n                                                                    Appuyez sur Entrée pour continuer...")
	lecteur.ReadString('\n')
}
