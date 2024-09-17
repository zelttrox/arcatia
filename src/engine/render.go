package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.Blue)

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}
func (e *Engine) GameOverRendering() {
	rl.ClearBackground(rl.Black)

	rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[Enter] to replay", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderHealth()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-350, 20, rl.RayWhite)
	rl.DrawText("MONEY : "+strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight())/2-300, 40, rl.RayWhite)
	rl.DrawText("PV : "+strconv.Itoa(e.Player.Health), int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight())/2-250, 40, rl.RayWhite)
}

func (e *Engine) PauseRendering() {
	rl.ClearBackground(rl.Red)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q]/[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	//rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.IsAlive {
			monster.UpdateAnimation()
			monster.Draw()
		}
	}
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	dialogBoxColor := rl.NewColor(200, 200, 185, 150)

	rl.DrawRectangle(450, 300, 200, 50, dialogBoxColor)

	rl.DrawText(
		sentence,
		int32(450),
		int32(300),
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RenderHealth() {
	rl.BeginMode2D(e.Camera)

	rl.DrawRectangle(int32(e.Player.Position.X)+45, int32(e.Player.Position.Y)+40, int32(50), 5, rl.DarkBrown)
	rl.DrawRectangle(int32(e.Player.Position.X)+45, int32(e.Player.Position.Y)+40, int32(e.Player.Health)/2, 5, rl.Red)

	for _, monster := range e.Monsters {
		if monster.IsAlive {
			rl.DrawRectangle(int32(monster.Position.X)+45, int32(monster.Position.Y)+40, int32(50), 5, rl.DarkBlue)
			rl.DrawRectangle(int32(monster.Position.X)+45, int32(monster.Position.Y)+40, int32((monster.Health*50)/monster.MaxHealth), 5, rl.Blue)
		}
	}

	rl.EndMode2D()
}
