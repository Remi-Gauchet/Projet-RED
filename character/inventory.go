package character

import "fmt"

// AddInventory ajoute un objet à l'inventaire, en respectant la taille maximale.
func (c *Character) AddInventory(item string) error {
	if len(c.Inventaire) >= MaxInventaire {
		return fmt.Errorf("inventaire plein (maximum %d objets)", MaxInventaire)
	}
	c.Inventaire = append(c.Inventaire, item)
	return nil
}

// RemoveInventory retire la première occurrence d'un objet de l'inventaire.
func (c *Character) RemoveInventory(item string) error {
	index := -1
	for i, objet := range c.Inventaire {
		if objet == item {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("objet '%s' introuvable dans l'inventaire", item)
	}
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
	return nil
}
