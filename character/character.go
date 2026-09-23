package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"scarlet/outils"
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
	XP                int // <-- Ajout de l'expérience actuelle
	XPMax             int // <-- Ajout du palier d'expérience
	PVMax             int
	PVActuels         int
	ManaMax           int
	ManaActuel        int
	Inventaire        []string
	InventaireMax     int
	UpgradesRestantes int
	Or                int
	Sorts             []string
	Equipement        Equipment
	QueteMagieNoire   int
}

func (c Character) String() string {
	return fmt.Sprintf("{%s, %s lvl %d (XP %d/%d), PV%d/%d, Mana%d/%d, Inventaire %d/%d\nÉquipement: Tête: %s | Torse: %s | Pieds: %s}",
		c.Nom, c.Classe, c.Niveau, c.XP, c.XPMax, c.PVActuels, c.PVMax, c.ManaActuel, c.ManaMax, len(c.Inventaire), c.InventaireMax,
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

	outils.ClearScreen()
	// Ajout de 30 lignes de décalage vers le bas
	fmt.Print(strings.Repeat("\n", 30))

	lignesNom := []string{"Création de personnage", "", "Entrez le nom de votre personnage :"}
	outils.Encadrer(lignesNom, 60)
	fmt.Print("                                                    Votre réponse : ")

	nom, _ := lecteur.ReadString('\n')
	nom = strings.TrimSpace(nom)
	nom = piscine.Capitalize(nom)

	classesValides := []string{"elfe", "nain", "humain"}
	var classe string

	for {
		outils.ClearScreen()
		// Ajout de 30 lignes de décalage vers le bas
		fmt.Print(strings.Repeat("\n", 30))

		lignesClasse := []string{
			"Choisissez votre classe :",
			"",
			"- Elfe",
			"- Nain",
			"- Humain",
		}
		outils.Encadrer(lignesClasse, 60)
		fmt.Print("                                                    Votre choix : ")

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

		outils.ClearScreen()
		// Ajout de 30 lignes de décalage vers le bas
		fmt.Print(strings.Repeat("\n", 30))

		lignesErreur := []string{
			"Classe invalide !",
			"Veuillez choisir parmi : elfe, nain, humain.",
		}
		outils.Encadrer(lignesErreur, 60)
		fmt.Print("\n                                                    Appuyez sur Entrée pour continuer...")
		lecteur.ReadString('\n')
	}

	pvMax, pvActuels := pvSelonClasse(classe)
	manaMax := manaSelonClasse(classe)

	return Character{
		Nom:               nom,
		Classe:            classe,
		Niveau:            niveau,
		XP:                0,   // <-- Initialisé à 0
		XPMax:             100, // <-- Initialisé à 100
		PVMax:             pvMax,
		PVActuels:         pvActuels,
		ManaMax:           manaMax,
		ManaActuel:        manaMax,
		Inventaire:        inventaire,
		InventaireMax:     InventaireBase,
		UpgradesRestantes: MaxUpgrades,
		Or:                OrDepart,
		Sorts:             []string{},
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
