package ascii

import (
	"fmt"
	"os"
)

func AfficherASCII(chemin string) { // majuscule pour l'exporter
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		fmt.Printf("[Erreur] Impossible de charger l'image : %v\n", err)
		return
	}
	fmt.Println(string(contenu))
}
