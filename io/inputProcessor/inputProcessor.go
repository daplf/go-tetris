package inputProcessor

import (
	"github.com/daplf/go-tetris/game"
	"github.com/daplf/go-tetris/io/renderer"
	"github.com/faiface/pixel/pixelgl"
)

// GetInput checks if there is new input and returns it
func GetInput(renderer *renderer.Renderer) game.Move {
	move := getJustPressed(renderer)

	if move == game.NoMove {
		move = getRepeated(renderer)
	}

	if renderer.Window().Closed() {
		move = game.Closed
	}

	return move
}

func getJustPressed(renderer *renderer.Renderer) game.Move {
	var move game.Move

	if renderer.Window().JustPressed(pixelgl.KeyDown) {
		move = game.MoveDown
	}

	if renderer.Window().JustPressed(pixelgl.KeyRight) {
		move = game.MoveRight
	}

	if renderer.Window().JustPressed(pixelgl.KeyLeft) {
		move = game.MoveLeft
	}

	if renderer.Window().JustPressed(pixelgl.KeyA) || renderer.Window().JustPressed(pixelgl.KeyS) {
		move = game.RotateLeft
	}

	if renderer.Window().JustPressed(pixelgl.KeyW) || renderer.Window().JustPressed(pixelgl.KeyD) {
		move = game.RotateRight
	}

	if renderer.Window().JustPressed(pixelgl.KeyP) {
		move = game.Paused
	}

	return move
}

func getRepeated(renderer *renderer.Renderer) game.Move {
	var move game.Move

	if renderer.Window().Repeated(pixelgl.KeyDown) {
		move = game.MoveDown
	}

	if renderer.Window().Repeated(pixelgl.KeyRight) {
		move = game.MoveRight
	}

	if renderer.Window().Repeated(pixelgl.KeyLeft) {
		move = game.MoveLeft
	}

	if renderer.Window().Repeated(pixelgl.KeyA) || renderer.Window().Repeated(pixelgl.KeyS) {
		move = game.RotateLeft
	}

	if renderer.Window().Repeated(pixelgl.KeyW) || renderer.Window().Repeated(pixelgl.KeyD) {
		move = game.RotateRight
	}

	return move
}
