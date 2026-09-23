package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"scarlet/character"
	"scarlet/outils"
)

// ----------------------------------------------------------------------------
// STRUCTURES
// ----------------------------------------------------------------------------

type Attaque struct {
	Nom    string
	Chance int // en pourcentage, la somme des Chance d'un monstre doit faire 100
	Degats int
}

type Monstre struct {
	Nom          string
	FichierASCII string // nom du fichier dans BanqueASCII, ex: "gobelin.txt"
	PVMax        int
	PVActuels    int
	Attaques     []Attaque
	ToursGeles   int
}

// ----------------------------------------------------------------------------
// INITIALISATION DU GOBELIN
// ----------------------------------------------------------------------------

func InitGobelin() Monstre {
	return Monstre{
		Nom:          "Gobelin d'entrainement",
		FichierASCII: "gobelin.txt", // <-- vérifie que ce nom correspond bien au fichier dans BanqueASCII/
		PVMax:        40,
		PVActuels:    40,
		Attaques: []Attaque{
			{Nom: "Griffure", Chance: 50, Degats: 15},
			{Nom: "Morsure", Chance: 20, Degats: 25},
			{Nom: "Coup de bâton", Chance: 15, Degats: 10},
			{Nom: "Jet de pierre", Chance: 10, Degats: 20},
			{Nom: "Cri sauvage", Chance: 5, Degats: 30},
		},
	}
}

// choisirAttaque tire une attaque au hasard selon les probabilités définies.
func choisirAttaque(attaques []Attaque) Attaque {
	tirage := rand.Intn(100)
	cumul := 0

	for _, a := range attaques {
		cumul += a.Chance
		if tirage < cumul {
			return a
		}
	}
	return attaques[len(attaques)-1]
}

// ----------------------------------------------------------------------------
// AFFICHAGE CÔTE À CÔTE (STYLE POKÉMON)
// ----------------------------------------------------------------------------

const largeurColonne = 40 // largeur réservée à l'ASCII de gauche, ajuste selon tes fichiers

// chargerLignes lit un fichier ASCII et renvoie son contenu ligne par ligne.
func chargerLignes(chemin string) []string {
	data, err := os.ReadFile(chemin)
	if err != nil {
		return []string{"[ASCII introuvable : " + chemin + "]"}
	}
	return strings.Split(string(data), "\n")
}

// AfficherArene affiche le joueur à gauche et le monstre à droite, avec leurs PV,
// façon écran de combat Pokémon.
func AfficherArene(c *character.Character, m *Monstre) {
	lignesJoueur := chargerLignes("BanqueASCII/" + c.Classe + ".txt")
	lignesMonstre := chargerLignes("BanqueASCII/" + m.FichierASCII)

	maxLignes := len(lignesJoueur)
	if len(lignesMonstre) > maxLignes {
		maxLignes = len(lignesMonstre)
	}

	fmt.Printf("%-*s   %s\n", largeurColonne, c.Nom, m.Nom)
	fmt.Printf("%-*s   PV : %d/%d\n", largeurColonne, fmt.Sprintf("PV : %d/%d", c.PVActuels, c.PVMax), m.PVActuels, m.PVMax)
	fmt.Println()

	for i := 0; i < maxLignes; i++ {
		gauche := ""
		if i < len(lignesJoueur) {
			gauche = lignesJoueur[i]
		}
		droite := ""
		if i < len(lignesMonstre) {
			droite = lignesMonstre[i]
		}
		fmt.Printf("%-*s   %s\n", largeurColonne, gauche, droite)
	}
}

// ----------------------------------------------------------------------------
// SORTS
// ----------------------------------------------------------------------------

var coutMana = map[string]int{
	"Boule de Feu":   10,
	"Glace":          15,
	"Lumière Divine": 20,
}

func lancerSort(c *character.Character, monstre *Monstre, nomSort string) bool {
	cout, existe := coutMana[nomSort]
	if !existe {
		fmt.Println("Sort inconnu.")
		return false
	}
	if c.ManaActuel < cout {
		fmt.Println("Pas assez de mana pour lancer ce sort.")
		return false
	}

	switch nomSort {
	case "Boule de Feu":
		degats := 15
		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("\n%s lance Boule de Feu et inflige %d dégâts à %s !\n", c.Nom, degats, monstre.Nom)

	case "Glace":
		monstre.ToursGeles = 2
		fmt.Printf("\n%s lance Glace ! %s est gelé et ne pourra pas agir pendant 2 tours.\n", c.Nom, monstre.Nom)

	case "Lumière Divine":
		soin := 40
		c.PVActuels += soin
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Printf("\n%s invoque Lumière Divine et récupère %d PV !\n", c.Nom, soin)

	default:
		fmt.Println("Ce sort n'a pas encore d'effet défini.")
		return false
	}

	c.ManaActuel -= cout
	fmt.Printf("PV : %d/%d | Mana : %d/%d\n", c.PVActuels, c.PVMax, c.ManaActuel, c.ManaMax)
	time.Sleep(1 * time.Second)
	return true
}

