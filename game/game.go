package game

import (
	"math/rand"
	"time"

	"github.com/daplf/go-tetris/game/board"
	"github.com/daplf/go-tetris/game/piece"
	"github.com/daplf/go-tetris/game/piece/block"
)

const (
	// MoveDown represents a downwards move
	MoveDown = "MoveDown"

	// MoveRight represents a rightwards move
	MoveRight = "MoveRight"

	// MoveLeft represents a leftwards move
	MoveLeft = "MoveLeft"

	// RotateLeft represents a leftwards rotation move
	RotateLeft = "RotateLeft"

	// RotateRight represents a rightwards rotation move
	RotateRight = "RotateRight"

	// NoMove represents no move
	NoMove = ""

	// Closed is a flag used to tell the game to finish (because the window was closed)
	Closed = "Closed"

	oneSecond = 1

	scoreMultiplier = 20
)

// Move type
type Move = string

// Game holds the game logic
type Game struct {
	running      bool
	board        *board.Board
	currentPiece *piece.Piece
	lastTime     time.Time
	score        int
}

// IsRunning checks if game is running
func (game *Game) IsRunning() bool {
	return game.running
}

// Board returns the game's board
func (game *Game) Board() *board.Board {
	return game.board
}

// Score returns the game's score
func (game *Game) Score() int {
	return game.score
}

// CreateGame creates a new game
func CreateGame() *Game {
	rand.Seed(time.Now().UnixNano())

	board := board.CreateBoard()

	currentPiece := generateNewPiece(board)

	return &Game{
		running:      true,
		board:        board,
		currentPiece: currentPiece,
		lastTime:     time.Now(),
	}
}

// CreateGameWithDimensions creates a new game using custom dimensions for the board
func CreateGameWithDimensions(width, height board.Size) *Game {
	rand.Seed(time.Now().UnixNano())

	board := board.CreateBoardWithDimensions(width, height)

	currentPiece := generateNewPiece(board)

	return &Game{
		running:      true,
		board:        board,
		currentPiece: currentPiece,
		lastTime:     time.Now(),
	}
}

func generateNewPiece(board *board.Board) *piece.Piece {
	random := rand.Intn(piece.NumPieces)

	var pieceType block.Type
	var pieceCoords [][]block.Position

	switch random {
	case 0:
		pieceType = piece.PieceI
		pieceCoords = piece.PieceICoords[0]
		break
	case 1:
		pieceType = piece.PieceJ
		pieceCoords = piece.PieceJCoords[0]
		break
	case 2:
		pieceType = piece.PieceL
		pieceCoords = piece.PieceLCoords[0]
		break
	case 3:
		pieceType = piece.PieceO
		pieceCoords = piece.PieceOCoords[0]
		break
	case 4:
		pieceType = piece.PieceS
		pieceCoords = piece.PieceSCoords[0]
		break
	case 5:
		pieceType = piece.PieceT
		pieceCoords = piece.PieceTCoords[0]
		break
	case 6:
		pieceType = piece.PieceZ
		pieceCoords = piece.PieceZCoords[0]
		break
	}

	blocks := make([]*block.Block, 4)

	for i := 0; i < 4; i++ {
		x := board.Width()/2 + pieceCoords[0][i]
		y := board.Height() - 2 + pieceCoords[1][i]

		if board.Squares()[y][x] != nil {
			return nil
		}

		blocks[i] = block.CreateBlock(x, y, pieceType)
	}

	piece := piece.CreatePiece(blocks)

	board.SetSquares(piece.Blocks())

	return piece
}

// Update updates the game
func (game *Game) Update(move Move) {
	if move != NoMove {
		if move == Closed {
			game.running = false
			return
		}

		game.makeMove(move)
	}

	game.fallCurrentPiece()
}

func (game *Game) makeMove(move Move) bool {
	res := false

	switch move {
	case MoveDown:
		res = game.movePieceDown()
		if !res {
			game.lastTime = time.Now()
			game.executeFallCurrentPiece()
		}
		break
	case MoveRight:
		res = game.movePieceRight()
		break
	case MoveLeft:
		res = game.movePieceLeft()
		break
	case RotateLeft:
		res = game.rotatePieceLeft()
		break
	case RotateRight:
		res = game.rotatePieceRight()
		break
	}

	return res
}

// movePieceDown moves current piece down
func (game *Game) movePieceDown() bool {
	blocks := game.currentPiece.Blocks()

	return game.board.MoveBlocksDown(blocks)
}

// movePieceRight moves current piece right
func (game *Game) movePieceRight() bool {
	blocks := game.currentPiece.Blocks()

	return game.board.MoveBlocksRight(blocks)
}

// movePieceLeft moves current piece left
func (game *Game) movePieceLeft() bool {
	blocks := game.currentPiece.Blocks()

	return game.board.MoveBlocksLeft(blocks)
}

// rotatePieceRight moves current piece right
func (game *Game) rotatePieceRight() bool {
	blocks := game.currentPiece.Blocks()

	oldState := game.currentPiece.State()
	newState := game.board.RotateBlocksRight(blocks, oldState)

	game.currentPiece.SetState(newState)

	return newState != oldState
}

// rotatePieceLeft moves current piece left
func (game *Game) rotatePieceLeft() bool {
	blocks := game.currentPiece.Blocks()

	oldState := game.currentPiece.State()
	newState := game.board.RotateBlocksLeft(blocks, oldState)

	game.currentPiece.SetState(newState)

	return newState != oldState
}

// fallCurrentPiece moves current piece down if possible
func (game *Game) fallCurrentPiece() {
	now := time.Now()

	if now.Sub(game.lastTime).Seconds() > oneSecond {
		game.lastTime = now
		res := game.movePieceDown()

		if !res {
			game.executeFallCurrentPiece()
		}
	}
}

// executeFallCurrentPiece moves current piece down
func (game *Game) executeFallCurrentPiece() {
	numRowsDestroyed := game.Board().DestroyFullRows()
	if numRowsDestroyed > 0 {
		game.score += numRowsDestroyed * scoreMultiplier
	}

	game.currentPiece = generateNewPiece(game.board)

	if game.currentPiece == nil {
		game.running = false
	}
}
