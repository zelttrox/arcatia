package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position  rl.Vector2
	Health    int
	Money     int
	Speed     float32
	Damage    int
	Inventory []item.Item

	IsAlive bool
	IsRunning bool

	Sprite      rl.Texture2D
	IsAnimated  bool
	FrameWidth  int
	FrameHeight int
	MaxFrames   int
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 5
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}

var playerCurrentFrame int
var playerFrameCount int
var playerAnimationSpeed int = 5

func (p *Player) UpdateAnimation() {
	if playerFrameCount >= playerAnimationSpeed {
		playerCurrentFrame++
		playerFrameCount = 0
	} else {
		playerFrameCount++
	}
	if playerCurrentFrame >= p.MaxFrames {
		playerCurrentFrame = 0
	}
}

func (p *Player) Draw() {
	frameRec := rl.Rectangle{
		X: float32(p.FrameWidth * playerCurrentFrame), Y: 0,
		Width:  float32(p.FrameWidth),
		Height: float32(p.FrameHeight),
	}
	position := rl.Rectangle{
		X: p.Position.X, Y: p.Position.Y,
		Width:  float32(p.FrameWidth),
		Height: float32(p.FrameHeight),
	}
	rl.DrawTexturePro(p.Sprite, frameRec, position, rl.Vector2{}, 0, rl.White)
}
