package pdflexgo

import (
	"log"

	"github.com/kjk/flex"
)

type border struct {
	color string
}

func (block *BlockElement) BorderTopWidth(width float64) *BlockElement {

	block.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(width))

	return block
}

func (block *BlockElement) BorderBottomWidth(width float64) *BlockElement {

	block.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(width))

	return block
}

func (block *BlockElement) BorderLeftWidth(width float64) *BlockElement {

	block.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(width))

	return block
}

func (block *BlockElement) BorderRightWidth(width float64) *BlockElement {

	block.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(width))

	return block
}

func (block *BlockElement) BorderWidth(top float64, right float64, bottom float64, left float64) *BlockElement {
	block.BorderTopWidth(top)
	block.BorderRightWidth(right)
	block.BorderBottomWidth(bottom)
	block.BorderLeftWidth(left)

	return block
}

func (block *BlockElement) BorderTopColor(color string) *BlockElement {

	block.border[edgeTopIndex].color = color

	return block
}

func (block *BlockElement) BorderRightColor(color string) *BlockElement {

	block.border[edgeRightIndex].color = color

	return block
}

func (block *BlockElement) BorderBottomColor(color string) *BlockElement {

	block.border[edgeBottomIndex].color = color

	return block
}

func (block *BlockElement) BorderLeftColor(color string) *BlockElement {

	block.border[edgeLeftIndex].color = color

	return block
}

func (block *BlockElement) BorderColor(top string, right string, bottom string, left string) *BlockElement {
	block.BorderTopColor(top)
	block.BorderRightColor(right)
	block.BorderBottomColor(bottom)
	block.BorderLeftColor(left)

	return block
}

func (block *BlockElement) renderBorders(pdf *Pdf) *BlockElement {

	block.renderBorderTop(pdf)
	block.renderBorderRight(pdf)
	block.renderBorderBottom(pdf)
	block.renderBorderLeft(pdf)

	return block
}

func (block *BlockElement) renderBorderTop(pdf *Pdf) *BlockElement {

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

func (block *BlockElement) renderBorderRight(pdf *Pdf) *BlockElement {

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

func (block *BlockElement) renderBorderBottom(pdf *Pdf) *BlockElement {

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

func (block *BlockElement) renderBorderLeft(pdf *Pdf) *BlockElement {

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

func (block *BlockElement) BorderVerticalWidth(vertical float64) *BlockElement {
	block.
		BorderTopWidth(vertical).
		BorderBottomWidth(vertical)

	return block
}

func (block *BlockElement) BorderHorizontalWidth(horizontal float64) *BlockElement {
	block.
		BorderLeftWidth(horizontal).
		BorderRightWidth(horizontal)

	return block
}

func (block *BlockElement) BorderAllWidth(border float64) *BlockElement {
	block.
		BorderTopWidth(border).
		BorderRightWidth(border).
		BorderBottomWidth(border).
		BorderLeftWidth(border)

	return block
}

func (block *BlockElement) BorderVerticalColor(color string) *BlockElement {
	block.
		BorderTopColor(color).
		BorderBottomColor(color)

	return block
}

func (block *BlockElement) BorderHorizontalColor(color string) *BlockElement {
	block.
		BorderLeftColor(color).
		BorderRightColor(color)

	return block
}

func (block *BlockElement) BorderAllColor(color string) *BlockElement {
	block.
		BorderTopColor(color).
		BorderRightColor(color).
		BorderBottomColor(color).
		BorderLeftColor(color)

	return block
}
