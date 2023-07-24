package pdflexgo

import (
	"log"

	"github.com/kjk/flex"
)

type border struct {
	color string
}

func (block *Block) BorderTopWidth(width float32) *Block {

	block.flexNode.StyleSetBorder(flex.EdgeTop, width)

	return block
}

func (block *Block) BorderBottomWidth(width float32) *Block {

	block.flexNode.StyleSetBorder(flex.EdgeBottom, width)

	return block
}

func (block *Block) BorderLeftWidth(width float32) *Block {

	block.flexNode.StyleSetBorder(flex.EdgeLeft, width)

	return block
}

func (block *Block) BorderRightWidth(width float32) *Block {

	block.flexNode.StyleSetBorder(flex.EdgeRight, width)

	return block
}

func (block *Block) BorderWidth(width float32) *Block {
	return block.
		BorderTopWidth(width).
		BorderRightWidth(width).
		BorderBottomWidth(width).
		BorderLeftWidth(width)
}

func (block *Block) BorderTopColor(color string) *Block {

	block.border[edgeTopIndex].color = color

	return block
}

func (block *Block) BorderRightColor(color string) *Block {

	block.border[edgeRightIndex].color = color

	return block
}

func (block *Block) BorderBottomColor(color string) *Block {

	block.border[edgeBottomIndex].color = color

	return block
}

func (block *Block) BorderLeftColor(color string) *Block {

	block.border[edgeLeftIndex].color = color

	return block
}

func (block *Block) BorderColor(color string) *Block {
	return block.
		BorderTopColor(color).
		BorderRightColor(color).
		BorderBottomColor(color).
		BorderLeftColor(color)
}

func (block *Block) buildBorders() *Block {

	block.
		buildBorderTop().
		buildBorderRight().
		buildBorderBottom().
		buildBorderLeft()

	return block
}

func (block *Block) buildBorderTop() *Block {

	if block.border[edgeTopIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeTopIndex].color)
		if err != nil {
			log.Fatal(err)
		}
		block.pdf.fpdf.SetDrawColor(r, g, b)
	}

	block.pdf.fpdf.SetLineWidth(float64(block.flexNode.LayoutGetBorder(flex.EdgeTop)))
	block.pdf.fpdf.Line(
		float64(block.X()),
		float64(block.Y()),
		float64(block.flexNode.LayoutGetWidth()+block.X()),
		float64(block.Y()))

	return block
}

func (block *Block) buildBorderRight() *Block {

	if block.border[edgeRightIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeRightIndex].color)
		if err != nil {
			log.Fatal(err)
		}

		block.pdf.fpdf.SetDrawColor(r, g, b)
		block.pdf.fpdf.SetLineWidth(float64(block.flexNode.LayoutGetBorder(flex.EdgeRight)))
		block.pdf.fpdf.Line(
			float64(block.X()+block.flexNode.LayoutGetWidth()),
			float64(block.Y()),
			float64(block.X()+block.flexNode.LayoutGetWidth()),
			float64(block.X()+block.flexNode.LayoutGetHeight()))
	}

	return block
}

func (block *Block) buildBorderBottom() *Block {

	if block.border[edgeBottomIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeBottomIndex].color)
		if err != nil {
			log.Fatal(err)
		}
		block.pdf.fpdf.SetDrawColor(r, g, b)
	}

	block.pdf.fpdf.SetLineWidth(float64(block.flexNode.LayoutGetBorder(flex.EdgeBottom)))
	block.pdf.fpdf.Line(
		float64(block.X()),
		float64(block.Y()+block.flexNode.LayoutGetHeight()),
		float64(block.X()+block.flexNode.LayoutGetWidth()),
		float64(block.Y()+block.flexNode.LayoutGetHeight()))

	return block
}

func (block *Block) buildBorderLeft() *Block {

	if block.border[edgeLeftIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeLeftIndex].color)
		if err != nil {
			log.Fatal(err)
		}

		block.pdf.fpdf.SetDrawColor(r, g, b)
		block.pdf.fpdf.SetLineWidth(float64(block.flexNode.LayoutGetBorder(flex.EdgeLeft)))
		block.pdf.fpdf.Line(
			float64(block.X()),
			float64(block.Y()+block.flexNode.LayoutGetBottom()),
			float64(block.X()),
			float64(block.X()+block.flexNode.LayoutGetHeight()))
	}

	return block
}
