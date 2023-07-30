package pdflexgo

import (
	"log"

	"github.com/kjk/flex"
)

type border struct {
	color string
}

func (block *Block) BorderTop(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(width))

	return block
}

func (block *Block) BorderBottom(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(width))

	return block
}

func (block *Block) BorderLeft(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(width))

	return block
}

func (block *Block) BorderRight(width float64) *Block {

	block.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(width))

	return block
}

func (block *Block) Border(top float64, right float64, bottom float64, left float64) *Block {
	block.BorderTop(top)
	block.BorderRight(right)
	block.BorderBottom(bottom)
	block.BorderLeft(left)

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

func (block *Block) BorderColor(top string, right string, bottom string, left string) *Block {
	block.BorderTopColor(top)
	block.BorderRightColor(right)
	block.BorderBottomColor(bottom)
	block.BorderLeftColor(left)

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

	if block.getFlexNode().LayoutGetBorder(flex.EdgeTop) > 0 {
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
			float64(block.X()+block.getFlexNode().LayoutGetWidth()),
			float64(block.Y()))
	}

	return block
}

func (block *Block) renderBorderRight(pdf *Pdf) *Block {

	if block.getFlexNode().LayoutGetBorder(flex.EdgeRight) > 0 {
		if block.border[edgeRightIndex].color != "" {
			r, g, b, err := hexToRGB(block.border[edgeRightIndex].color)
			if err != nil {
				log.Fatal(err)
			}

			pdf.fpdf.SetDrawColor(r, g, b)
		}
		pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeRight)))

		pdf.fpdf.Line(
			float64(block.X()+block.getFlexNode().LayoutGetWidth()),
			float64(block.Y()),
			float64(block.X()+block.getFlexNode().LayoutGetWidth()),
			float64(block.Y()+block.getFlexNode().LayoutGetHeight()))
	}

	return block
}

func (block *Block) renderBorderBottom(pdf *Pdf) *Block {

	if block.getFlexNode().LayoutGetBorder(flex.EdgeBottom) > 0 {
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
	}

	return block
}

func (block *Block) renderBorderLeft(pdf *Pdf) *Block {

	if block.getFlexNode().LayoutGetBorder(flex.EdgeLeft) > 0 {
		if block.border[edgeLeftIndex].color != "" {
			r, g, b, err := hexToRGB(block.border[edgeLeftIndex].color)
			if err != nil {
				log.Fatal(err)
			}

			pdf.fpdf.SetDrawColor(r, g, b)
		}

		pdf.fpdf.SetLineWidth(float64(block.getFlexNode().LayoutGetBorder(flex.EdgeLeft)))
		pdf.fpdf.Line(
			float64(block.X()),
			float64(block.Y()),
			float64(block.X()),
			float64(block.Y()+block.getFlexNode().LayoutGetHeight()))
	}

	return block
}

func (block *Block) BorderVertical(vertical float64) *Block {
	block.
		BorderTop(vertical).
		BorderBottom(vertical)

	return block
}

func (block *Block) BorderHorizontal(horizontal float64) *Block {
	block.
		BorderLeft(horizontal).
		BorderRight(horizontal)

	return block
}

func (block *Block) BorderAll(border float64) *Block {
	block.
		BorderTop(border).
		BorderRight(border).
		BorderBottom(border).
		BorderLeft(border)

	return block
}

func (block *Block) BorderVerticalColor(color string) *Block {
	block.
		BorderTopColor(color).
		BorderBottomColor(color)

	return block
}

func (block *Block) BorderHorizontalColor(color string) *Block {
	block.
		BorderLeftColor(color).
		BorderRightColor(color)

	return block
}

func (block *Block) BorderAllColor(color string) *Block {
	block.
		BorderTopColor(color).
		BorderRightColor(color).
		BorderBottomColor(color).
		BorderLeftColor(color)

	return block
}
