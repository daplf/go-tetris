package board

import (
	"github.com/daplf/go-tetris/game/piece"
	"github.com/daplf/go-tetris/game/piece/block"
)

// Size is a measure for board sizes (height or width)
type Size = int

const (
	boardWidth  = 10
	boardHeight = 20
)

// Board holds the board logic
type Board struct {
	width   Size
	height  Size
	squares [][]*block.Block
}

// Width returns the width of the board
func (board *Board) Width() Size {
	return board.width
}

// Height returns the height of the board
func (board *Board) Height() Size {
	return board.height
}

// Squares returns the squares in the board
func (board *Board) Squares() [][]*block.Block {
	return board.squares
}

// CreateBoard creates a new default board
func CreateBoard() *Board {
	squares := initSquares(boardWidth, boardHeight)

	return &Board{
		width:   boardWidth,
		height:  boardHeight,
		squares: squares,
	}
}

// CreateBoardWithDimensions creates a new board with custom dimensions
func CreateBoardWithDimensions(width Size, height Size) *Board {
	squares := initSquares(width, height)

	return &Board{
		width:   width,
		height:  height,
		squares: squares,
	}
}

// initSquares creates a 2D array containing the squares of the board based on input dimensions
func initSquares(width, height Size) [][]*block.Block {
	squares := make([][]*block.Block, height)

	for row := range squares {
		squares[row] = make([]*block.Block, width)
	}

	return squares
}

// SetSquares sets some squares to piece blocks
func (board *Board) SetSquares(blocks []*block.Block) {
	for _, block := range blocks {
		board.squares[block.Y()][block.X()] = block
	}
}

// MoveBlocksDown moves blocks down if possible
func (board *Board) MoveBlocksDown(blocks []*block.Block) bool {
	movePossible := true

	for _, block := range blocks {
		if block.Y() <= 0 {
			movePossible = false
			break
		}

		if board.squares[block.Y()-1][block.X()] != nil {
			ok := false

			for _, neighbour := range blocks {
				if board.squares[block.Y()-1][block.X()] == neighbour {
					ok = true
				}
			}

			if !ok {
				movePossible = false
				break
			}
		}
	}

	if movePossible {
		for i := 0; i < len(blocks); i++ {
			board.squares[blocks[i].Y()][blocks[i].X()] = nil
		}

		for i := 0; i < len(blocks); i++ {
			blocks[i].MoveDown()
			board.squares[blocks[i].Y()][blocks[i].X()] = blocks[i]
		}
	}

	return movePossible
}

// MoveBlocksRight moves blocks right if possible
func (board *Board) MoveBlocksRight(blocks []*block.Block) bool {
	movePossible := true

	for _, block := range blocks {
		if block.X() >= board.width-1 {
			movePossible = false
			break
		}

		if board.squares[block.Y()][block.X()+1] != nil {
			ok := false

			for _, neighbour := range blocks {
				if board.squares[block.Y()][block.X()+1] == neighbour {
					ok = true
				}
			}

			if !ok {
				movePossible = false
				break
			}
		}
	}

	if movePossible {
		for i := 0; i < len(blocks); i++ {
			board.squares[blocks[i].Y()][blocks[i].X()] = nil
		}

		for i := 0; i < len(blocks); i++ {
			blocks[i].MoveRight()
			board.squares[blocks[i].Y()][blocks[i].X()] = blocks[i]
		}
	}

	return movePossible
}

// MoveBlocksLeft moves blocks right if possible
func (board *Board) MoveBlocksLeft(blocks []*block.Block) bool {
	movePossible := true

	for _, block := range blocks {
		if block.X() <= 0 {
			movePossible = false
			break
		}

		if board.squares[block.Y()][block.X()-1] != nil {
			ok := false

			for _, neighbour := range blocks {
				if board.squares[block.Y()][block.X()-1] == neighbour {
					ok = true
				}
			}

			if !ok {
				movePossible = false
				break
			}
		}
	}

	if movePossible {
		for i := 0; i < len(blocks); i++ {
			board.squares[blocks[i].Y()][blocks[i].X()] = nil
		}

		for i := 0; i < len(blocks); i++ {
			blocks[i].MoveLeft()
			board.squares[blocks[i].Y()][blocks[i].X()] = blocks[i]
		}
	}

	return movePossible
}

