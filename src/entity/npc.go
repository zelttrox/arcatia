package entity

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	Position rl.Vector2
	Health   int
	Money    int
	Speed    float32
	Damage   int

	IsAlive bool
	Dir     int

	IsAnimated  bool
	FrameWidth  int
	FrameHeight int
	MaxFrames   int

	SpriteIdle rl.Texture2D
}

func (p *NPC) Attack(m *Monster) {
	m.Health -= p.Damage
}

func (p *NPC) ToString() {
	fmt.Printf(`
	Joueur:
		Vie : %d,
		Argent: %d,
	
	\n`, p.Health, p.Money)
}

func (p NPC) NPCIdle() {
	CurrentFrame = 0
}

func (p *NPC) NPCDraw() {
	frameRec := rl.Rectangle{
		X: float32(p.FrameWidth + p.FrameWidth*CurrentFrame), Y: 160,
		Width: float32(p.FrameWidth), Height: float32(p.FrameHeight),
	}
	position := rl.Rectangle{
		X: p.Position.X + 15, Y: p.Position.Y,
		Width: float32(p.FrameWidth), Height: float32(p.FrameHeight),
	}
	rl.DrawTexturePro(p.SpriteIdle, frameRec, position, rl.Vector2{}, 0, rl.White)
}
