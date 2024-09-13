package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layer struct {
	Data    []int   `json:"data"`
	Height  int     `json:"height"`
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Opacity float32 `json:"opacity"`
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Width   int     `json:"width"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

type TileSet struct {
	Columns     int    `json:"columns"`
	FirstGid    int    `json:"firstgid"`
	Image       string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
	Margin      int    `json:"margin"`
	Name        string `json:"name"`
	Spacing     int    `json:"spacing"`
	TileCount   int    `json:"tilecount"`
	TileHeight  int    `json:"tileheight"`
	TileWidth   int    `json:"tilewidth"`
}

type MapJSON struct {
	CompressionLevel int       `json:"compressionLevel"`
	Height           int       `json:"height"`
	Infinite         bool      `json:"infinite"`
	Layers           []Layer   `json:"layers"`
	NextLayerID      int       `json:"nextlayerid"`
	NextObjectID     int       `json:"nextobjectid"`
	Orientation      string    `json:"orientation"`
	RenderOrder      string    `json:"renderorder"`
	TiledVersion     string    `json:"tiledversion"`
	TileHeight       int       `json:"tileheight"`
	TileSets         []TileSet `json:"tilesets"`
	TileWidth        int       `json:"tilewidth"`
	Type             string    `json:"type"`
	Version          string    `json:"version"`
	Width            int       `json:"width"`
}

func (e *Engine) InitMap(mapFile string) {
	/*
		Naive & slow map loader, render all layers everywhere each frame:
		- Parse JSON
		- Load required textures
		- For each layer
			- For each tile
				- Find closest TileSet GID to select correct texture
				- Get X and Y coordinates of the tile
				- Draw tile
				- Move to next position (line, or column)
	*/
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Unmarshal(file, &e.MapJSON)

	//Load all required textures from TileSets
	for _, TileSet := range e.MapJSON.TileSets {
		path := path.Dir(mapFile) + "/"
		e.Sprites[TileSet.Name] = rl.LoadTexture(path + TileSet.Image)
	}
}

func (e *Engine) RenderMap() {
	/*
		Naive & slow map loader, render all layers everywhere each frame:
		- Parse JSON
		- Load required textures
		- For each layer
			- For each tile
				- Find closest TileSet GID to select correct texture
				- Get X and Y coordinates of the tile
				- Draw tile
				- Move to next position (line, or column)
	*/

	// Prepare source and destination rectangle (only X and Y will change on both)
	srcRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileHeight), Height: float32(e.MapJSON.TileHeight)}
	destRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileWidth), Height: float32(e.MapJSON.TileWidth)}
	column_counter := -1

	for _, Layer := range e.MapJSON.Layers {
		for _, tile := range Layer.Data {
			if tile != 0 {
				wantedTileSet := e.MapJSON.TileSets[0]
				for _, TileSet := range e.MapJSON.TileSets { // Get correct texture
					if TileSet.FirstGid < tile {
						wantedTileSet = TileSet
					}
				}

				index := tile - wantedTileSet.FirstGid

				srcRectangle.X = float32(index)
				srcRectangle.Y = 0

				if index > wantedTileSet.Columns { // If Tile number exceeds columns (overflow), adjust, find X and Y coordinates
					srcRectangle.X = float32(index % wantedTileSet.Columns)
					srcRectangle.Y = float32(index / wantedTileSet.Columns)
				}

				srcRectangle.X *= float32(e.MapJSON.TileWidth)
				srcRectangle.Y *= float32(e.MapJSON.TileHeight)

				rl.DrawTexturePro(
					e.Sprites[wantedTileSet.Name],
					srcRectangle,
					destRectangle,
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
			}

			// After each draw, move to the right. When at max width, new line (like a typewriter)
			destRectangle.X += 32
			column_counter += 1
			if column_counter >= e.MapJSON.Width {
				destRectangle.X = 0
				destRectangle.Y += 32
				column_counter = 0
			}

		}
		destRectangle.X, destRectangle.Y, column_counter = 0, 0, 0
	}
}
