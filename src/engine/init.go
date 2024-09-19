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
	rl.InitWindow(ScreenWidth, ScreenHeight, "Meow, Meow, Meow Meow")

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
		Position:  rl.Vector2{X: 128, Y: 180},
		Health:    100,
		Money:     1000,
		Damage:    5,
		Speed:     2,
		Inventory: []item.Item{},

		IsAlive:   true,
		IsRunning: false,

		HealCount: 10,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		SpriteIdle:       rl.LoadTexture("textures/entities/cat/2_Cat_Run-Sheet.png"),
		SpriteRun:        []rl.Texture2D{rl.LoadTexture("textures/entities/cat/2_Cat_Run-Sheet_Right.png"), rl.LoadTexture("textures/entities/cat/2_Cat_Run-Sheet_Left.png")},
		HomescreenSprite: rl.LoadTexture("textures/menu/homescreen.png"),
		GameoverSprite:   rl.LoadTexture("textures/menu/gameover.png"),
		GoodGameSprite:   rl.LoadTexture("textures/menu/goodgame.jpg"),
	}
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "distributeur",
		Position:  rl.Vector2{X: 240, Y: 460},
		Health:    200,
		MaxHealth: 200,
		Damage:    0,
		Loot:      []item.Item{},
		Worth:     0,
		Speed:     0,
		Origine:   rl.Vector2{X: 240, Y: 460},

		IsAlive: true,

		IsAnimated:  false,
		FrameWidth:  100,
		FrameHeight: 100,
		MaxFrames:   10,
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "chien",
		Position:  rl.Vector2{X: 212, Y: 20},
		Health:    50,
		MaxHealth: 50,
		Damage:    10,
		Loot:      []item.Item{},
		Worth:     200,
		Speed:     2,
		Origine:   rl.Vector2{X: 212, Y: 20},

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/dog/GoldenIdle.png"),

		IsAnimated:  true,
		FrameWidth:  64,
		FrameHeight: 64,
		MaxFrames:   6,
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Mouse",
		Position:  rl.Vector2{X: 300, Y: 500},
		Health:    25,
		MaxHealth: 25,
		Damage:    2,
		Loot:      []item.Item{},
		Worth:     40,
		Speed:     2,
		Origine:   rl.Vector2{X: 300, Y: 500},

		IsAlive: true,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		Sprite: rl.LoadTexture("textures/entities/mice/rat_1.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Mouse",
		Position:  rl.Vector2{X: 310, Y: 510},
		Health:    25,
		MaxHealth: 25,
		Damage:    2,
		Loot:      []item.Item{},
		Worth:     40,
		Speed:     2,
		Origine:   rl.Vector2{X: 310, Y: 510},

		IsAlive: true,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		Sprite: rl.LoadTexture("textures/entities/mice/rat_1.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Mouse",
		Position:  rl.Vector2{X: 290, Y: 490},
		Health:    25,
		MaxHealth: 25,
		Damage:    2,
		Loot:      []item.Item{},
		Worth:     40,
		Speed:     2,
		Origine:   rl.Vector2{X: 290, Y: 490},

		IsAlive: true,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		Sprite: rl.LoadTexture("textures/entities/mice/rat_1.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Mouse",
		Position:  rl.Vector2{X: 320, Y: 520},
		Health:    25,
		MaxHealth: 25,
		Damage:    2,
		Loot:      []item.Item{},
		Worth:     40,
		Speed:     2,
		Origine:   rl.Vector2{X: 320, Y: 520},

		IsAlive: true,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		Sprite: rl.LoadTexture("textures/entities/mice/rat_1.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Mouse",
		Position:  rl.Vector2{X: 330, Y: 530},
		Health:    25,
		MaxHealth: 25,
		Damage:    2,
		Loot:      []item.Item{},
		Worth:     40,
		Speed:     2,
		Origine:   rl.Vector2{X: 323, Y: 523},

		IsAlive: true,

		IsAnimated:  true,
		FrameWidth:  32,
		FrameHeight: 32,
		MaxFrames:   10,

		Sprite: rl.LoadTexture("textures/entities/mice/rat_1.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:      "Boss",
		Position:  rl.Vector2{X: 360, Y: 1030},
		Health:    200,
		MaxHealth: 200,
		Damage:    25,
		Loot:      []item.Item{},
		Worth:     12,
		Speed:     2,
		Origine:   rl.Vector2{X: 360, Y: 1030},

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/dog/RottweilerIdle.png"),

		IsAnimated:  true,
		FrameWidth:  64,
		FrameHeight: 64,
		MaxFrames:   6,
	})

	// e.Monsters = append(e.Monsters, entity.Monster{
	// 	Name:      "chien",
	// 	Position:  rl.Vector2{X: 212, Y: 20},
	// 	Health:    100,
	// 	MaxHealth: 100,
	// 	Damage:    15,
	// 	Loot:      []item.Item{},
	// 	Worth:     200,
	// 	Speed:     1,
	// 	Origine:   rl.Vector2{X: 212, Y: 20},

	// 	IsAlive: true,
	// 	Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),

	// 	IsAnimated:  true,
	// 	FrameWidth:  100,
	// 	FrameHeight: 100,
	// 	MaxFrames:   10,
	// })

	e.Player.Money = 0
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		4.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/What_Was_I_Made_For.mp3")

	rl.PlayMusicStream(e.Music)
}
