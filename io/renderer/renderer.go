package renderer

import (
	"fmt"

	"github.com/daplf/go-tetris/game"
	"github.com/daplf/go-tetris/game/board"
	"github.com/daplf/go-tetris/game/piece"
	"github.com/daplf/go-tetris/game/piece/block"
	"github.com/daplf/go-tetris/utils/consts"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type windowSize = int
type colorType = float64

const (
	windowWidthPixels  = 500
	windowHeightPixels = 800
	boardWidthPixels   = 400
	boardHeightPixels  = 800
	scoreTextXPixels   = 430
	scoreTextYPixels   = 760
	pausedTextXPixels  = 430
	pausedTextYPixels  = 560
	windowTitle        = "Tetris"
	noColor            = "No color!"
)

// Renderer holds rendering logic
type Renderer struct {
	window     *pixelgl.Window
	scoreText  *text.Text
	pausedText *text.Text
}

// Window returns the window
func (renderer *Renderer) Window() *pixelgl.Window {
	return renderer.window
}

// CreateRenderer creates a new renderer with default dimensions
func CreateRenderer() *Renderer {
	window := setupWindow()
	scoreText := setupScoreText()
	pausedText := setupPausedText()

	return &Renderer{
		window:     window,
		scoreText:  scoreText,
		pausedText: pausedText,
	}
}

// setupWindow sets up the OpenGL context and window
func setupWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  windowTitle,
		Bounds: pixel.R(0, 0, windowWidthPixels, windowHeightPixels),
		VSync:  true,
	}

	window, error := pixelgl.NewWindow(cfg)
	if error != nil {
		panic(error)
	}

	return window
}

// setupScoreText sets up the text used for the score.
func setupScoreText() *text.Text {
	scoreAtlas := text.NewAtlas(
		basicfont.Face7x13,
		[]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
	)

	scoreText := text.New(pixel.V(scoreTextXPixels, scoreTextYPixels), scoreAtlas)
	scoreText.Color = colornames.Yellow

	return scoreText
}

// setupPauseText sets up the text used for the paused status indicator.
func setupPausedText() *text.Text {
	pausedAtlas := text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)

	pausedText := text.New(pixel.V(pausedTextXPixels, pausedTextYPixels), pausedAtlas)
	pausedText.Color = colornames.Yellow

	return pausedText
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

	renderer.drawInfoTab(game)

	renderer.window.Update()
}

// drawBlock draws a block on the screen
func (renderer *Renderer) drawBlock(block *block.Block, boardWidth, boardHeight board.Size) {
	r, g, b, error := getBlockColor(block)

	if error == consts.NoError {
		blockWidth := float64(boardWidthPixels / boardWidth)
		blockHeight := float64(boardHeightPixels / boardHeight)
		x1 := float64(block.X()) * blockWidth
		y1 := float64(block.Y()) * blockHeight
		x2 := x1 + blockWidth
		y2 := y1 + blockHeight

		drawPolygon(
			renderer.window,
			pixel.RGB(r, g, b),
			[][2]float64{
				{x1, y1},
				{x2, y1},
				{x2, y2},
				{x1, y2},
			},
		)
	}
}

// drawScore draws the score on the screen.
func (renderer *Renderer) drawInfoTab(game *game.Game) {
	drawPolygon(
		renderer.window,
		pixel.RGB(0, 0, 1),
		[][2]float64{
			{boardWidthPixels, 0},
			{windowHeightPixels, 0},
			{windowHeightPixels, windowHeightPixels},
			{boardWidthPixels, windowHeightPixels},
		},
	)

	renderer.scoreText.Clear()
	fmt.Fprintln(renderer.scoreText, game.Score())
	renderer.scoreText.Draw(renderer.window, pixel.IM.Scaled(renderer.scoreText.Orig, 2))

	pausedText := ""
	if game.IsPaused() {
		pausedText = "Paused"
	}

	renderer.pausedText.Clear()
	fmt.Fprintln(renderer.pausedText, pausedText)
	renderer.pausedText.Draw(renderer.window, pixel.IM.Scaled(renderer.pausedText.Orig, 1.5))
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

// drawPolygon draws a poligon on the given target
func drawPolygon(target pixel.Target, color pixel.RGBA, vertices [][2]float64) {
	imd := imdraw.New(nil)

	imd.Color = color

	for _, row := range vertices {
		imd.Push(pixel.V(row[0], row[1]))
	}

	imd.Polygon(0)

	imd.Draw(target)
}
