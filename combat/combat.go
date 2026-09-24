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
	XPAccordee   int
}

// ----------------------------------------------------------------------------
// INITIALISATION DES MONSTRES
// ----------------------------------------------------------------------------

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
		XPAccordee: 0, // Généré aléatoirement (35 à 50) pour les niveaux >= 2
	}
}

func InitRoiGobelin() Monstre {
	return Monstre{
		Nom:          "Roi Gobelin",
		FichierASCII: "roi_gobelin.txt",
		PVMax:        90,
		PVActuels:    90,
		Attaques: []Attaque{
			{Nom: "Coup de Sceptre", Chance: 40, Degats: 20},
			{Nom: "Charge Royale", Chance: 25, Degats: 30},
			{Nom: "Piétinement", Chance: 20, Degats: 15},
			{Nom: "Ordre de Massacre", Chance: 10, Degats: 35},
			{Nom: "Fureur du Roi", Chance: 5, Degats: 45},
		},
		XPAccordee: 80,
	}
}

func InitDemon() Monstre {
	return Monstre{
		Nom:          "Démon",
		FichierASCII: "demon.txt",
		PVMax:        150,
		PVActuels:    150,
		Attaques: []Attaque{
			{Nom: "Griffe Infernale", Chance: 35, Degats: 25},
			{Nom: "Souffle de Soufre", Chance: 25, Degats: 35},
			{Nom: "Poing Ténébreux", Chance: 20, Degats: 30},
			{Nom: "Flammes Abyssales", Chance: 15, Degats: 45},
			{Nom: "Rituel de Douleur", Chance: 5, Degats: 60},
		},
		XPAccordee: 150,
	}
}

func InitDragon() Monstre {
	return Monstre{
		Nom:          "Dragon",
		FichierASCII: "dragon.txt",
		PVMax:        250,
		PVActuels:    250,
		Attaques: []Attaque{
			{Nom: "Morsure Écailleuse", Chance: 35, Degats: 30},
			{Nom: "Coup de Queue", Chance: 25, Degats: 40},
			{Nom: "Griffure Ardente", Chance: 20, Degats: 45},
			{Nom: "Souffle de Feu", Chance: 15, Degats: 60},
			{Nom: "Rugissement Dévastateur", Chance: 5, Degats: 80},
		},
		XPAccordee: 300,
	}
}

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
// AFFICHAGE ET ENCADREMENT
// ----------------------------------------------------------------------------

const largeurColonne = 80
const largeurBoite = 33

func chargerLignes(chemin string) []string {
	data, err := os.ReadFile(chemin)
	if err != nil {
		return []string{"[ASCII introuvable]"}
	}
	return strings.Split(string(data), "\n")
}

func afficherASCIIBrut(chemin string) {
	for _, ligne := range chargerLignes(chemin) {
		fmt.Println(ligne)
	}
}

func encadrerLignes(lignes []string, largeur int) []string {
	var res []string
	res = append(res, "┌"+strings.Repeat("─", largeur)+"┐")

	for _, l := range lignes {
		rc := utf8.RuneCountInString(l)
		if rc > largeur {
			runes := []rune(l)
			l = string(runes[:largeur])
			rc = largeur
		}
		res = append(res, fmt.Sprintf("│%-*s│", largeur, l))
	}

	res = append(res, "└"+strings.Repeat("─", largeur)+"┘")
	return res
}

