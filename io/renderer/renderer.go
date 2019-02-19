package renderer

import (
	"github.com/daplf/go-tetris/game"
	"github.com/daplf/go-tetris/game/board"
	"github.com/daplf/go-tetris/game/piece"
	"github.com/daplf/go-tetris/game/piece/block"
	"github.com/daplf/go-tetris/utils/consts"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type windowSize = int
type colorType = float64

const (
	windowWidth  = 400
	windowHeight = 800
	windowTitle  = "Tetris"
	noColor      = "No color!"
)

// Renderer holds rendering logic
type Renderer struct {
	window *pixelgl.Window
}

// Window returns the window
func (renderer *Renderer) Window() *pixelgl.Window {
	return renderer.window
}

// CreateRenderer creates a new renderer with default dimensions
func CreateRenderer() *Renderer {
	window := setupPixel()

	return &Renderer{
		window: window,
	}
}

// setupPixel setsup the OpenGL context and window
func setupPixel() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  windowTitle,
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	window, error := pixelgl.NewWindow(cfg)
	if error != nil {
		panic(error)
	}

	return window
}

// DrawBoard draws the board on the screen
func (renderer *Renderer) DrawBoard(game *game.Game) {
	renderer.window.Clear(colornames.Black)

	squares := game.Board().Squares()
	width := game.Board().Width()
	height := game.Board().Height()

	for _, row := range squares {
		for _, block := range row {
			if block != nil {
				renderer.drawBlock(block, width, height)
			}
		}
	}

	renderer.window.Update()
}

// drawBlock draws a block on the screen
func (renderer *Renderer) drawBlock(block *block.Block, boardWidth, boardHeight board.Size) {
	imd := imdraw.New(nil)

	r, g, b, error := getBlockColor(block)

	if error == consts.NoError {
		imd.Color = pixel.RGB(r, g, b)

		blockWidth := float64(windowWidth / boardWidth)
		blockHeight := float64(windowHeight / boardHeight)
		x1 := float64(block.X()) * blockWidth
		y1 := float64(block.Y()) * blockHeight
		x2 := x1 + blockWidth
		y2 := y1 + blockHeight

		imd.Push(pixel.V(x1, y1))
		imd.Push(pixel.V(x2, y1))
		imd.Push(pixel.V(x2, y2))
		imd.Push(pixel.V(x1, y2))
		imd.Polygon(0)

		imd.Draw(renderer.window)
	}
}

// getBlockColor gets a block's color
func getBlockColor(block *block.Block) (colorType, colorType, colorType, consts.ErrorType) {
	var pieceColor = []colorType{0.0, 0.0, 0.0}
	error := consts.NoError

	switch block.Type() {
	case piece.PieceI:
		pieceColor = pieceIColor
	case piece.PieceJ:
		pieceColor = pieceJColor
	case piece.PieceL:
		pieceColor = pieceLColor
	case piece.PieceO:
		pieceColor = pieceOColor
	case piece.PieceS:
		pieceColor = pieceSColor
	case piece.PieceT:
		pieceColor = pieceTColor
	case piece.PieceZ:
		pieceColor = pieceZColor
	default:
		error = noColor
	}

	return pieceColor[0], pieceColor[1], pieceColor[2], error
}
