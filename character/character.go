package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/piscine"
)

const InventaireBase = 10
const OrDepart = 100
const MaxUpgrades = 3
const CoutUpgrade = 30
const BonusUpgrade = 10

type Character struct {
	Nom               string
	Classe            string
	Niveau            int
	PVMax             int
	PVActuels         int
	ManaMax           int
	ManaActuel        int
	Inventaire        []string
	InventaireMax     int
	UpgradesRestantes int
	Or                int
	Equipement        Equipment
}

func (c Character) String() string {
	return fmt.Sprintf("{%s, %s lvl %d, PV%d/%d, Mana%d/%d, Inventaire %d/%d\nÉquipement: Tête: %s | Torse: %s | Pieds: %s}",
		c.Nom, c.Classe, c.Niveau, c.PVActuels, c.PVMax, c.ManaActuel, c.ManaMax, len(c.Inventaire), c.InventaireMax,
		afficherEmplacement(c.Equipement.Tete),
		afficherEmplacement(c.Equipement.Torse),
		afficherEmplacement(c.Equipement.Pieds))
}

func afficherEmplacement(objet string) string {
	if objet == "" {
		return "Aucun"
	}
	return objet
}

func InitCharacter(niveau int, inventaire []string) (Character, error) {
	if len(inventaire) > InventaireBase {
		return Character{}, fmt.Errorf("inventaire trop grand : %d objets fournis, maximum autorisé %d", len(inventaire), InventaireBase)
	}

	lecteur := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom de votre personnage : ")
	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)
	nom = piscine.Capitalize(nom)

	classesValides := []string{"elfe", "nain", "humain"}
	var classe string

	for {
		fmt.Print("Choisissez votre classe (elfe / nain / humain) : ")
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

	pvMax, pvActuels := pvSelonClasse(classe)
	manaMax := manaSelonClasse(classe)

	return Character{
		Nom:               nom,
		Classe:            classe,
		Niveau:            niveau,
		PVMax:             pvMax,
		PVActuels:         pvActuels,
		ManaMax:           manaMax,
		ManaActuel:        manaMax,
		Inventaire:        inventaire,
		InventaireMax:     InventaireBase,
		UpgradesRestantes: MaxUpgrades,
		Or:                OrDepart,
		Equipement:        Equipment{},
	}, nil
}

func pvSelonClasse(classe string) (int, int) {
	switch classe {
	case "elfe":
		return 80, 40
	case "nain":
		return 120, 60
	case "humain":
		return 100, 50
	default:
		return 100, 50
	}
}

func manaSelonClasse(classe string) int {
	switch classe {
	case "elfe":
		return 120
	case "nain":
		return 80
	case "humain":
		return 100
	default:
		return 100
	}
}
