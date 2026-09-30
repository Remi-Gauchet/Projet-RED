package outils

import (
	"bufio"
	"fmt"
	"os"
)

func ClearScreen() {
	fmt.Print("\x1b[H\x1b[2J")
}

func AttendreEntree() {
	fmt.Println("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
