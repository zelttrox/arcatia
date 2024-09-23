package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME   engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
	WIN      engine = iota
)

var SALLE int = 1

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	NPC      entity.NPC

	Music       rl.Music
	DeathMusic  rl.Music
	MusicVolume float32
	ShowFPS     bool
	Godmode     bool
	SpeedBoost  bool

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
