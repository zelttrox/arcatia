package engine

import (
	"main/src/entity"

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

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}

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

	e.MonsterCollisions()
}

func (e *Engine) MonsterCollisions() {

	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-20 &&
			monster.Position.X < e.Player.Position.X+20 &&
			monster.Position.Y > e.Player.Position.Y-20 &&
			monster.Position.Y < e.Player.Position.Y+20 {

			if monster.Name == "claude" {
				e.NormalTalk(monster, "Bonjour")
				if rl.IsKeyPressed(rl.KeyE) {
					//lancer un combat ?
				}
			}
		} else {
			//...
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

	//Musique
	rl.UpdateMusicStream(e.Music)
}
