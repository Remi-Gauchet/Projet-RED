package character

import (
	"fmt"

	"scarlet/ascii"
)

func (c Character) DisplayInfo() {
	chemin := "BanqueASCII/" + c.Classe + ".txt"
	ascii.AfficherASCII(chemin)
	fmt.Println(c)

	if c.QueteMagieNoire == 1 {
		fmt.Println("\n--- Quête en cours ---")
		fmt.Println("Découvrir l'origine de la magie noire qui s'est répandue dans la région.")
	}
}
