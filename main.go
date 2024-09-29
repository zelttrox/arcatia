package main

import (
	"flag"
	"fmt"
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// Gestion des arguments
	argFPS := flag.Bool("fps", false, "Fullscreen mode")
	argFullscreen := flag.Bool("fs", false, "Show FPS")
	argHelp := flag.Bool("h", false, "Help")
	argGodmode := flag.Bool("gm", false, "Godmode")
	argSpeeedBoost := flag.Bool("spd", false, "Speed boost")
	argCipher := flag.Bool("cipher", false, "Cipher")
	argBin := flag.Bool("bin", false, "Binary")
	flag.Parse()

	// Arg Help (-h)
	if *argHelp {
		fmt.Println("Args:")
		fmt.Println("-h for Help")
		fmt.Println("-fs for Fullscreen")
		fmt.Println("-fps to Show FPS")
		fmt.Println("-gm for Godmode")
		fmt.Println("-spd for Speed boost")
		fmt.Println("-cipher for Cipher Dialogues")
		fmt.Println("-bin for Binary Dialogues")
	} else {

		// Arg Fullscreen (-fs)
		if *argFullscreen {
			rl.InitWindow(1920, 1080, "Fullscreen")
		} else {
			rl.InitWindow(1400, 800, "Meow, Meow, Meow Meow")
		}

		var e engine.Engine
		e.Init()
		e.Load()
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		if *argFPS {
			e.ShowFPS = true
		}
		if *argGodmode {
			e.Godmode = true
		}
		if *argSpeeedBoost {
			e.SpeedBoost = true
		}
		if *argCipher {
			e.CipherDialogue = true
		}
		if *argBin {
			e.BinDialogue = true
		}

		e.Run()
		e.Unload()
		e.Close()
	}
}
