package character

import (
	"bufio"
	"fmt"
	"os"
	"scarlet/piscine"
	"strings"
)

const MaxInventaire = 10

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

func (c Character) String() string {
	return fmt.Sprintf("{%s, %s lvl %d, PV%d/%d, Inventaire %d/%d}",
		c.Nom, c.Classe, c.Niveau, c.PVActuels, c.PVMax, len(c.Inventaire), MaxInventaire)
}

func InitCharacter(niveau int, pvMax int, pvActuels int, inventaire []string) (Character, error) {
	if len(inventaire) > MaxInventaire {
		return Character{}, fmt.Errorf("inventaire trop grand : %d objets fournis, maximum autorisé %d", len(inventaire), MaxInventaire)
	}

	lecteur := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom de votre personnage : ")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)
	nom = piscine.Capitalize(nom)

	classesValides := []string{"elfe", "nain", "humain"}
	var classe string

	for {
		fmt.Print("Choisissez votre classe. Elfe, nain ou humain) : ")
		saisie, _ := lecteur.ReadString('\n')
		saisie = strings.TrimSpace(saisie)
		saisie = strings.ToLower(saisie)

		valide := false
		for _, c := range classesValides {
			if saisie == c {
				valide = true
				break
			}
		}

		if valide {
			classe = saisie
			break
		}

		fmt.Println("Classe invalide. Veuillez choisir parmi : elfe, nain, humain.")
	}

	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}, nil
}
