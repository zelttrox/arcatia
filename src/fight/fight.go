package fight

import (
	"fmt"
	"main/src/entity"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

var cmt int = 0

func Fightm(player *entity.Player, monster *entity.Monster) {
	cmt++
	if cmt%60 == 0 {
		monster.Attack(player)
		fmt.Println(player.Health)
		cmt = 0
	}
	if player.Health <= 0 {
		player.IsAlive = false
	}
}

func Fightp(player *entity.Player, monster *entity.Monster) {
	player.Attack(monster)
	fmt.Println(monster.Health)
	if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
		monster.IsAlive = false
	}

}
