package main

import (
	"fmt"
	"os"

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
		100,
		40,
		[]string{"Potion de soins", "Potion de soins", "Potion de soins"},
		100,
	)

	if err != nil {
		fmt.Println("Erreur lors de la création du personnage :", err)
		return
	}

	afficherIntro(c1.Classe)

	exploration.Explorer(&c1)
}
