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

	IsAlive   bool
	IsRunning bool
	Dir       int

	IsAnimated  bool
	FrameWidth  int
	FrameHeight int
	MaxFrames   int

	SpriteIdle rl.Texture2D
	SpriteRun  []rl.Texture2D
}

func (p *Player) Attack(m *Monster) {
	m.Health -= p.Damage
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}

var CurrentFrame2 int
var frameCount2 int
var speed2 int = 5

func (p Player) PlayerAnimation() {
	if frameCount2 >= speed2 {
		CurrentFrame2++
		frameCount2 = 0
	} else {
		frameCount2++
	}
	if CurrentFrame2 >= p.MaxFrames {
		CurrentFrame2 = 0
	}
}

func (p Player) PlayerIdle() {
	CurrentFrame = 0
}

func (p *Player) PlayerDraw() {
	frameRec := rl.Rectangle{
		X: float32(p.FrameWidth + p.FrameWidth*CurrentFrame), Y: 0,
		Width: float32(p.FrameWidth), Height: float32(p.FrameHeight),
	}
	position := rl.Rectangle{
		X: p.Position.X + 15, Y: p.Position.Y,
		Width: float32(p.FrameWidth), Height: float32(p.FrameHeight),
	}
	rl.DrawTexturePro(p.SpriteRun[p.Dir], frameRec, position, rl.Vector2{}, 0, rl.White)
}
