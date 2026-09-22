package character

import "fmt"

const CoutAuberge = 5

// SeReposer fait payer le prix d'une nuit au personnage et restaure entièrement
// ses points de vie et son mana.
func (c *Character) SeReposer() error {
	if c.Or < CoutAuberge {
		return fmt.Errorf("vous n'avez pas assez d'or pour payer la chambre (%d nécessaires, vous en possédez %d)", CoutAuberge, c.Or)
	}

	c.Or -= CoutAuberge
	c.PVActuels = c.PVMax
	c.ManaActuel = c.ManaMax

	fmt.Println("Malgré les bruits des poivrots et d'un couple libidineux dans la chambre d'à côté, vous parvenez à trouver le sommeil.")
	fmt.Println("Vos PV et votre mana sont entièrement restaurés !")
	return nil
}
