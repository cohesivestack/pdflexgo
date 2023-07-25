package pdflexgo

import (
	"log"

	"github.com/kjk/flex"
)

type border struct {
	color string
}

func (block *Block) BorderTopWidth(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(width))

	return block
}

func (block *Block) BorderBottomWidth(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(width))

	return block
}

func (block *Block) BorderLeftWidth(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(width))

	return block
}

func (block *Block) BorderRightWidth(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(width))

	return block
}

func (block *Block) BorderWidth(width float64) *Block {
	block.BorderTopWidth(width)
	block.BorderRightWidth(width)
	block.BorderBottomWidth(width)
	block.BorderLeftWidth(width)

	return block
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
	block.BorderTopColor(color)
	block.BorderRightColor(color)
	block.BorderBottomColor(color)
	block.BorderLeftColor(color)

	return block
}

func (block *Block) renderBorders(pdf *Pdf) *Block {

	block.renderBorderTop(pdf)
	block.renderBorderRight(pdf)
	block.renderBorderBottom(pdf)
	block.renderBorderLeft(pdf)

	return block
}

func (block *Block) renderBorderTop(pdf *Pdf) *Block {

	if block.border[edgeTopIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeTopIndex].color)
		if err != nil {
			log.Fatal(err)
		}
		pdf.fpdf.SetDrawColor(r, g, b)
	}

	pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeTop)))
	pdf.fpdf.Line(
		float64(block.X()),
		float64(block.Y()),
		float64(block.getFlexNode().LayoutGetWidth()+block.X()),
		float64(block.Y()))

	return block
}

func (block *Block) renderBorderRight(pdf *Pdf) *Block {

	if block.border[edgeRightIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeRightIndex].color)
		if err != nil {
			log.Fatal(err)
		}

		pdf.fpdf.SetDrawColor(r, g, b)
		pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeRight)))
		pdf.fpdf.Line(
			float64(block.X()+block.getFlexNode().LayoutGetWidth()),
			float64(block.Y()),
			float64(block.X()+block.getFlexNode().LayoutGetWidth()),
			float64(block.X()+block.getFlexNode().LayoutGetHeight()))
	}

	return block
}

func (block *Block) renderBorderBottom(pdf *Pdf) *Block {

	if block.border[edgeBottomIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeBottomIndex].color)
		if err != nil {
			log.Fatal(err)
		}
		pdf.fpdf.SetDrawColor(r, g, b)
	}

	pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeBottom)))
	pdf.fpdf.Line(
		float64(block.X()),
		float64(block.Y()+block.getFlexNode().LayoutGetHeight()),
		float64(block.X()+block.getFlexNode().LayoutGetWidth()),
		float64(block.Y()+block.getFlexNode().LayoutGetHeight()))

	return block
}

func (block *Block) renderBorderLeft(pdf *Pdf) *Block {

	if block.border[edgeLeftIndex].color != "" {
		r, g, b, err := hexToRGB(block.border[edgeLeftIndex].color)
		if err != nil {
			log.Fatal(err)
		}

		pdf.fpdf.SetDrawColor(r, g, b)
		pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeLeft)))
		pdf.fpdf.Line(
			float64(block.X()),
			float64(block.Y()+block.getFlexNode().LayoutGetBottom()),
			float64(block.X()),
			float64(block.X()+block.getFlexNode().LayoutGetHeight()))
	}

	return block
}
