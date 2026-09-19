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
	inventaireTexte := strings.Join(c.Inventaire, " ")
	return fmt.Sprintf("{%s, %s lvl %d, PV%d/%d [%s]}",
		c.Nom, c.Classe, c.Niveau, c.PVActuels, c.PVMax, inventaireTexte)
}

func InitCharacter(niveau int, pvMax int, pvActuels int, inventaire []string) (Character, error) {
	if len(inventaire) > MaxInventaire {
		return Character{}, fmt.Errorf("inventaire trop grand : %d objets fournis, maximum autorisé %d", len(inventaire), MaxInventaire)
	}

	lecteur := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom de votre personnage : ")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)
	nom = piscine.Capitalize(nom) // "rémi" ou "RÉMI" → "Rémi"

	fmt.Print("Entrez la classe de votre personnage : ")
	classe, _ := lecteur.ReadString('\n')
	classe = strings.TrimSpace(classe)

	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}, nil
}