// RotateBlocksRight rotates blocks left if possible
func (board *Board) RotateBlocksRight(blocks []*block.Block, state piece.State) piece.State {
	newState := state + 1

	if newState > 3 {
		newState = piece.NormalState
	}

	return board.rotate(blocks, state, newState)
}

// RotateBlocksLeft rotates blocks left if possible
func (board *Board) RotateBlocksLeft(blocks []*block.Block, state piece.State) piece.State {
	newState := state - 1

	if newState < 0 {
		newState = piece.NumStates - 1
	}

	return board.rotate(blocks, state, newState)
}

// rotate rotates a block to a new position
func (board *Board) rotate(blocks []*block.Block, state, newState piece.State) piece.State {
	movePossible := true

	blockType := blocks[0].Type()

	oldCoords, newCoords := getOldAndNewCoords(blockType, state, newState)

	for i := range blocks {
		newX := blocks[i].X() - (oldCoords[0][i] - newCoords[0][i])
		newY := blocks[i].Y() - (oldCoords[1][i] - newCoords[1][i])

		if newX >= board.width || newX < 0 || newY >= board.height || newY < 0 {
			movePossible = false
			break
		}

		if board.squares[newY][newX] != nil {
			ok := false

			for _, neighbour := range blocks {
				if board.squares[newY][newX] == neighbour {
					ok = true
				}
			}

			if !ok {
				movePossible = false
				break
			}
		}
	}

	if movePossible {
		for i := 0; i < len(blocks); i++ {
			board.squares[blocks[i].Y()][blocks[i].X()] = nil
		}

		for i := 0; i < len(blocks); i++ {
			newX := blocks[i].X() - (oldCoords[0][i] - newCoords[0][i])
			newY := blocks[i].Y() - (oldCoords[1][i] - newCoords[1][i])
			blocks[i].SetX(newX)
			blocks[i].SetY(newY)
			board.squares[blocks[i].Y()][blocks[i].X()] = blocks[i]
		}

		state = newState
	}

	return state
}

// getOldAndNewCoords gets the old and new coords of the piece
func getOldAndNewCoords(blockType block.Type, oldState, newState piece.State) ([][]block.Position, [][]block.Position) {
	var oldCoords [][]block.Position
	var newCoords [][]block.Position

	switch blockType {
	case piece.PieceI:
		oldCoords = piece.PieceICoords[oldState]
		newCoords = piece.PieceICoords[newState]
		break
	case piece.PieceJ:
		oldCoords = piece.PieceJCoords[oldState]
		newCoords = piece.PieceJCoords[newState]
		break
	case piece.PieceL:
		oldCoords = piece.PieceLCoords[oldState]
		newCoords = piece.PieceLCoords[newState]
		break
	case piece.PieceO:
		oldCoords = piece.PieceOCoords[oldState]
		newCoords = piece.PieceOCoords[newState]
		break
	case piece.PieceS:
		oldCoords = piece.PieceSCoords[oldState]
		newCoords = piece.PieceSCoords[newState]
		break
	case piece.PieceT:
		oldCoords = piece.PieceTCoords[oldState]
		newCoords = piece.PieceTCoords[newState]
		break
	case piece.PieceZ:
		oldCoords = piece.PieceZCoords[oldState]
		newCoords = piece.PieceZCoords[newState]
		break
	}

	return oldCoords, newCoords
}

// DestroyFullRows destroys full rows
func (board *Board) DestroyFullRows() {
	fall := 0
	for i := range board.squares {
		full := true

		for j := range board.squares[i] {
			if board.squares[i][j] == nil {
				full = false
			}
		}

		if full {
			for j := range board.squares[i] {
				board.squares[i][j] = nil
			}
			fall++
		} else {
			oldI := i

			for j := 0; j < fall; j++ {
				for k := range board.squares[i] {
					block := board.squares[i][k]
					if block != nil {
						block.MoveDown()
						board.squares[i-1][k] = block
						board.squares[i][k] = nil
					}
				}

				i--
			}

			i = oldI
		}
	}
}
