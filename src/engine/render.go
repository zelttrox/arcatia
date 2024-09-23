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
	rl.DrawTexture(e.Player.HomescreenSprite, 0, 0, rl.White)

	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-505, int32(rl.GetScreenHeight())/2-40, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-505, int32(rl.GetScreenHeight())/2+100-40, 40, rl.RayWhite)

}
func (e *Engine) GameOverRendering() {
	rl.DrawTexture(e.Player.GameoverSprite, 0, 0, rl.White)

}
func (e *Engine) GoodGameRendering() {
	rl.DrawTexture(e.Player.GoodGameSprite, 0, 0, rl.White)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Brown)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderHealth()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawFPS(int32(rl.GetScreenWidth()/2+600), int32(rl.GetScreenHeight()/2-350))
	rl.DrawRectangle(int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight()/2-350), 200, 25, rl.DarkBrown)
	rl.DrawRectangle(int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight()/2-350), int32(e.Player.Health)*2, 25, rl.Red)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-550-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2+350, 20, rl.RayWhite)
	rl.DrawText("MONEY : "+strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight())/2-300, 25, rl.RayWhite)
}

func (e *Engine) PauseRendering() {

	rl.DrawTexture(e.Player.PausescreenSprite, 0, 0, rl.White)

	rl.DrawText("[Enter] to Resume", int32(rl.GetScreenWidth())/2-150, int32(rl.GetScreenHeight())/2-40, 40, rl.RayWhite)
	rl.DrawText("[R] to Restart", int32(rl.GetScreenWidth())/2-150, int32(rl.GetScreenHeight())/2+100-40, 40, rl.RayWhite)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-150, int32(rl.GetScreenHeight())/2+200-40, 40, rl.RayWhite)

}

func (e *Engine) RenderPlayer() {

	if e.Player.IsRunning {
		e.Player.PlayerAnimation()
		e.Player.PlayerDraw()
	} else {
		e.Player.PlayerIdle()
		e.Player.PlayerDraw()
	}

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.IsAlive {
			monster.UpdateAnimation()
			monster.Draw()
			if monster.Name == "distributeur" {
				dialogBoxColor := rl.NewColor(200, 200, 185, 150)
				rl.DrawRectangle(int32(monster.Position.X)-25, int32(monster.Position.Y)-10, 191, 9, dialogBoxColor)

				rl.DrawText(
					"[E] Food(+Heal) or [R] Drink(+Damage)",
					int32(monster.Position.X)-25,
					int32(monster.Position.Y)-10,
					int32(1),
					rl.DarkBrown,
				)
			}
		}
	}
}

func (e *Engine) RenderNPC() {

	e.NPC.NPCIdle()
	e.NPC.NPCDraw()
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	dialogBoxColor := rl.NewColor(200, 200, 185, 150)

	for _, monster := range e.Monsters {
		if monster.IsAlive {
			rl.DrawRectangle(int32(monster.Position.X), int32(monster.Position.Y), 200, 50, dialogBoxColor)

			rl.DrawText(
				sentence,
				int32(450),
				int32(300),
				10,
				rl.RayWhite,
			)

			rl.EndMode2D()
		}
	}
}

func (e Engine) RenderHealth() {
	rl.BeginMode2D(e.Camera)

	for _, monster := range e.Monsters {
		if !(monster.Name == "distributeur") {
			if monster.IsAlive {
				rl.DrawRectangle(int32(monster.Position.X)+45, int32(monster.Position.Y)+40, int32(50), 5, rl.DarkBlue)
				rl.DrawRectangle(int32(monster.Position.X)+45, int32(monster.Position.Y)+40, int32((monster.Health*50)/monster.MaxHealth), 5, rl.Blue)
			}
		}
	}

	rl.EndMode2D()
}
