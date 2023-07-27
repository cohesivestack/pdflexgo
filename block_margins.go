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

func (block *Block) GetMarginTop() float64 {

	return float64(block.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
}

func (block *Block) GetMarginBottom() float64 {

	return float64(block.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
}

func (block *Block) GetMarginLeft() float64 {

	return float64(block.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)
}

func (block *Block) GetMarginRight() float64 {

	return float64(block.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
}

func (block *Block) Margin(top float64, right float64, bottom float64, left float64) *Block {
	block.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)

	return block
}

func (block *Block) MarginVertical(vertical float64) *Block {
	block.
		MarginTop(vertical).
		MarginBottom(vertical)

	return block
}

func (block *Block) MarginHorizontal(horizontal float64) *Block {
	block.
		MarginLeft(horizontal).
		MarginRight(horizontal)

	return block
}

func (block *Block) MarginAll(margin float64) *Block {
	block.
		MarginTop(margin).
		MarginRight(margin).
		MarginBottom(margin).
		MarginLeft(margin)

	return block
}
