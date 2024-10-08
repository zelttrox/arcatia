package engine

import (
	"fmt"
	"main/src/fight"
	"main/src/item"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Meow.mp3")
		rl.SetMusicVolume(e.Music, e.MusicVolume)
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) Reset() {
	SALLE = 1
	e.Player.Health = 100
	e.Player.Money = 300
	e.Player.Inventory = []item.Item{}
	e.Player.IsAlive = true
	e.Player.Position = rl.Vector2{X: 128, Y: 180}
	for i := range e.Monsters {
		monster := &e.Monsters[i]
		monster.IsAlive = true
		monster.Health = monster.MaxHealth
		monster.Position = monster.Origine
	}
	e.StateMenu = PLAY
	e.StateEngine = INGAME
}

func (e *Engine) GameOverLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.Reset()
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) GoodGameLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.Reset()
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {

	if e.Godmode {
		e.Player.Health = 100
	}
	if e.SpeedBoost {
		e.Player.Speed = 6
	}

	e.Player.IsRunning = false

	e.Collisions()

	// Mouvement
	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 20, Y: e.Player.Position.Y + 20}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Meow.mp3")
		rl.SetMusicVolume(e.Music, e.MusicVolume)
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {
	e.ChasePlayer()
	e.MonsterCollisions()
	e.NPCCollisions()
}

func (e *Engine) MonsterCollisions() {
	if e.Player.Money < 0 {
		e.Player.Money = 0
	}
	for i := range e.Monsters {
		monster := &e.Monsters[i]
		if monster.IsAlive {
			if monster.Position.X > e.Player.Position.X-40 &&
				monster.Position.X < e.Player.Position.X+40 &&
				monster.Position.Y > e.Player.Position.Y-40 &&
				monster.Position.Y < e.Player.Position.Y+40 {
				rl.PlaySound(e.Player.BossSound)
				e.PlaySound()
				if rl.IsKeyPressed(rl.KeyR) {
					if monster.Name == "distributeur" && e.Player.Money >= 100 {
						e.Player.Money -= 100 + (e.Player.Damage - 5)
						e.Player.Damage += 3
					}
				}
				if rl.IsKeyPressed(rl.KeyE) {
					if monster.Name == "distributeur" && e.Player.Money >= 50 && e.Player.Health < 100 {
						if e.Player.Health >= 85 {
							e.Player.Health = 100
							e.Player.Money -= 50 + e.Player.HealCount
							e.Player.HealCount += 10
						} else {
							e.Player.Health += 15
							e.Player.Money -= 50 + e.Player.HealCount
							e.Player.HealCount += 10
						}

					} else {
						fight.Fightp(&e.Player, monster)

					}
				}
			}
			if monster.Position.X > e.Player.Position.X-20 &&
				monster.Position.X < e.Player.Position.X+20 &&
				monster.Position.Y > e.Player.Position.Y-20 &&
				monster.Position.Y < e.Player.Position.Y+20 {
				fight.Fightm(&e.Player, monster)
				if !e.Player.IsAlive {
					e.StateEngine = GAMEOVER
				}
				if monster.Name == "Boss" && !monster.IsAlive {
					e.StateEngine = WIN
				}
			}

		}
	}
}

func (e *Engine) PlaySound() {
	for i := range e.Monsters {
		monster := &e.Monsters[i]
		if monster.Name == "Mouse1" && !monster.SoundPlayed {
			rl.PlaySound(e.Player.MouseSound)
			monster.SoundPlayed = true
		}
		if monster.Name == "Boss" && !monster.SoundPlayed {
			rl.SetSoundVolume(e.Player.BossSound, 1.0)
			rl.PlaySound(e.Player.BossSound)
			monster.SoundPlayed = true
			fmt.Println("Playing sound")
		}
	}
}

func (e *Engine) NPCCollisions() {
	if e.NPC.IsAlive {
		if e.NPC.Position.X > e.Player.Position.X-40 &&
			e.NPC.Position.X < e.Player.Position.X+40 &&
			e.NPC.Position.Y > e.Player.Position.Y-40 &&
			e.NPC.Position.Y < e.Player.Position.Y+40 {
			if e.Player.NeverMet {
				e.RenderDialog("Tom, my sweet litle kitty!")
				e.Player.NeverMet = false
			}
		} else if !e.Player.NeverMet && !e.Player.KilledMice {
			e.RenderDialog("I need you to get rid of those \n mice in the kitchen.")
		} else if e.Player.KilledMice {
			e.RenderDialog("If only you could understand me.. I would tell you to get rid of the neighbor's dog..")
		} else if e.Player.KilledBoss {
			e.RenderDialog("You finally made this dog shut it! I love you so much!")
		} else {
			e.RenderDialog("I don't know why I am talking..")
		}
	}
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.Close()
		rl.StopMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.Reset()
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) ChasePlayer() {

	for i := range e.Monsters {
		monster := &e.Monsters[i]
		if !(monster.Name == "distributeur") {
			distrel := math.Sqrt(math.Pow(float64(e.Player.Position.X)-float64(monster.Position.X), 2) + math.Pow(float64(e.Player.Position.Y)-float64(monster.Position.Y), 2))
			if distrel <= 60 && distrel > 20 {
				rl.PlaySound(e.Player.BossSound)
				xrel := float64(e.Player.Position.X-monster.Position.X) / distrel
				yrel := float64(e.Player.Position.Y-monster.Position.Y) / distrel
				if float32(float64(xrel)*float64(monster.Speed)) < 0 {
				}
				if float32(float64(xrel)*float64(monster.Speed)) >= 0 {
				}
				monster.Position.X += float32(float64(xrel) * float64(monster.Speed))
				monster.Position.Y += float32(float64(yrel) * float64(monster.Speed))
			}
		}
	}
}
