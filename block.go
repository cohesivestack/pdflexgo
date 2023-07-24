package pdflexgo

import "github.com/kjk/flex"

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
	flexNode *flex.Node
	border   [edgeCount]*border
	pdf      *Pdf
	parent   *Block
	children []*Block
}

func NewBlock() *Block {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	return &Block{
		flexNode: node,
		border:   [edgeCount]*border{{}, {}, {}, {}},
	}
}

func (block *Block) Children(children ...*Block) *Block {

	for _, child := range children {
		block.flexNode.InsertChild(child.flexNode, len(block.flexNode.Children))
	}

	block.children = append(block.children, children...)

	return block
}

func (block *Block) X() float32 {
	x := block.flexNode.LayoutGetLeft()
	if block.flexNode.Parent != nil {
		x += block.flexNode.Parent.LayoutGetLeft()
	}

	return x
}

func (block *Block) Y() float32 {
	y := block.flexNode.LayoutGetTop()
	if block.flexNode.Parent != nil {
		y += block.flexNode.Parent.LayoutGetTop()
	}

	return y
}

func (block *Block) build() *Block {

	block.buildBorders()

	for _, child := range block.children {
		child.build()
	}

	return block
}

func (block *Block) setPdf(pdf *Pdf) *Block {
	block.pdf = pdf

	for _, block := range block.children {
		block.setPdf(pdf)
	}

	return block
}
