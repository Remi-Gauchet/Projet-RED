package outils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Encadrer(lignes []string, largeur int) {
	// Marge fixée à 74 espaces
	marge := strings.Repeat(" ", 74)

	fmt.Println(marge + "┌" + strings.Repeat("─", largeur+2) + "┐")
	for _, l := range lignes {
		rc := utf8.RuneCountInString(l)
		if rc > largeur {
			runes := []rune(l)
			l = string(runes[:largeur])
			rc = largeur
		}
		fmt.Printf("%s│ %s%s │\n", marge, l, strings.Repeat(" ", largeur-rc))
	}
	fmt.Println(marge + "└" + strings.Repeat("─", largeur+2) + "┘")
}