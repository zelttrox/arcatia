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
		if x >= 160 && x <= 170 && y >= 1020 && y < 1026 {
			SALLE = 2
			e.Player.Position = rl.Vector2{X: 204, Y: 580}
		}

		if x >= 316 && x <= 320 && y < 1156 || x >= 394 && x <= 396 && y < 1016 || x >= 350 && x <= 354 && y < 1000 ||
			x >= 300 && x <= 304 && y < 948 || x >= 264 && x <= 268 && y < 900 || x >= 250 && x <= 254 && y < 884 ||
			x >= 234 && x <= 238 && y < 872 || x >= 186 && x <= 190 && y < 856 || x >= 336 && x <= 338 && y > 1154 && y < 1174 ||
			x >= 88 && x <= 90 && y < 1018 && y > 922 {

			CanGoRight = false
		}
		if x <= 340 && x >= 336 && y < 1156 || x <= 20 && x >= 18 && y < 1065 || x <= 42 && x >= 40 && y < 922 ||
			x <= 56 && x >= 54 && y < 904 || x <= 66 && x >= 64 && y < 887 || x <= 104 && x >= 100 && y < 870 ||
			x <= 134 && x >= 130 && y < 855 || x <= 250 && x >= 248 && y > 922 && y < 1018 ||
			x <= 352 && x >= 348 && y > 1148 && y < 1168 {
			CanGoLeft = false
		}

		if x < 20 && y < 1072 || x < 34 && y < 928 || x < 48 && y < 912 || x < 68 && y < 896 || x < 96 && y < 880 || x < 128 && y < 864 || x < 190 && y < 848 ||
			x > 192 && y < 862 || x > 234 && y < 876 || x > 250 && y < 894 || x > 268 && y < 914 || x > 300 && y < 960 || x > 316 && y < 990 || x > 348 && y < 1012 || x > 396 && y < 1026 ||
			y < 1164 && y > 1160 && x > 316 && x < 340 || y < 1180 && y > 1176 && x > 334 && x < 364 ||
			x < 250 && x > 88 && y <= 1020 && y >= 1018 {
			CanGoUp = false
		}
		if x > 88 && x < 250 && y >= 922 && y <= 924 || x < 370 && x > 350 && y <= 1164 && y >= 1162 || x < 350 && x > 330 && y <= 1150 && y >= 1148 {
			CanGoDown = false
		}

		if CanGoUp {
			if y > 780 {
				if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
					e.Player.Position.Y -= e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoDown {
			if y < 1250 {
				if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
					e.Player.Position.Y += e.Player.Speed
					e.Player.IsRunning = true
				}
			}
		}
		if CanGoLeft {
			if x > -24 {
				if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
					e.Player.Position.X -= e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 1
				}
			}
		}
		if CanGoRight {
			if x < 440 {
				if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
					e.Player.Position.X += e.Player.Speed
					e.Player.IsRunning = true
					e.Player.Dir = 0
				}
			}
		}
	}
}
