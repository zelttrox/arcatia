package item

import (
	"fmt"
)

type Item struct {
	Name         string
	Price        int
	IsConsumable bool
	IsEquippable bool
}

func (i *Item) ToString() {
	fmt.Printf("Je suis un item qui vaut %d €\n", i.Price)
}
