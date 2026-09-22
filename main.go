package main

import (
	"bufio"
	"fmt"
	"os"

	"scarlet/audio"
	"scarlet/character"
	"scarlet/exploration"
)

func afficherIntro(classe string) {
	chemin := "introduction/" + classe + ".txt"
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		fmt.Printf("[Erreur] Impossible de charger l'introduction : %v\n", err)
		return
	}
	fmt.Println(string(contenu))
}

func main() {
	c1, err := character.InitCharacter(
		1,
		[]string{"Potion de soins", "Potion de soins", "Potion de soins"},
	)
	if err != nil {
		fmt.Println("Erreur lors de la création du personnage :", err)
		return
	}

	afficherIntro(c1.Classe)

	attendreEntree()

	audio.PlayMusic("BanqueSon/musique/ambiance.ogg")

	exploration.Explorer(&c1)
}

func attendreEntree() {
	fmt.Println("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
