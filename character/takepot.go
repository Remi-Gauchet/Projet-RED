package character

import "fmt"

func (c *Character) TakePot() error {
    if c.PVActuels == c.PVMax {
        return fmt.Errorf("points de vie déjà au maximum, impossible de boire une potion")
    }

    err := c.RemoveInventory("Potion de soins")
    if err != nil {
        return fmt.Errorf("Aucune potion de soins dans l'inventaire")
    }

    c.PVActuels += 50
    if c.PVActuels > c.PVMax {
        c.PVActuels = c.PVMax
    }

    fmt.Printf("PV%d/%d\n", c.PVActuels, c.PVMax)
    return nil
}
