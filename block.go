package pdflexgo

//go:generate go run generator/main.go

import (
	"github.com/jung-kurt/gofpdf"
)

const edgeCount = 4

type BlockElement struct {
	abstractElement

	children []Node
}

func Block() *BlockElement {

	block := &BlockElement{}

	block.abstractElement.initialize()

	return block
}

func (block *BlockElement) markRequiredAsDirty() {
	for _, child := range block.children {
		child.markRequiredAsDirty()
	}
}

func (block *BlockElement) render(pdf *Pdf) {
	block.renderElement(pdf)

	for _, child := range block.children {
		child.render(pdf)
	}
}

func (block *BlockElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	for _, child := range block.children {
		child.preRender(defaultProps, fpdf)
	}
}
