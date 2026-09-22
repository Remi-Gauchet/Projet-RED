package character

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// infoEquipement décrit l'emplacement et le bonus de PV d'un équipement.
type infoEquipement struct {
	Slot  string
	Bonus int
}

var statsEquipement = map[string]infoEquipement{
	"Chapeau de l'aventurier": {Slot: "Tete", Bonus: 10},
	"Tunique de l'aventurier": {Slot: "Torse", Bonus: 25},
	"Bottes de l'aventurier":  {Slot: "Pieds", Bonus: 15},
}

// GererEquipement permet au joueur d'équiper ou de retirer un équipement.
func (c *Character) GererEquipement() {
	lecteur := bufio.NewReader(os.Stdin)

	for {
		var equipables []string
		for _, objet := range c.Inventaire {
			if _, ok := statsEquipement[objet]; ok {
				equipables = append(equipables, objet)
			}
		}

		fmt.Println("\n--- Équipement ---")
		fmt.Printf("Tête: %s | Torse: %s | Pieds: %s\n",
			afficherEmplacement(c.Equipement.Tete),
			afficherEmplacement(c.Equipement.Torse),
			afficherEmplacement(c.Equipement.Pieds))

		if len(equipables) == 0 {
			fmt.Println("Vous n'avez rien à vous mettre.")
		} else {
			fmt.Println("\nObjets équipables dans votre inventaire :")
			for i, objet := range equipables {
				fmt.Printf("%d. Équiper %s\n", i+1, objet)
			}
		}

		fmt.Println("r. Retirer un équipement")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "0":
			return
		case "r":
			retirerEquipement(c, lecteur)
		default:
			num, err := strconv.Atoi(choix)
			if err != nil || num < 1 || num > len(equipables) {
				fmt.Println("Choix invalide.")
				continue
			}
			equiperObjet(c, equipables[num-1])
		}
	}
}

func equiperObjet(c *Character, nom string) {
	info := statsEquipement[nom]

	var emplacement *string
	switch info.Slot {
	case "Tete":
		emplacement = &c.Equipement.Tete
	case "Torse":
		emplacement = &c.Equipement.Torse
	case "Pieds":
		emplacement = &c.Equipement.Pieds
	}

	// Si un objet occupe déjà cet emplacement, on le retire d'abord (retour à l'inventaire)
	if *emplacement != "" {
		ancienneInfo := statsEquipement[*emplacement]
		c.PVMax -= ancienneInfo.Bonus
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		c.AddInventory(*emplacement)
	}

	err := c.RemoveInventory(nom)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	*emplacement = nom
	c.PVMax += info.Bonus
	c.PVActuels += info.Bonus

	fmt.Printf("Vous avez équipé : %s (+%d PV max)\n", nom, info.Bonus)
}

func retirerEquipement(c *Character, lecteur *bufio.Reader) {
	fmt.Println("\nQuel emplacement voulez-vous vider ?")
	fmt.Println("1. Tête")
	fmt.Println("2. Torse")
	fmt.Println("3. Pieds")
	fmt.Print("Votre choix : ")

	choix, _ := lecteur.ReadString('\n')
	choix = strings.TrimSpace(choix)

	var emplacement *string
	switch choix {
	case "1":
		emplacement = &c.Equipement.Tete
	case "2":
		emplacement = &c.Equipement.Torse
	case "3":
		emplacement = &c.Equipement.Pieds
	default:
		fmt.Println("Choix invalide.")
		return
	}

	if *emplacement == "" {
		fmt.Println("Rien n'est équipé à cet emplacement.")
		return
	}

	info := statsEquipement[*emplacement]

	err := c.AddInventory(*emplacement)
	if err != nil {
		fmt.Println("Impossible de retirer l'équipement :", err)
		return
	}

	c.PVMax -= info.Bonus
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}

	fmt.Printf("Vous avez retiré : %s (-%d PV max)\n", *emplacement, info.Bonus)
	*emplacement = ""
}
