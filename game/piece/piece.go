package piece

import (
	"github.com/daplf/go-tetris/game/piece/block"
)

const (
	// NormalState is the normal piece state
	NormalState = 0

	// RightState is the right piece state
	RightState = 1

	// InvertedState is the inverted piece state
	InvertedState = 2

	// LeftState is the left piece state
	LeftState = 3

	// NumStates is the number of possible states
	NumStates = 4
)

// State type used by Piece
type State = int

// Piece contains a piece's logic
type Piece struct {
	blocks []*block.Block
	state  State
}

// Blocks returns the piece's blocks
func (piece *Piece) Blocks() []*block.Block {
	return piece.blocks
}

// State returns the piece's state
func (piece *Piece) State() State {
	return piece.state
}

// CreatePiece creates a new piece
func CreatePiece(blocks []*block.Block) *Piece {
	return &Piece{
		blocks: blocks,
		state:  NormalState,
	}
}

// SetState sets a new state for the piece
func (piece *Piece) SetState(newState State) {
	piece.state = newState
}
