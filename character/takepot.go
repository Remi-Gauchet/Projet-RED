package character

import "fmt"

func (c *Character) TakePot() error {
	if c.PVActuels == c.PVMax {
		return fmt.Errorf("Je suis déjà en pleine forme. Gardons cette potion pour plus tard !")
	}

	index := -1
	for i, objet := range c.Inventaire {
		if objet == "Potion de soins" {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Je n'ai plus de potion de soins...")
	}

	// Retire la potion trouvée à l'index "index" du slice
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)

	c.PVActuels += 50
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}

	fmt.Printf("PV%d/%d\n", c.PVActuels, c.PVMax)
	return nil
}
