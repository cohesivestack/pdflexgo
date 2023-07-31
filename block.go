package pdflexgo

//go:generate go run generator/main.go

import (
	"math"

	"github.com/kjk/flex"
)

const edgeCount = 4
const edgeTopIndex = 0
const edgeRightIndex = 1
const edgeBottomIndex = 2
const edgeLeftIndex = 3

type Block struct {
	AbstractElement
	children []Element
	border   [edgeCount]*border
}

func NewBlock() *Block {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	block := &Block{
		border: [edgeCount]*border{{}, {}, {}, {}},
	}

	block.AbstractElement.setFlexNode(node)
	block.MarginAll(0)
	block.BorderAllWidth(0)
	block.BorderAllColor("#000000")
	block.FlexDirection(DefaultFlexDirection)
	return block
}

func (block *Block) Children(children ...Element) *Block {

	for _, child := range children {
		block.getFlexNode().InsertChild(child.getFlexNode(), len(block.getFlexNode().Children))
		block.children = append(block.children, child)
	}

	return block
}

func (block *Block) render(pdf *Pdf) {

	block.renderBorders(pdf)

	for _, child := range block.children {
		child.render(pdf)
	}
}

func (block *Block) Width(width float64) *Block {
	block.getFlexNode().StyleSetWidth(float32(width))
	return block
}

func (block *Block) WidthAuto() *Block {
	block.getFlexNode().StyleSetWidthAuto()
	return block
}

func (block *Block) Height(height float64) *Block {
	block.getFlexNode().StyleSetHeight(float32(height))
	return block
}

func (block *Block) Display(display flex.Display) *Block {
	block.getFlexNode().StyleSetDisplay(display)
	return block
}

func (block *Block) GetWidth() float64 {
	return float64(block.getFlexNode().StyleGetWidth().Value)
}

func (block *Block) GetHeight() float64 {
	return float64(block.getFlexNode().StyleGetHeight().Value)
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

func (block *Block) FlexBasisAuto() *Block {
	flex.NodeStyleSetFlexBasisAuto(block.getFlexNode())
	return block
}