// ----------------------------------------------------------------------------
// TOUR DU MONSTRE
// ----------------------------------------------------------------------------

func TourMonstre(monstre *Monstre, c *character.Character) {
	fmt.Printf("\n--- Tour de %s ---\n", monstre.Nom)

	if monstre.ToursGeles > 0 {
		fmt.Printf("%s est gelé et ne peut pas agir (%d tour(s) restant(s)).\n", monstre.Nom, monstre.ToursGeles)
		monstre.ToursGeles--
		time.Sleep(1 * time.Second)
		return
	}

	attaque := choisirAttaque(monstre.Attaques)

	c.PVActuels -= attaque.Degats
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}

	fmt.Printf("%s utilise %s et inflige %d dégâts à %s.\n", monstre.Nom, attaque.Nom, attaque.Degats, c.Nom)
	fmt.Printf("%s PV : %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
	time.Sleep(1 * time.Second)
}

// ----------------------------------------------------------------------------
// TOUR DU JOUEUR
// ----------------------------------------------------------------------------

func TourJoueur(c *character.Character, monstre *Monstre, lecteur *bufio.Reader) {
	fmt.Printf("\n--- Tour de %s (%s) ---\n", c.Nom, c.Classe)

	for {
		fmt.Println("\n=== MENU COMBAT ===")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser une potion")
		if len(c.Sorts) > 0 {
			fmt.Println("3. Lancer un sort")
		}
		fmt.Print("Choisissez une action : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			degats := 5
			monstre.PVActuels -= degats
			if monstre.PVActuels < 0 {
				monstre.PVActuels = 0
			}
			fmt.Printf("\n%s attaque et inflige %d dégâts à %s !\n", c.Nom, degats, monstre.Nom)
			fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)
			time.Sleep(1 * time.Second)
			return

		case "2":
			err := c.TakePot()
			if err != nil {
				fmt.Println("Erreur :", err)
				continue
			}
			return

		case "3":
			if len(c.Sorts) == 0 {
				fmt.Println("Choix invalide.")
				continue
			}
			fmt.Println("Sorts connus :")
			for i, sort := range c.Sorts {
				fmt.Printf("%d. %s (coût : %d mana)\n", i+1, sort, coutMana[sort])
			}
			fmt.Print("Quel sort voulez-vous lancer ? ")

			choixSort, _ := lecteur.ReadString('\n')
			choixSort = strings.TrimSpace(choixSort)

			num, err := strconv.Atoi(choixSort)
			if err != nil || num < 1 || num > len(c.Sorts) {
				fmt.Println("Choix invalide.")
				continue
			}

			if lancerSort(c, monstre, c.Sorts[num-1]) {
				return
			}

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// ----------------------------------------------------------------------------
// BOUCLE PRINCIPALE DE COMBAT
// ----------------------------------------------------------------------------

func TrainingFight(c *character.Character) {
	monstre := InitGobelin()
	tour := 1
	lecteur := bufio.NewReader(os.Stdin)

	outils.ClearScreen()
	fmt.Println("========================================")
	fmt.Printf("   DEBUT DU COMBAT : %s vs %s\n", c.Nom, monstre.Nom)
	fmt.Println("========================================")

	for c.PVActuels > 0 && monstre.PVActuels > 0 {
		outils.ClearScreen()
		fmt.Printf("\n========== TOUR %d ==========\n", tour)
		AfficherArene(c, &monstre)

		TourJoueur(c, &monstre, lecteur)
		if monstre.PVActuels <= 0 {
			fmt.Printf("\n🎉 Félicitations ! Vous avez vaincu %s !\n", monstre.Nom)
			break
		}

		TourMonstre(&monstre, c)
		if c.PVActuels <= 0 {
			fmt.Printf("\n☠️ Vous avez été vaincu par %s...\n", monstre.Nom)
			break
		}

		tour++
	}

	if c.PVActuels <= 0 {
		fmt.Println("Vous êtes ressuscité avec 50% de vos PV max.")
		c.PVActuels = c.PVMax / 2
	}
}
