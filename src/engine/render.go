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
	e.Inventory()
	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderNPC()
	e.RenderHealth()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	if e.ShowFPS {
		rl.DrawFPS(int32(rl.GetScreenWidth()/2+600), int32(rl.GetScreenHeight()/2-350))
	}
	rl.DrawRectangle(int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight()/2-350), 200, 25, rl.DarkBrown)
	rl.DrawRectangle(int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight()/2-350), int32(e.Player.Health)*2, 25, rl.Red)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-550-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2+350, 20, rl.RayWhite)
	rl.DrawText("MONEY : "+strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight())/2-300, 25, rl.RayWhite)
	rl.DrawText("DMG : "+strconv.Itoa(e.Player.Damage), int32(rl.GetScreenWidth()/2-650), int32(rl.GetScreenHeight())/2-300+40, 25, rl.RayWhite)
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

	for _, monster := range e.Monsters {
		if !(monster.Name == "distributeur") {
			if monster.IsAlive {
				if monster.Name == ("Mouse1") || monster.Name == ("Mouse2") || monster.Name == ("Mouse3") || monster.Name == ("Mouse4") ||
					monster.Name == ("Mouse5") {
					rl.DrawRectangle(int32(monster.Position.X)+15, int32(monster.Position.Y)+10, int32(20), 2, rl.DarkBrown)
					rl.DrawRectangle(int32(monster.Position.X)+15, int32(monster.Position.Y)+10, int32((monster.Health*20)/monster.MaxHealth), 2, rl.Red)
				}
				if monster.Name == "chien" {
					rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+20, int32(30), 4, rl.DarkBrown)
					rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+20, int32((monster.Health*30)/monster.MaxHealth), 4, rl.Red)

				}
			}
		}
	}

}
func (e *Engine) Inventory() {
	i := 0
	for _, item := range e.Player.Inventory {
		rl.DrawTexture(item.Sprite, int32(e.Player.Position.X)-150, int32(int(e.Player.Position.Y)+((-60)+(i*20))), rl.LightGray)
		i++
	}
}