func AfficherArene(c *character.Character, m *Monstre) {
	// 1. Partie visuelle du Joueur (gauche)
	imgJoueur := chargerLignes("BanqueASCII/" + c.Classe + ".txt")
	actionsJoueur := []string{
		"1. Coup Risqué (50%, 20 deg)",
		"2. Frappe Rapide (10 deg)",
		"3. Potion de soins",
	}

	if c.Niveau >= 2 {
		actionsJoueur = append(actionsJoueur, "4. Coup Assommant (20 deg, 40% stun)")
		actionsJoueur = append(actionsJoueur, "5. Lancer un sort")
	} else {
		actionsJoueur = append(actionsJoueur, "4. Lancer un sort")
	}

	var gauche []string
	gauche = append(gauche, imgJoueur...)
	gauche = append(gauche, fmt.Sprintf(" %s (Niveau %d)", c.Nom, c.Niveau))
	gauche = append(gauche, fmt.Sprintf(" PV : %d/%d | Mana : %d/%d", c.PVActuels, c.PVMax, c.ManaActuel, c.ManaMax))

	if c.Niveau >= 2 {
		gauche = append(gauche, fmt.Sprintf(" EXP : %d/%d", c.XP, c.XPMax))
	}

	gauche = append(gauche, "")
	gauche = append(gauche, " Actions :")
	gauche = append(gauche, encadrerLignes(actionsJoueur, largeurBoite)...)

	// 2. Partie visuelle du Monstre (droite)
	imgMonstre := chargerLignes("BanqueASCII/" + m.FichierASCII)
	var attaquesMonstre []string
	for _, a := range m.Attaques {
		attaquesMonstre = append(attaquesMonstre, fmt.Sprintf("- %s (%d%%, %ddg)", a.Nom, a.Chance, a.Degats))
	}

	var droite []string
	droite = append(droite, imgMonstre...)
	droite = append(droite, fmt.Sprintf(" %s", m.Nom))
	droite = append(droite, fmt.Sprintf(" PV : %d/%d", m.PVActuels, m.PVMax))
	droite = append(droite, "")
	droite = append(droite, " Attaques possibles :")
	droite = append(droite, encadrerLignes(attaquesMonstre, largeurBoite)...)

	// 3. Fusion ligne par ligne des deux colonnes
	maxLignes := len(gauche)
	if len(droite) > maxLignes {
		maxLignes = len(droite)
	}

	for i := 0; i < maxLignes; i++ {
		strGauche := ""
		if i < len(gauche) {
			strGauche = gauche[i]
		}

		strDroite := ""
		if i < len(droite) {
			strDroite = droite[i]
		}

		padG := largeurColonne - utf8.RuneCountInString(strGauche)
		if padG < 0 {
			padG = 0
		}
		fmt.Printf("%s%s   %s\n", strGauche, strings.Repeat(" ", padG), strDroite)
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
		fmt.Printf("\n%s lance Glace ! %s est gelé pour 2 tours.\n", c.Nom, monstre.Nom)

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
		fmt.Print("Choisissez votre action (numéro) : ")

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
				fmt.Printf("Coup réussi ! Vous infligez %d dégâts à %s !\n", degats, monstre.Nom)
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
			if c.Niveau >= 2 {
				degats := 20 * c.Niveau
				monstre.PVActuels -= degats
				if monstre.PVActuels < 0 {
					monstre.PVActuels = 0
				}
				fmt.Printf("\n%s utilise Coup Assommant et inflige %d dégâts !\n", c.Nom, degats)

				if rand.Intn(100) < 40 {
					monstre.ToursGeles = 1
					fmt.Printf("🎯 %s est assommé pour 1 tour !\n", monstre.Nom)
				}
				time.Sleep(1500 * time.Millisecond)
				return
			}
			gererMenuSorts(c, monstre, lecteur)
			return

		case "5":
			if c.Niveau >= 2 {
				gererMenuSorts(c, monstre, lecteur)
				return
			}
			fmt.Println("Choix invalide.")
			time.Sleep(1 * time.Second)

		default:
			fmt.Println("Choix invalide.")
			time.Sleep(1 * time.Second)
		}
	}
}

func gererMenuSorts(c *character.Character, monstre *Monstre, lecteur *bufio.Reader) {
	if len(c.Sorts) == 0 {
		fmt.Println("Aucun sort connu.")
		time.Sleep(1 * time.Second)
		return
	}
	fmt.Println("\nSorts connus :")
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
		return
	}

	lancerSort(c, monstre, c.Sorts[num-1])
	time.Sleep(500 * time.Millisecond)
}

// ----------------------------------------------------------------------------
// GESTION DE L'EXPERIENCE
// ----------------------------------------------------------------------------

