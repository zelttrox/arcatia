package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 300, Y: 300},
		Health:    100,
		Money:     1000,
		Damage:    5,
		Speed:     4,
		Inventory: []item.Item{},

		IsAlive:   true,
		IsRunning: false,

		IsAnimated:  true,
		FrameWidth:  100,
		FrameHeight: 100,
		MaxFrames:   8,

		SpriteIdle: rl.LoadTexture("textures/entities/soldier/Soldier-Idle.png"),
		SpriteRun:  rl.LoadTexture("textures/entities/soldier/Soldier-Walk.png"),
	}
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "claude",
		Position:  rl.Vector2{X: 400, Y: 320},
		Health:    20,
		MaxHealth: 20,
		Damage:    5,
		Loot:      []item.Item{},
		Worth:     12,
		Speed:     2,
		Origine:   rl.Vector2{X: 400, Y: 320},

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),

		IsAnimated:  true,
		FrameWidth:  100,
		FrameHeight: 100,
		MaxFrames:   6,
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "claude",
		Position:  rl.Vector2{X: 500, Y: 600},
		Health:    20,
		MaxHealth: 20,
		Damage:    5,
		Loot:      []item.Item{},
		Worth:     12,
		Speed:     2,
		Origine:   rl.Vector2{X: 500, Y: 600},

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),

		IsAnimated:  true,
		FrameWidth:  100,
		FrameHeight: 100,
		MaxFrames:   6,
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "claude",
		Position:  rl.Vector2{X: 600, Y: 600},
		Health:    200,
		MaxHealth: 200,
		Damage:    35,
		Loot:      []item.Item{},
		Worth:     200,
		Speed:     1,
		Origine:   rl.Vector2{X: 600, Y: 600},

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),

		IsAnimated:  true,
		FrameWidth:  100,
		FrameHeight: 100,
		MaxFrames:   6,
	})

	e.Player.Money = 300
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
