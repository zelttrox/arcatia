package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Collisions() {

	switch SALLE {

	case 1:

		if e.Player.Position.X > 80 && e.Player.Position.X < 120 && e.Player.Position.Y > 180 && e.Player.Position.Y < 220 {
			SALLE = 2
			e.Player.Position = rl.Vector2{X: 108, Y: 472}
		}

		if e.Player.Position.Y > 69 {
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				e.Player.Position.Y -= e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.Y < 198 {
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
		if e.Player.Position.X < 340 {
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				e.Player.Position.X += e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 0
			}

		}
	case 2:
		if e.Player.Position.X > 210 && e.Player.Position.X < 240 && e.Player.Position.Y > 610 && e.Player.Position.Y < 630 {
			SALLE = 3
			e.Player.Position = rl.Vector2{X: 184, Y: 1024}
		}
		if e.Player.Position.Y > 472 {
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				e.Player.Position.Y -= e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.Y < 612 {
			if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
				e.Player.Position.Y += e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.X > 56 {
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				e.Player.Position.X -= e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 1
			}
		}
		if e.Player.Position.X < 408 {
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				e.Player.Position.X += e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 0
			}
		}
	case 3:
		if e.Player.Position.Y > 792 {
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				e.Player.Position.Y -= e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.Y < 1260 {
			if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
				e.Player.Position.Y += e.Player.Speed
				e.Player.IsRunning = true
			}
		}
		if e.Player.Position.X > -10 {
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				e.Player.Position.X -= e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 1
			}
		}
		if e.Player.Position.X < 460 {
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				e.Player.Position.X += e.Player.Speed
				e.Player.IsRunning = true
				e.Player.Dir = 0
			}
		}
	}
}
