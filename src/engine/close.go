package engine

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Close() {
	rl.CloseAudioDevice()
	rl.CloseWindow()
	os.Exit(0)

}
