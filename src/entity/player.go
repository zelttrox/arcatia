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

	Sprite rl.Texture2D
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
