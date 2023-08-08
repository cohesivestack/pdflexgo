package pdflexgo

//go:generate go run generator/main.go

import (
	"math"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

const edgeCount = 4
const edgeTopIndex = 0
const edgeRightIndex = 1
const edgeBottomIndex = 2
const edgeLeftIndex = 3

type BlockElement struct {
	AbstractElement
	children []Element
	border   [edgeCount]*border
}

func Block() *BlockElement {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	block := &BlockElement{
		border: [edgeCount]*border{{}, {}, {}, {}},
	}

	block.AbstractElement.setFlexNode(node)
	block.MarginAll(0)
	block.BorderAllWidth(0)
	block.BorderAllColor("#000000")
	block.FlexDirection(DefaultFlexDirection)

	return block
}

func (block *BlockElement) Children(children ...Element) *BlockElement {

	for _, child := range children {
		block.getFlexNode().InsertChild(child.getFlexNode(), len(block.getFlexNode().Children))
		block.children = append(block.children, child)
	}

	return block
}

func (block *BlockElement) render(pdf *Pdf) {

	block.renderBorders(pdf)

	for _, child := range block.children {
		child.render(pdf)
	}
}

func (block *BlockElement) setPreRenderFpdf(fpdf *gofpdf.Fpdf) {

	block.AbstractElement.setPreRenderFpdf(fpdf)

	for _, child := range block.children {
		child.setPreRenderFpdf(fpdf)
	}
}

func (block *BlockElement) WidthAuto() *BlockElement {
	block.getFlexNode().StyleSetWidthAuto()
	return block
}

func (block *BlockElement) HeightAuto() *BlockElement {
	block.getFlexNode().StyleSetHeightAuto()
	return block
}

func (block *BlockElement) GetWidth() float64 {
	return float64(block.getFlexNode().StyleGetWidth().Value)
}

func (block *BlockElement) GetHeight() float64 {
	return float64(block.getFlexNode().StyleGetHeight().Value)
}

func (block *BlockElement) FlexAuto() *BlockElement {
	return block.
		FlexGrow(1).
		FlexShrink(1).
		FlexBasis(math.NaN())
}

func (block *BlockElement) FlexNone() *BlockElement {
	return block.
		FlexGrow(0).
		FlexShrink(0).
		FlexBasis(math.NaN())
}

func (block *BlockElement) FlexBasisAuto() *BlockElement {
	flex.NodeStyleSetFlexBasisAuto(block.getFlexNode())
	return block
}