func AjouterXP(c *character.Character, xpGagnee int) {
	if c.Niveau == 1 {
		// Passage direct au Niveau 2 pour le premier Gobelin
		c.Niveau = 2
		c.XP = 0
		c.XPMax = 100 // Objectif pour le niveau 3
		c.PVMax *= 2
		c.PVActuels = c.PVMax
		c.ManaMax *= 2
		c.ManaActuel = c.ManaMax
		fmt.Println("\n🌟 VOUS PASSEZ NIVEAU 2 !")
		fmt.Println("-> Vos dégâts et soins sont désormais doublés !")
		fmt.Println("-> Nouvelle attaque débloquée : Coup Assommant !")
		return
	}

	c.XP += xpGagnee
	fmt.Printf("\n✨ Vous gagnez %d points d'expérience ! (%d/%d EXP)\n", xpGagnee, c.XP, c.XPMax)

	// Verification et montée de niveau
	for c.XP >= c.XPMax {
		c.XP -= c.XPMax
		c.Niveau++

		// Multiplication du seuil par 1.5
		c.XPMax = int(float64(c.XPMax) * 1.5)

		c.PVMax += 20
		c.PVActuels = c.PVMax
		c.ManaMax += 10
		c.ManaActuel = c.ManaMax

		fmt.Printf("\n🌟 FELICITATIONS ! Vous passez Niveau %d !\n", c.Niveau)
		fmt.Printf("-> PV Max augmentés à %d\n", c.PVMax)
		fmt.Printf("-> Mana Max augmenté à %d\n", c.ManaMax)
		fmt.Printf("-> Prochain niveau à %d EXP requis !\n", c.XPMax)
	}
}

// ----------------------------------------------------------------------------
// ÉCRAN DE MORT
// ----------------------------------------------------------------------------

const coutResurrection = 10

func GererMort(c *character.Character, lecteur *bufio.Reader) {
	outils.ClearScreen()
	afficherASCIIBrut("BanqueASCII/mort.txt")
	chemin := "BanqueSon/" + c.Classe + "/mort.ogg"
	audio.PlaySound(chemin)

	fmt.Println("\nMême les héros finissent par céder à ce monde impitoyable...")

	for {
		fmt.Println("\n1. Ressusciter à la cathédrale (Or -10, PV au maximum)")
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
			fmt.Printf("\nUne lumière aveuglante vous illumine. Vous réapparaissez dans la cathédrale. (Or restant : %d)\n", c.Or)
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
// BOUCLE PRINCIPALE DE COMBAT
// ----------------------------------------------------------------------------

func Combattre(c *character.Character, monstre Monstre) bool {
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

			var xpGagnee int
			if strings.Contains(strings.ToLower(monstre.Nom), "gobelin") && !strings.Contains(strings.ToLower(monstre.Nom), "roi") {
				xpGagnee = rand.Intn(16) + 35 // entre 35 et 50 EXP
			} else {
				xpGagnee = monstre.XPAccordee
			}

			AjouterXP(c, xpGagnee)

			time.Sleep(3 * time.Second)
			break
		}

		TourMonstre(&monstre, c)
		if c.PVActuels <= 0 {
			outils.ClearScreen()
			AfficherArene(c, &monstre)
			fmt.Printf("\n☠️ Vous avez été vaincu par %s...\n", monstre.Nom)
			time.Sleep(2 * time.Second)
			break
		}
	}

	if c.PVActuels <= 0 {
		GererMort(c, lecteur)
		return true
	}

	return false
}

func TrainingFight(c *character.Character) {
	Combattre(c, InitGobelin())
}

func AffronterRoiGobelin(c *character.Character) {
	Combattre(c, InitRoiGobelin())
}

func AffronterDemon(c *character.Character) {
	Combattre(c, InitDemon())
}

func AffronterDragon(c *character.Character) {
	Combattre(c, InitDragon())
}

// ----------------------------------------------------------------------------
// LANCEMENT PAR NOM DE MONSTRE
// ----------------------------------------------------------------------------

func Lancer(nomMonstre string, c *character.Character) (bool, error) {
	var mort bool

	switch strings.ToLower(strings.TrimSpace(nomMonstre)) {
	case "gobelin":
		audio.PlayMusic("BanqueSon/musique/anciencombat.ogg")
		mort = Combattre(c, InitGobelin())
	case "roi gobelin", "roigobelin", "roi_gobelin":
		audio.PlayMusic("BanqueSon/musique/anciencombat.ogg")
		mort = Combattre(c, InitRoiGobelin())
	case "demon", "démon":
		audio.PlayMusic("BanqueSon/musique/incendie.ogg")
		mort = Combattre(c, InitDemon())
	case "dragon":
		audio.PlayMusic("BanqueSon/musique/cataclysme.ogg")
		mort = Combattre(c, InitDragon())
	default:
		return false, fmt.Errorf("monstre inconnu : '%s'", nomMonstre)
	}

	return mort, nil
}
