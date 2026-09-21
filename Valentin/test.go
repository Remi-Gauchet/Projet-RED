package main
 
import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)
 
// Caractères du plus foncé au plus clair
const chars = "@%#*+=-:. "
 
func imageToASCII(path string, width int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
 
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
 
	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()
 
	// Ajuster la hauteur pour compenser la forme des caractères
	aspectRatio := float64(origHeight) / float64(origWidth)
	height := int(float64(width) * aspectRatio * 0.55)
 
	result := ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Position correspondante dans l'image originale
			srcX := x * origWidth / width
			srcY := y * origHeight / height
 
			r, g, b, _ := img.At(bounds.Min.X+srcX, bounds.Min.Y+srcY).RGBA()
			// Luminosité moyenne (valeurs sur 16 bits -> ramenées sur 0-255)
			gray := (r + g + b) / 3 / 257
 
			index := int(gray) * (len(chars) - 1) / 255
			result += string(chars[index])
		}
		result += "\n"
	}
 
	return result, nil
}
 
func main() {
	result, err := imageToASCII("dragon.png", 120)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
 
	fmt.Println(result)
 
	err = os.WriteFile("dragon_ascii.txt", []byte(result), 0644)
	if err != nil {
		fmt.Println("Erreur écriture:", err)
	}
}
 