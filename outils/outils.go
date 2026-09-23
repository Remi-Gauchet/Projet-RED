package outils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func AttendreEntree() {
	fmt.Println("\nAppuyez sur Entrée pour continuer...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
