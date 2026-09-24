package outils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Encadrer(lignes []string) {
	marge := strings.Repeat(" ", 30)

	// Calcule la largeur nécessaire = longueur de la ligne la plus longue
	largeur := 0
	for _, l := range lignes {
		rc := utf8.RuneCountInString(l)
		if rc > largeur {
			largeur = rc
		}
	}

	fmt.Println(marge + "┌" + strings.Repeat("─", largeur+2) + "┐")
	for _, l := range lignes {
		rc := utf8.RuneCountInString(l)
		fmt.Printf("%s│ %s%s │\n", marge, l, strings.Repeat(" ", largeur-rc))
	}
	fmt.Println(marge + "└" + strings.Repeat("─", largeur+2) + "┘")
}
