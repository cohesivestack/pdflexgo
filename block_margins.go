package pdflexgo

import (
	"github.com/kjk/flex"
)

func (block *Block) MarginTopWidth(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(width))

	return block
}

func (block *Block) MarginBottomWidth(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(width))

	return block
}

func (block *Block) MarginLeftWidth(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(width))

	return block
}

func (block *Block) MarginRightWidth(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(width))

	return block
}

func (block *Block) MarginWidth(width float64) *Block {
	block.MarginTopWidth(width)
	block.MarginRightWidth(width)
	block.MarginBottomWidth(width)
	block.MarginLeftWidth(width)

	return block
}
