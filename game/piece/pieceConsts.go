package piece

import (
	"github.com/daplf/go-tetris/game/piece/block"
)

const (
	// PieceI is a label for the I piece
	PieceI = "PieceI"

	// PieceJ is a label for the J piece
	PieceJ = "PieceJ"

	// PieceL is a label for the L piece
	PieceL = "PieceL"

	// PieceO is a label for the O piece
	PieceO = "PieceO"

	// PieceS is a label for the S piece
	PieceS = "PieceS"

	// PieceT is a label for the T piece
	PieceT = "PieceT"

	// PieceZ is a label for the Z piece
	PieceZ = "PieceZ"

	// NumPieces is the number of available pieces
	NumPieces = 7
)

var (
	// PieceICoords holds PieceI's coordinates
	PieceICoords = [][][]block.Position{{{-2, -1, 0, 1}, {0, 0, 0, 0}}, {{0, 0, 0, 0}, {1, 0, -1, -2}}, {{1, 0, -1, -2}, {-1, -1, -1, -1}}, {{-1, -1, -1, -1}, {-2, -1, 0, 1}}}

	// PieceJCoords holds PieceJ's coordinates
	PieceJCoords = [][][]block.Position{{{-1, 0, 1, 1}, {0, 0, 0, -1}}, {{0, 0, 0, -1}, {1, 0, -1, -1}}, {{1, 0, -1, -1}, {0, 0, 0, 1}}, {{0, 0, 0, 1}, {-1, 0, 1, 1}}}

	// PieceLCoords holds PieceL's coordinates
	PieceLCoords = [][][]block.Position{{{-1, 0, 1, 1}, {0, 0, 0, 1}}, {{0, 0, 0, 1}, {1, 0, -1, -1}}, {{1, 0, -1, -1}, {0, 0, 0, -1}}, {{0, 0, 0, -1}, {-1, 0, 1, 1}}}

	// PieceOCoords holds PieceO's coordinates
	PieceOCoords = [][][]block.Position{{{-1, -1, 0, 0}, {0, -1, 0, -1}}, {{-1, -1, 0, 0}, {0, -1, 0, -1}}, {{-1, -1, 0, 0}, {0, -1, 0, -1}}, {{-1, -1, 0, 0}, {0, -1, 0, -1}}}

	// PieceSCoords holds PieceS's coordinates
	PieceSCoords = [][][]block.Position{{{-2, -1, -1, 0}, {-1, -1, 0, 0}}, {{-2, -2, -1, -1}, {1, 0, 0, -1}}, {{0, -1, -1, -2}, {1, 1, 0, 0}}, {{0, 0, -1, -1}, {-1, 0, 0, 1}}}

	// PieceTCoords holds PieceT's coordinates
	PieceTCoords = [][][]block.Position{{{-1, 0, 1, 0}, {0, 0, 0, -1}}, {{0, 0, 0, -1}, {1, 0, -1, 0}}, {{1, 0, -1, 0}, {0, 0, 0, 1}}, {{0, 0, 0, 1}, {-1, 0, 1, 0}}}

	// PieceZCoords holds PieceZ's coordinates
	PieceZCoords = [][][]block.Position{{{-2, -1, -1, 0}, {0, 0, -1, -1}}, {{-1, -1, -2, -2}, {1, 0, 0, -1}}, {{0, -1, -1, -2}, {0, 0, 1, 1}}, {{-1, -1, 0, 0}, {-1, 0, 0, 1}}}
)
