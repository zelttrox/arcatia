package fight

import (
	"main/src/entity"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player *entity.Player, monster *entity.Monster) {

	for { // Boucle infinie
		// Check si le joueur ou le monstre est vaincu. Si c'est le cas, on sort de la boucle

		if player.Health <= 0 {
			player.IsAlive = false
			break
		} else if monster.Health <= 0 {
			player.Inventory = append(player.Inventory, monster.Loot...)
			player.Money += monster.Worth
			monster.IsAlive = false
			break
		}
		player.Attack(monster)
		monster.Attack(player)
	}
}
