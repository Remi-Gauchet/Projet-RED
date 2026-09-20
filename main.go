package main

import (
	"fmt"

	"scarlet/character"
)

func main() {
	c1, err := character.InitCharacter(
		1,
		100,
		40,
		[]string{"Potion de soins", "Potion de soins", "Potion de soins"},
	)
	if err != nil {
		fmt.Println("Erreur lors de la création du personnage :", err)
		return
	}

	c1.DisplayInfo()
	c1.AccessInventory()
}
