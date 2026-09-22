package character

import "fmt"

// UpgradeInventorySlot augmente la taille maximale de l'inventaire, contre de l'or.
// Limité à MaxUpgrades améliorations au total.
func (c *Character) UpgradeInventorySlot() error {
	if c.UpgradesRestantes <= 0 {
		return fmt.Errorf("vous avez déjà amélioré votre inventaire. Vous ne pourrez jamais transporter davantage.")
	}

	if c.Or < CoutUpgrade {
		return fmt.Errorf("vous n'avez pas assez d'or (%d nécessaires, vous en possédez %d)", CoutUpgrade, c.Or)
	}

	c.Or -= CoutUpgrade
	c.InventaireMax += BonusUpgrade
	c.UpgradesRestantes--

	fmt.Printf("Capacité de l'inventaire augmentée ! Nouvelle taille maximum : %d\n", c.InventaireMax)
	return nil
}
