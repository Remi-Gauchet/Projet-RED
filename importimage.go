package main

import (
	"fmt"
	"os"
)

// Fonction pour afficher un art ASCII depuis un fichier
// On prend en paramètre un "chemin" (string) qui représente l'emplacement du fichier
func afficherASCII(chemin string) {
	// os.ReadFile lit tout le contenu du fichier d'un coup
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		fmt.Printf("[Erreur] Impossible de charger l'image : %v\n", err)
		return
	}

	// On convertit les octets ([]byte) en string et on affiche
	fmt.Println(string(contenu))
}

func main() {
	fmt.Println("Un monstre sauvage apparaît !")

	// On appelle la fonction en lui passant le chemin sous forme de texte (string)
	// Assure-toi que le dossier "BanqueASCII" et le fichier "gobelin.txt" existent bien au même endroit !
	afficherASCII("BanqueASCII/gobelin.txt")
}
