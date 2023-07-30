package pdflexgo

import (
	"math"

	"github.com/kjk/flex"
)

func (block *Block) FlexGrow(grow float64) *Block {
	block.getFlexNode().StyleSetFlexGrow(float32(grow))
	return block
}

func (block *Block) FlexShrink(shrink float64) *Block {
	block.getFlexNode().StyleSetFlexShrink(float32(shrink))
	return block
}

func (block *Block) FlexBasis(basis float64) *Block {
	block.getFlexNode().StyleSetFlexBasis(float32(basis))
	return block
}

func (block *Block) FlexAuto() *Block {
	return block.
		FlexGrow(1).
		FlexShrink(1).
		FlexBasis(math.NaN())
}

func (block *Block) FlexNone() *Block {
	return block.
		FlexGrow(0).
		FlexShrink(0).
		FlexBasis(math.NaN())
}

func (block *Block) FlexDirection(direction FlexDirection) *Block {
	block.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return block
}
