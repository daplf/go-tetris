package block

// Type is the block type (the piece it belongs to)
type Type = string

// Position is an alias for the block position
type Position = int

// Block is part of a Piece
type Block struct {
	x         Position
	y         Position
	blockType Type
}

// CreateBlock creates a new block at a specified position
func CreateBlock(x, y Position, blockType Type) *Block {
	return &Block{
		x:         x,
		y:         y,
		blockType: blockType,
	}
}

// X returns the block x position
func (block *Block) X() Position {
	return block.x
}

// Y returns the block y position
func (block *Block) Y() Position {
	return block.y
}

// Type returns the block type
func (block *Block) Type() Type {
	return block.blockType
}

// MoveDown moves the block down
func (block *Block) MoveDown() {
	block.y--
}

// MoveLeft moves the block down
func (block *Block) MoveLeft() {
	block.x--
}

// MoveRight moves the block down
func (block *Block) MoveRight() {
	block.x++
}

// SetX sets a new x for the block
func (block *Block) SetX(newX Position) {
	block.x = newX
}

// SetY sets a new y for the block
func (block *Block) SetY(newY Position) {
	block.y = newY
}
