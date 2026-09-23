package character

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"scarlet/outils"
)

// AccessInventory affiche le contenu de l'inventaire, puis propose un menu
// d'actions selon les objets présents (fermer l'inventaire, boire une potion,
// apprendre un sort...)
func (c *Character) AccessInventory() {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		var menuLignes []string

		// Préparation de l'affichage de l'inventaire
		if len(c.Inventaire) == 0 {
			menuLignes = append(menuLignes, "Inventaire vide.")
		} else {
			detail := strings.Join(c.Inventaire, " - ")
			menuLignes = append(menuLignes, "Inventaire : "+detail)
		}

		auMoinsUnePotion := false
		var livresTrouves []string
		for _, objet := range c.Inventaire {
			if objet == "Potion de soins" {
				auMoinsUnePotion = true
			}
			if strings.HasPrefix(objet, PrefixeLivreSort) {
				livresTrouves = append(livresTrouves, objet)
			}
		}

		// Préparation des options du menu
		menuLignes = append(menuLignes, "")
		menuLignes = append(menuLignes, "Que voulez-vous faire ?")
		menuLignes = append(menuLignes, "1. Fermer l'inventaire")
		if auMoinsUnePotion {
			menuLignes = append(menuLignes, "2. Boire une potion de soins")
		}
		if len(livresTrouves) > 0 {
			menuLignes = append(menuLignes, "3. Apprendre un sort")
		}

		// Affichage du cadre style Undertale
		outils.Encadrer(menuLignes, 55)

		fmt.Print("Selon votre choix, tapez le numéro puis entrée : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			return // sort de la fonction, referme l'inventaire

		case "2":
			if !auMoinsUnePotion {
				fmt.Println("Choix invalide.")
				fmt.Println()
				continue
			}
			err := c.TakePot()
			if err != nil {
				fmt.Println("Erreur :", err)
			}
			fmt.Println() // ligne vide pour aérer avant de réafficher le menu

		case "3":
			if len(livresTrouves) == 0 {
				fmt.Println("Choix invalide.")
				fmt.Println()
				continue
			}
			c.choisirLivreAApprendre(livresTrouves, lecteur)
			fmt.Println()

		default:
			fmt.Println("Choix invalide.")
			fmt.Println()
		}
	}
}

// choisirLivreAApprendre liste les livres de sort disponibles dans
// l'inventaire et demande au joueur lequel il souhaite étudier.
func (c *Character) choisirLivreAApprendre(livres []string, lecteur *bufio.Reader) {
	var lignes []string
	lignes = append(lignes, "Livres disponibles :")
	for i, livre := range livres {
		lignes = append(lignes, fmt.Sprintf("%d. %s", i+1, livre))
	}

	outils.Encadrer(lignes, 55)

	fmt.Print("Quel livre voulez-vous étudier ? ")

	choix, _ := lecteur.ReadString('\n')
	choix = strings.TrimSpace(choix)

	num, err := strconv.Atoi(choix)
	if err != nil || num < 1 || num > len(livres) {
		fmt.Println("Choix invalide.")
		return
	}

	err = c.ApprendreSort(livres[num-1])
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}
