package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.SpriteRun = []rl.Texture2D{rl.LoadTexture("textures/entities/cat/2_Cat_Run-Sheet_Right.png"), rl.LoadTexture("textures/entities/cat/2_Cat_Run-Sheet_Left.png")}
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}
}
