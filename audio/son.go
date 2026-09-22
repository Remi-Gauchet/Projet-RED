package audio

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
)

const sampleRate = beep.SampleRate(44100)

var (
	speakerInitialized bool
	currentMusic       *beep.Ctrl // référence à la musique de fond actuellement en cours
)

// PlaySound joue un son sans attendre la fin (non-bloquant).
func PlaySound(path string) {
	playInternal(path, false)
}

// PlaySoundBlocking joue un son et attend qu'il soit terminé avant de continuer.
func PlaySoundBlocking(path string) {
	playInternal(path, true)
}

func playInternal(path string, wait bool) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur ouverture son:", err)
		return
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		fmt.Println("Erreur décodage son:", err)
		return
	}

	if !speakerInitialized {
		speaker.Init(sampleRate, sampleRate.N(time.Second/10))
		speakerInitialized = true
	}

	resampled := beep.Resample(4, format.SampleRate, sampleRate, streamer)

	done := make(chan bool, 1) // bufferisé : l'envoi ne bloque jamais, même sans lecteur
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		streamer.Close()
		done <- true
	})))

	if wait {
		<-done
	}
}

// PlayMusic arrête la musique de fond en cours (s'il y en a une)
// et lance la nouvelle musique, jouée une seule fois.
func PlayMusic(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur ouverture musique:", err)
		return
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		fmt.Println("Erreur décodage musique:", err)
		return
	}

	if !speakerInitialized {
		speaker.Init(sampleRate, sampleRate.N(time.Second/10))
		speakerInitialized = true
	}

	resampled := beep.Resample(4, format.SampleRate, sampleRate, streamer)

	speaker.Lock()
	if currentMusic != nil {
		currentMusic.Paused = true
	}
	speaker.Unlock()

	ctrl := &beep.Ctrl{Streamer: resampled, Paused: false}
	currentMusic = ctrl

	speaker.Play(ctrl)
}

// StopMusic arrête la musique de fond en cours, sans en lancer une nouvelle.
func StopMusic() {
	speaker.Lock()
	if currentMusic != nil {
		currentMusic.Paused = true
	}
	speaker.Unlock()
}
