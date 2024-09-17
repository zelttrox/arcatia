package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name      string
	Position  rl.Vector2
	Health    int
	MaxHealth int
	Damage    int
	Loot      []item.Item
	Worth     int //valeur en argent quand tué
	Speed     int
	Origine   rl.Vector2

	IsAlive bool

	Sprite  rl.Texture2D
	Sprites []rl.Texture2D

	IsAnimated  bool
	FrameWidth  int
	FrameHeight int
	MaxFrames   int
}

func (m *Monster) Attack(p *Player) {
	p.Health -= m.Damage
}

func (m *Monster) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}

var CurrentFrame int
var frameCount int
var speed int = 5

func (m Monster) UpdateAnimation() {
	if frameCount >= speed {
		CurrentFrame++
		frameCount = 0
	} else {
		frameCount++
	}
	if CurrentFrame >= m.MaxFrames {
		CurrentFrame = 0
	}
}

func (m *Monster) Draw() {
	frameRec := rl.Rectangle{
		X: float32(m.FrameWidth + m.FrameWidth*CurrentFrame), Y: 0,
		Width: float32(m.FrameWidth), Height: float32(m.FrameHeight),
	}
	position := rl.Rectangle{
		X: m.Position.X + 20, Y: m.Position.Y + 10,
		Width: float32(m.FrameWidth), Height: float32(m.FrameHeight),
	}
	rl.DrawTexturePro(m.Sprite, frameRec, position, rl.Vector2{}, 0, rl.White)
}
