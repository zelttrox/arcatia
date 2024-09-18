package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Collisions() {
	CanGoUp := true
	CanGoDown := true
	CanGoLeft := true
	CanGoRight := true

	x := e.Player.Position.X
	y := e.Player.Position.Y

	switch SALLE {

	case 1:

		if e.Player.Position.X >= 68 && e.Player.Position.X <= 80 && e.Player.Position.Y >= 180 && e.Player.Position.Y <= 200 {
			SALLE = 2
			e.Player.Position = rl.Vector2{X: 108, Y: 472}
		}
		if x >= 130 && x <= 132 && y < 134 || x <= 84 && x >= 80 && y > 120 {
			CanGoRight = false
		}
		if x >= 148 && x <= 152 && y < 134 || x <= 104 && x >= 100 && y > 120 {
			CanGoLeft = false
		}
		if x > 212 && y >= 148 && y <= 152 || x > 132 && x < 184 && y >= 148 && y <= 152 {
			CanGoUp = false
		}
		if x > 212 && y >= 124 && y <= 128 || x > 132 && x < 184 && y >= 124 && y <= 128 {
			CanGoDown = false
		}

		if CanGoUp {
			if e.Player.Position.Y > 69 {
				if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
					e.Player.Position.Y -= e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoDown {
			if e.Player.Position.Y < 198 {
				if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
					e.Player.Position.Y += e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoLeft {
			if e.Player.Position.X > 70 {
				if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
					e.Player.Position.X -= e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 1
				}
			}
		}
		if CanGoRight {
			if e.Player.Position.X < 328 {
				if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
					e.Player.Position.X += e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 0
				}

			}
		}
	case 2:
		if x >= 200 && x <= 220 && y >= 592 && y <= 596 {
			SALLE = 3
			e.Player.Position = rl.Vector2{X: 184, Y: 1024}
		}
		if x <= 80 && x <= 88 && y >= 460 && y <= 470 {
			SALLE = 1
			e.Player.Position = rl.Vector2{X: 100, Y: 132}
		}

		if x > 250 && y <= 484 || x < 92 && y < 484 || (x < 100 || x > 120 && x < 188 || x > 228 && x < 264) && y <= 544 && y >= 540 {
			CanGoUp = false
		}
		if x >= 248 && y < 482 || (x >= 168 && x <= 172 || x >= 228 && x <= 232) && y > 532 {
			CanGoRight = false
		}
		if x < 78 && y >= 572 || (x < 100 || x > 120 && x < 188 || x > 228 && x < 264) && y <= 532 && y >= 528 {
			CanGoDown = false
		}
		if x <= 80 && y > 574 || (x >= 184 && x <= 188 || x >= 248 && x <= 252) && y > 532 {
			CanGoLeft = false
		}

		if CanGoUp {
			if y > 462 {
				if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
					e.Player.Position.Y -= e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoDown {
			if y < 596 {
				if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
					e.Player.Position.Y += e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoLeft {
			if x > 42 {
				if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
					e.Player.Position.X -= e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 1
				}
			}
		}
		if CanGoRight {
			if x < 394 {
				if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
					e.Player.Position.X += e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 0
				}
			}
		}
	case 3:
		if x > 160 && x < 180 && y > 1000 && y < 1010 {
			SALLE = 2
			e.Player.Position = rl.Vector2{X: 204, Y: 580}
		}

		if y > 792 {
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				e.Player.Position.Y -= e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if y < 1260 {
			if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
				e.Player.Position.Y += e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if x > -10 {
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				e.Player.Position.X -= e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 1
			}
		}
		if x < 460 {
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				e.Player.Position.X += e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 0
			}
		}
	}
}
