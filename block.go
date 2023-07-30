package pdflexgo

import (
	"math"

	"github.com/kjk/flex"
)

type Align string
type Justify string

const edgeCount = 4
const edgeTopIndex = 0
const edgeRightIndex = 1
const edgeBottomIndex = 2
const edgeLeftIndex = 3

const (
	AlignStart  Align = "start"
	AlignEnd    Align = "end"
	AlignCenter Align = "center"
)

const (
	JustifyStart  Justify = "start"
	JustifyEnd    Justify = "end"
	JustifyCenter Justify = "center"
)

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
	block.BorderAll(0)
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

func (block *Block) Flex(flext float64) *Block {
	block.getFlexNode().StyleSetFlex(float32(math.NaN()))
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
