package character

import (
	"bufio"
	"fmt"
	"os"
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

// InitCharacter demande au joueur son nom et sa classe via le terminal,
// puis crée un Character avec les autres valeurs passées en paramètres.
func InitCharacter(niveau int, pvMax int, pvActuels int, inventaire []string) (Character, error) {
	if len(inventaire) > MaxInventaire {
		return Character{}, fmt.Errorf("inventaire trop grand : %d objets fournis, maximum autorisé %d", len(inventaire), MaxInventaire)
	}

	lecteur := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom de votre personnage : ")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom) // enlève le retour à la ligne et les espaces en trop

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
