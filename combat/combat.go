package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"scarlet/audio"
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
	FichierASCII string
	PVMax        int
	PVActuels    int
	Attaques     []Attaque
	ToursGeles   int
}

// ----------------------------------------------------------------------------
// INITIALISATION DES MONSTRES
// ----------------------------------------------------------------------------

// InitGobelin crée le gobelin d'entrainement.
func InitGobelin() Monstre {
	return Monstre{
		Nom:          "Gobelin d'entrainement",
		FichierASCII: "gobelin.txt",
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

// InitRoiGobelin crée le Roi Gobelin, plus puissant que le gobelin de base.
func InitRoiGobelin() Monstre {
	return Monstre{
		Nom:          "Roi Gobelin",
		FichierASCII: "roi_gobelin.txt", // <-- à vérifier
		PVMax:        90,
		PVActuels:    90,
		Attaques: []Attaque{
			{Nom: "Coup de Sceptre", Chance: 40, Degats: 20},
			{Nom: "Charge Royale", Chance: 25, Degats: 30},
			{Nom: "Piétinement", Chance: 20, Degats: 15},
			{Nom: "Ordre de Massacre", Chance: 10, Degats: 35},
			{Nom: "Fureur du Roi", Chance: 5, Degats: 45},
		},
	}
}

// InitDemon crée le Démon, monstre de milieu de jeu.
func InitDemon() Monstre {
	return Monstre{
		Nom:          "Démon",
		FichierASCII: "demon.txt", // <-- à vérifier
		PVMax:        150,
		PVActuels:    150,
		Attaques: []Attaque{
			{Nom: "Griffe Infernale", Chance: 35, Degats: 25},
			{Nom: "Souffle de Soufre", Chance: 25, Degats: 35},
			{Nom: "Poing Ténébreux", Chance: 20, Degats: 30},
			{Nom: "Flammes Abyssales", Chance: 15, Degats: 45},
			{Nom: "Rituel de Douleur", Chance: 5, Degats: 60},
		},
	}
}

// InitDragon crée le Dragon, boss le plus puissant du trio.
func InitDragon() Monstre {
	return Monstre{
		Nom:          "Dragon",
		FichierASCII: "dragon.txt", // <-- à vérifier
		PVMax:        250,
		PVActuels:    250,
		Attaques: []Attaque{
			{Nom: "Morsure Écailleuse", Chance: 35, Degats: 30},
			{Nom: "Coup de Queue", Chance: 25, Degats: 40},
			{Nom: "Griffure Ardente", Chance: 20, Degats: 45},
			{Nom: "Souffle de Feu", Chance: 15, Degats: 60},
			{Nom: "Rugissement Dévastateur", Chance: 5, Degats: 80},
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
// AFFICHAGE ASCII
// ----------------------------------------------------------------------------

const largeurColonne = 40 // largeur du bloc joueur (gauche) et du bloc monstre (droite)
const largeurBoite = 90   // largeur intérieure totale du rectangle

func chargerLignes(chemin string) []string {
	data, err := os.ReadFile(chemin)
	if err != nil {
		return []string{"[ASCII introuvable : " + chemin + "]"}
	}
	return strings.Split(string(data), "\n")
}

func afficherASCIIBrut(chemin string) {
	for _, ligne := range chargerLignes(chemin) {
		fmt.Println(ligne)
	}
}

// encadrer affiche une liste de lignes à l'intérieur d'un rectangle façon
// écran de combat Undertale.
func encadrer(lignes []string) {
	fmt.Println("┌" + strings.Repeat("─", largeurBoite+2) + "┐")
	for _, l := range lignes {
		rc := utf8.RuneCountInString(l)
		if rc > largeurBoite {
			runes := []rune(l)
			l = string(runes[:largeurBoite])
			rc = largeurBoite
		}
		fmt.Printf("│ %s%s │\n", l, strings.Repeat(" ", largeurBoite-rc))
	}
	fmt.Println("└" + strings.Repeat("─", largeurBoite+2) + "┘")
}

// construireBlocJoueur construit le bloc de gauche : image du joueur, ses PV,
// puis la liste de ses actions possibles en combat.
func construireBlocJoueur(c *character.Character) []string {
	bloc := chargerLignes("BanqueASCII/" + c.Classe + ".txt")
	bloc = append(bloc, "")
	bloc = append(bloc, c.Nom)
	bloc = append(bloc, fmt.Sprintf("PV : %d/%d", c.PVActuels, c.PVMax))
	bloc = append(bloc, "")
	bloc = append(bloc, "Actions possibles :")
	bloc = append(bloc, "- Coup Risqué (50% de chances, 20 dégâts)")
	bloc = append(bloc, "- Frappe Rapide (10 dégâts)")
	for _, sort := range c.Sorts {
		bloc = append(bloc, fmt.Sprintf("- %s (coût : %d mana)", sort, coutMana[sort]))
	}
	return bloc
}

// construireBlocMonstre construit le bloc de droite : image du monstre, ses
// PV, puis la liste de ses attaques possibles.
func construireBlocMonstre(m *Monstre) []string {
	bloc := chargerLignes("BanqueASCII/" + m.FichierASCII)
	bloc = append(bloc, "")
	bloc = append(bloc, m.Nom)
	bloc = append(bloc, fmt.Sprintf("PV : %d/%d", m.PVActuels, m.PVMax))
	bloc = append(bloc, "")
	bloc = append(bloc, "Attaques possibles :")
	for _, a := range m.Attaques {
		bloc = append(bloc, fmt.Sprintf("- %s (%d%% de chances, %d dégâts)", a.Nom, a.Chance, a.Degats))
	}
	return bloc
}

// AfficherArene construit le bloc joueur et le bloc monstre, les fusionne
// ligne par ligne à la même hauteur, puis affiche le tout dans un rectangle.
func AfficherArene(c *character.Character, m *Monstre) {
	blocJoueur := construireBlocJoueur(c)
	blocMonstre := construireBlocMonstre(m)

	maxLignes := len(blocJoueur)
	if len(blocMonstre) > maxLignes {
		maxLignes = len(blocMonstre)
	}

	var contenu []string
	for i := 0; i < maxLignes; i++ {
		gauche := ""
		if i < len(blocJoueur) {
			gauche = blocJoueur[i]
		}
		droite := ""
		if i < len(blocMonstre) {
			droite = blocMonstre[i]
		}
		contenu = append(contenu, fmt.Sprintf("%-*s   %s", largeurColonne, gauche, droite))
	}

	encadrer(contenu)
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
		chemin := "BanqueSon/" + c.Classe + "/bouledefeu.ogg"
		audio.PlaySoundBlocking(chemin)
		degats := 15 * c.Niveau
		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("\n%s lance Boule de Feu et inflige %d dégâts à %s !\n", c.Nom, degats, monstre.Nom)

	case "Glace":
		monstre.ToursGeles = 2
		fmt.Printf("\n%s lance Glace ! %s est gelé et ne pourra pas agir pendant 2 tours.\n", c.Nom, monstre.Nom)

	case "Lumière Divine":
		soin := 40 * c.Niveau
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
	outils.ClearScreen()
	AfficherArene(c, monstre)
	fmt.Printf("\n--- Tour de %s ---\n", monstre.Nom)

	if monstre.ToursGeles > 0 {
		fmt.Printf("%s est gelé et ne peut pas agir (%d tour(s) restant(s)).\n", monstre.Nom, monstre.ToursGeles)
		monstre.ToursGeles--
		time.Sleep(1500 * time.Millisecond)
		return
	}

	attaque := choisirAttaque(monstre.Attaques)

	c.PVActuels -= attaque.Degats
	if c.PVActuels < 0 {
		c.PVActuels = 0
	}

	fmt.Printf("%s utilise %s et inflige %d dégâts à %s.\n", monstre.Nom, attaque.Nom, attaque.Degats, c.Nom)
	fmt.Printf("%s PV : %d/%d\n", c.Nom, c.PVActuels, c.PVMax)
	time.Sleep(1500 * time.Millisecond)
}

// ----------------------------------------------------------------------------
// TOUR DU JOUEUR
// ----------------------------------------------------------------------------

func TourJoueur(c *character.Character, monstre *Monstre, lecteur *bufio.Reader) {
	for {
		outils.ClearScreen()
		AfficherArene(c, monstre)
		fmt.Printf("\n--- Tour de %s (%s) ---\n", c.Nom, c.Classe)

		fmt.Println("\n=== MENU COMBAT ===")
		fmt.Println("1. Coup Risqué (50% de chance, 20 dégâts si réussi)")
		fmt.Println("2. Frappe Rapide (10 dégâts assurés)")
		fmt.Println("3. Utiliser une potion")
		if len(c.Sorts) > 0 {
			fmt.Println("4. Lancer un sort")
		}
		fmt.Print("Choisissez une action : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			fmt.Printf("\n%s tente un Coup Risqué...\n", c.Nom)
			time.Sleep(500 * time.Millisecond)

			if rand.Intn(2) == 0 {
				fmt.Println("Raté ! Vous perdez votre tour.")
			} else {
				degats := 20 * c.Niveau
				monstre.PVActuels -= degats
				if monstre.PVActuels < 0 {
					monstre.PVActuels = 0
				}
				fmt.Printf("Coup critique ! Vous infligez %d dégâts à %s !\n", degats, monstre.Nom)
				fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)
			}
			time.Sleep(1500 * time.Millisecond)
			return

		case "2":
			degats := 10 * c.Niveau
			monstre.PVActuels -= degats
			if monstre.PVActuels < 0 {
				monstre.PVActuels = 0
			}
			fmt.Printf("\n%s utilise Frappe Rapide et inflige %d dégâts à %s !\n", c.Nom, degats, monstre.Nom)
			fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)
			time.Sleep(1500 * time.Millisecond)
			return

		case "3":
			err := c.TakePot()
			if err != nil {
				fmt.Println("Erreur :", err)
				time.Sleep(1500 * time.Millisecond)
				continue
			}
			time.Sleep(1500 * time.Millisecond)
			return

		case "4":
			if len(c.Sorts) == 0 {
				fmt.Println("Choix invalide.")
				time.Sleep(1 * time.Second)
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
				time.Sleep(1 * time.Second)
				continue
			}

			if lancerSort(c, monstre, c.Sorts[num-1]) {
				time.Sleep(500 * time.Millisecond)
				return
			}
			time.Sleep(1 * time.Second)

		default:
			fmt.Println("Choix invalide.")
			time.Sleep(1 * time.Second)
		}
	}
}

// ----------------------------------------------------------------------------
// ÉCRAN DE MORT
// ----------------------------------------------------------------------------

const coutResurrection = 10

func GererMort(c *character.Character, lecteur *bufio.Reader) {
	outils.ClearScreen()
	afficherASCIIBrut("BanqueASCII/mort.txt")
	fmt.Println("\nMême les rois et les héros finissent par offrir leur sourire le plus misérable à ce monde impitoyable qui contemple leur déchéance avec une indifférence glaciale.")

	for {
		fmt.Println("\n1. Ressusciter à l'auberge (Or -10, PV au maximum)")
		fmt.Println("2. Quitter le jeu")
		fmt.Print("Votre choix : ")

		choix, _ := lecteur.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			if c.Or < coutResurrection {
				c.Or = 0
			} else {
				c.Or -= coutResurrection
			}
			c.PVActuels = c.PVMax
			fmt.Printf("\nVous reprenez connaissance à l'auberge, en pleine forme. (Or restant : %d)\n", c.Or)
			return

		case "2":
			fmt.Println("\nMerci d'avoir joué. À bientôt !")
			os.Exit(0)

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// ----------------------------------------------------------------------------
// BOUCLE PRINCIPALE DE COMBAT (GÉNÉRIQUE, TOUS MONSTRES)
// ----------------------------------------------------------------------------

func Combattre(c *character.Character, monstre Monstre) {
	lecteur := bufio.NewReader(os.Stdin)

	outils.ClearScreen()
	fmt.Println("========================================")
	fmt.Printf("   DEBUT DU COMBAT : %s vs %s\n", c.Nom, monstre.Nom)
	fmt.Println("========================================")
	time.Sleep(1500 * time.Millisecond)

	for c.PVActuels > 0 && monstre.PVActuels > 0 {
		TourJoueur(c, &monstre, lecteur)
		if monstre.PVActuels <= 0 {
			outils.ClearScreen()
			AfficherArene(c, &monstre)
			fmt.Printf("\n🎉 Félicitations ! Vous avez vaincu %s !\n", monstre.Nom)
			break
		}

		TourMonstre(&monstre, c)
		if c.PVActuels <= 0 {
			outils.ClearScreen()
			AfficherArene(c, &monstre)
			fmt.Printf("\n☠️ Vous avez été vaincu par %s...\n", monstre.Nom)
			break
		}
	}

	if c.PVActuels <= 0 {
		GererMort(c, lecteur)
	}
}

// TrainingFight : combat d'entrainement contre le gobelin.
func TrainingFight(c *character.Character) {
	Combattre(c, InitGobelin())
}

// AffronterRoiGobelin : combat contre le Roi Gobelin.
func AffronterRoiGobelin(c *character.Character) {
	Combattre(c, InitRoiGobelin())
}

// AffronterDemon : combat contre le Démon.
func AffronterDemon(c *character.Character) {
	Combattre(c, InitDemon())
}

// AffronterDragon : combat contre le Dragon.
func AffronterDragon(c *character.Character) {
	Combattre(c, InitDragon())
}

// ----------------------------------------------------------------------------
// LANCEMENT PAR NOM DE MONSTRE
// ----------------------------------------------------------------------------

// Lancer déclenche un combat contre le monstre désigné par son nom (insensible
// à la casse). Exemples : combat.Lancer("gobelin", c), combat.Lancer("dragon", c)
func Lancer(nomMonstre string, c *character.Character) error {
	switch strings.ToLower(strings.TrimSpace(nomMonstre)) {
	case "gobelin":
		audio.PlayMusic("BanqueSon/musique/anciencombat.ogg")
		Combattre(c, InitGobelin())
	case "roi gobelin", "roigobelin", "roi_gobelin":
		audio.PlayMusic("BanqueSon/musique/anciencombat.ogg")
		Combattre(c, InitRoiGobelin())
	case "demon", "démon":
		audio.PlayMusic("BanqueSon/musique/incendie.ogg")
		Combattre(c, InitDemon())
	case "dragon":
		audio.PlayMusic("BanqueSon/musique/cataclysme.ogg")
		Combattre(c, InitDragon())
	default:
		return fmt.Errorf("monstre inconnu : '%s'", nomMonstre)
	}
	audio.PlayMusic("BanqueSon/musique/ambiance.ogg")
	return nil
}
