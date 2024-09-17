package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/fight"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
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

func (e *Engine) reset() {
	e.Player.Health = 100
	e.Player.Money = 300
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
		e.reset()
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

	e.Player.IsRunning = false

	if e.Player.Position.Y > 10 {
		if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
			e.Player.Position.Y -= e.Player.Speed
			e.Player.IsRunning = true
		}
	}
		if e.Player.Position.Y < 230 {
			if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
				e.Player.Position.Y += e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.X > 90 {
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				e.Player.Position.X -= e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 1
			}
		}
		if e.Player.Position.X < 375 {
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				e.Player.Position.X += e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 0
			}
		}

	

	// Mouvement
	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-07-Simon_s-In-There-Somewhere.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {
	fmt.Println(e.Player.Position)
	e.ChasePlayer()
	e.MonsterCollisions()
}

func (e *Engine) MonsterCollisions() {
	for i := range e.Monsters {
		monster := &e.Monsters[i]
		if monster.IsAlive {
			if monster.Position.X > e.Player.Position.X-40 &&
				monster.Position.X < e.Player.Position.X+40 &&
				monster.Position.Y > e.Player.Position.Y-40 &&
				monster.Position.Y < e.Player.Position.Y+40 {
				if rl.IsKeyPressed(rl.KeyE) {
					fight.Fightp(&e.Player, monster)

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

			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}
	if rl.IsKeyPressed(rl.KeyR) {
		e.reset()
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) ChasePlayer() {

	for i := range e.Monsters {
		monster := &e.Monsters[i]
		distrel := math.Sqrt(math.Pow(float64(e.Player.Position.X)-float64(monster.Position.X), 2) + math.Pow(float64(e.Player.Position.Y)-float64(monster.Position.Y), 2))
		if distrel <= 200 && distrel > 20 {
			xrel := float64(e.Player.Position.X-monster.Position.X) / distrel
			yrel := float64(e.Player.Position.Y-monster.Position.Y) / distrel
			if float32(float64(xrel)*float64(monster.Speed)) < 0 {
				//monster.Sprite = rl.LoadTexture("textures/entities/orc/Orc-Idle-Reverse.png")
			}
			if float32(float64(xrel)*float64(monster.Speed)) >= 0 {
				//monster.Sprite = rl.LoadTexture("textures/entities/orc/Orc-Idle.png")
			}
			monster.Position.X += float32(float64(xrel) * float64(monster.Speed))
			monster.Position.Y += float32(float64(yrel) * float64(monster.Speed))
		}
	}
}
