package main

import (
	"fmt"
	"os"

	"scarlet/audio"
	"scarlet/character"
	"scarlet/exploration"
	"scarlet/outils"
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

	outils.ClearScreen()
	afficherIntro(c1.Classe)

	outils.AttendreEntree()

	audio.PlayMusic("BanqueSon/musique/ambiance.ogg")

	exploration.Explorer(&c1)
}
