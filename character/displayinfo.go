package character

import (
	"fmt"

	"scarlet/ascii"
)

func (c Character) DisplayInfo() {
	chemin := "BanqueASCII/" + c.Classe + ".txt"
	ascii.AfficherASCII(chemin)
	fmt.Println(c)

	if c.QueteMagieNoire == 0 {
		fmt.Println("\n--- Progression ---")
		fmt.Println("Je commence à prendre mes marques dans le duché de Scarlet. Je suis épuisé...")
		fmt.Println("Je me demande où est l'auberge. Et j'aimerais autant en profiter pour faire le plein de provisions.")
	}

	if c.QueteMagieNoire == 1 {
		fmt.Println("\n--- Progression ---")
		fmt.Println("Le maire a besoin de mon aide pour éliminer les gobelins qui harcèlent les fermes.")
		fmt.Println("La carte qu'il m'a donné devrait m'aider à me repérer en sortant de la ville.")
	}

	if c.QueteMagieNoire == 2 {
		fmt.Println("\n--- Progression ---")
		fmt.Println("J'ai récupéré un étrange cristal d'ombre sur la dépouille du roi des gobelins.")
		fmt.Println("Un expert de l'ombre et de la lumière pourra certainement m'en dire plus. ")
	}

	if c.QueteMagieNoire == 3 {
		fmt.Println("\n--- Progression ---")
		fmt.Println("Le cristal sert maintenant de boussole vers le lieu d'origine du rituel.")
		fmt.Println("Qui sait ce que j'y trouverai ? Je ferais mieux de me préparer.")
	}

	if c.QueteMagieNoire == 4 {
		fmt.Println("\n--- Progression ---")
		fmt.Println("La ville est en danger ! Ils ont besoin de mon aide.")
		fmt.Println("La puissante mage noire que le prêtre redoutait est donc derrière tout ça.")
		fmt.Println("Son armée est en route, je dois m'attendre au pire...")
	}
}
