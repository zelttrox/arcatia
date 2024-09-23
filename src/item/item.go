package item

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
	Name         string
	Price        int
	IsConsumable bool
	IsEquippable bool
	Sprite       rl.Texture2D
}

func (i *Item) ToString() {
	fmt.Printf("Je suis un item qui vaut %d €\n", i.Price)
}
