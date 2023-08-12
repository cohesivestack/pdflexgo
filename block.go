package pdflexgo

//go:generate go run generator/main.go

import (
	"log"

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
	children       []Element
	border         [edgeCount]*border
	backgrondColor string
}

func Block() *BlockElement {
	config := flex.NewConfig()
	config.UseWebDefaults = true
	node := flex.NewNodeWithConfig(config)

	block := &BlockElement{
		border: [edgeCount]*border{{}, {}, {}, {}},
	}

	//node.StyleSetFlexShrink(1)

	block.AbstractElement.setFlexNode(node)
	block.MarginAll(0)
	block.PaddingAll(0)
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

func (block *BlockElement) markDirty() {
	block.AbstractElement.markDirty()
	for _, child := range block.children {
		child.markDirty()
	}
}

func (block *BlockElement) render(pdf *Pdf) {

	block.renderBorders(pdf)

	if block.backgrondColor != "" {
		r, g, b, err := hexToRGB(block.backgrondColor)
		if err != nil {
			log.Fatal(err)
		}
		pdf.fpdf.SetFillColor(r, g, b)
		pdf.fpdf.Rect(
			float64(block.X()+block.getFlexNode().LayoutGetBorder(flex.EdgeLeft)),
			float64(block.Y()+block.getFlexNode().LayoutGetBorder(flex.EdgeTop)),
			float64(block.getFlexNode().LayoutGetWidth()-(block.getFlexNode().LayoutGetBorder(flex.EdgeLeft)+block.getFlexNode().LayoutGetBorder(flex.EdgeRight))),
			float64(block.getFlexNode().LayoutGetHeight()-(block.getFlexNode().LayoutGetBorder(flex.EdgeTop)+block.getFlexNode().LayoutGetBorder(flex.EdgeBottom))), "F")

	}

	for _, child := range block.children {
		child.render(pdf)
	}
}

func (block *BlockElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	for _, child := range block.children {
		child.preRender(defaultProps, fpdf)
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
		FlexBasisAuto()
}

func (block *BlockElement) FlexNone() *BlockElement {
	return block.
		FlexGrow(0).
		FlexShrink(0).
		FlexBasisAuto()
}

func (block *BlockElement) FlexBasisAuto() *BlockElement {
	flex.NodeStyleSetFlexBasisAuto(block.getFlexNode())
	return block
}

func (block *BlockElement) BackgroundColor(color string) *BlockElement {
	block.backgrondColor = color

	return block
}
