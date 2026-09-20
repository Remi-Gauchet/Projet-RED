package character

import (
	"fmt"
	"strings"
)

func (c Character) AccessInventory() {
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	detail := strings.Join(c.Inventaire, " - ")
	fmt.Println("Voici le contenu de votre sac :")
	fmt.Println(detail)
}
