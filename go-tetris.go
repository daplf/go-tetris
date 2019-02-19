package main

import (
	"github.com/daplf/go-tetris/game"
	"github.com/daplf/go-tetris/io/inputProcessor"
	"github.com/daplf/go-tetris/io/renderer"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	game := game.CreateGame()
	renderer := renderer.CreateRenderer()

	renderer.DrawBoard(game)

	for game.IsRunning() {
		renderer.DrawBoard(game)
		move := inputProcessor.GetInput(renderer)
		game.Update(move)
	}
}

func main() {
	pixelgl.Run(run)
}
