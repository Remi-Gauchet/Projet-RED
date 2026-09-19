package character

import (
	"fmt"

	"scarlet/ascii"
)

func (c Character) DisplayInfo() {
	chemin := "BanqueASCII/" + c.Classe + ".txt"
	ascii.AfficherASCII(chemin)
	fmt.Println(c)
}
