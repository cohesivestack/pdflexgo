package pdflexgo

import (
	"github.com/kjk/flex"
)

func (block *Block) MarginTop(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(width))

	return block
}

func (block *Block) MarginBottom(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(width))

	return block
}

func (block *Block) MarginLeft(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(width))

	return block
}

func (block *Block) MarginRight(width float64) *Block {

	block.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(width))

	return block
}

func (block *Block) Margin(width float64) *Block {
	block.MarginTop(width)
	block.MarginRight(width)
	block.MarginBottom(width)
	block.MarginLeft(width)

	return block
}
