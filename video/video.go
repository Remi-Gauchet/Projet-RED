package video

import (
	"embed"
	"fmt"
	"os"
	"sort"
	"time"
)

//go:embed frames_ascii/*.txt
var frameFS embed.FS

func PlayASCII(fps int) error {
	entries, err := frameFS.ReadDir("frames_ascii")
	if err != nil {
		return fmt.Errorf("lecture dossier frames_ascii: %w", err)
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}
	sort.Strings(names)

	fmt.Print("\x1b[2J")

	delay := time.Second / time.Duration(fps)

	for _, name := range names {
		data, err := frameFS.ReadFile("frames_ascii/" + name)
		if err != nil {
			continue
		}
		fmt.Print("\x1b[H")
		os.Stdout.Write(data)
		time.Sleep(delay)
	}

	return nil
}
