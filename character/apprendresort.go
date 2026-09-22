package character

import (
	"fmt"
	"strings"
)

const PrefixeLivreSort = "Livre de Sort : "

// ApprendreSort consomme un livre de sort de l'inventaire et ajoute le sort
// correspondant à la liste des sorts connus du personnage.
func (c *Character) ApprendreSort(nomLivre string) error {
	nomSort := strings.TrimPrefix(nomLivre, PrefixeLivreSort)

	for _, s := range c.Sorts {
		if s == nomSort {
			return fmt.Errorf("vous connaissez déjà le sort '%s'", nomSort)
		}
	}

	err := c.RemoveInventory(nomLivre)
	if err != nil {
		return fmt.Errorf("aucun exemplaire de '%s' dans l'inventaire", nomLivre)
	}

	c.Sorts = append(c.Sorts, nomSort)
	fmt.Printf("Vous avez appris le sort : %s !\n", nomSort)
	return nil
}
