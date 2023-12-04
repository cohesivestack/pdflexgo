// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import (
	"github.com/kjk/flex"
)

// MARGIN

func (block *BlockElement) MarginTop(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(margin))
	return block
}
func (block *BlockElement) MarginRight(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(margin))
	return block
}
func (block *BlockElement) MarginBottom(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(margin))
	return block
}
func (block *BlockElement) MarginLeft(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(margin))
	return block
}
func (block *BlockElement) MarginHorizontal(margin float64) *BlockElement {
	block.MarginLeft(margin).MarginRight(margin)
	return block
}
func (block *BlockElement) MarginVertical(margin float64) *BlockElement {
	block.MarginTop(margin).MarginBottom(margin)
	return block
}
func (block *BlockElement) MarginAll(margin float64) *BlockElement {
	block.MarginHorizontal(margin).MarginVertical(margin)
	return block
}
func (block *BlockElement) Margin(top float64, right float64, bottom float64, left float64) *BlockElement {
	block.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)
	return block
}

// MARGIN PERCENT

func (block *BlockElement) MarginPercentTop(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMarginPercent(flex.EdgeTop, float32(margin))
	return block
}

func (block *BlockElement) MarginPercentRight(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMarginPercent(flex.EdgeRight, float32(margin))
	return block
}

func (block *BlockElement) MarginPercentBottom(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMarginPercent(flex.EdgeBottom, float32(margin))
	return block
}

func (block *BlockElement) MarginPercentLeft(margin float64) *BlockElement {
	block.getFlexNode().StyleSetMarginPercent(flex.EdgeLeft, float32(margin))
	return block
}

func (block *BlockElement) MarginPercentHorizontal(margin float64) *BlockElement {
	block.MarginPercentLeft(margin).MarginPercentRight(margin)
	return block
}

func (block *BlockElement) MarginPercentVertical(margin float64) *BlockElement {
	block.MarginPercentTop(margin).MarginPercentBottom(margin)
	return block
}

func (block *BlockElement) MarginPercentAll(margin float64) *BlockElement {
	block.MarginPercentHorizontal(margin).MarginPercentVertical(margin)
	return block
}

// PADDING

// For Padding

func (block *BlockElement) PaddingTop(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPadding(flex.EdgeTop, float32(padding))
	return block
}

func (block *BlockElement) PaddingRight(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPadding(flex.EdgeRight, float32(padding))
	return block
}

func (block *BlockElement) PaddingBottom(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPadding(flex.EdgeBottom, float32(padding))
	return block
}

func (block *BlockElement) PaddingLeft(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPadding(flex.EdgeLeft, float32(padding))
	return block
}

func (block *BlockElement) PaddingHorizontal(padding float64) *BlockElement {
	block.PaddingLeft(padding).PaddingRight(padding)
	return block
}

func (block *BlockElement) PaddingVertical(padding float64) *BlockElement {
	block.PaddingTop(padding).PaddingBottom(padding)
	return block
}

func (block *BlockElement) PaddingAll(padding float64) *BlockElement {
	block.PaddingHorizontal(padding).PaddingVertical(padding)
	return block
}

func (block *BlockElement) Padding(top float64, right float64, bottom float64, left float64) *BlockElement {
	block.
		PaddingTop(top).
		PaddingRight(right).
		PaddingBottom(bottom).
		PaddingLeft(left)
	return block
}

// PADDING PERCENT

func (block *BlockElement) PaddingPercentTop(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPaddingPercent(flex.EdgeTop, float32(padding))
	return block
}

func (block *BlockElement) PaddingPercentRight(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPaddingPercent(flex.EdgeRight, float32(padding))
	return block
}

func (block *BlockElement) PaddingPercentBottom(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPaddingPercent(flex.EdgeBottom, float32(padding))
	return block
}

func (block *BlockElement) PaddingPercentLeft(padding float64) *BlockElement {
	block.getFlexNode().StyleSetPaddingPercent(flex.EdgeLeft, float32(padding))
	return block
}

func (block *BlockElement) PaddingPercentHorizontal(padding float64) *BlockElement {
	block.PaddingPercentLeft(padding).PaddingPercentRight(padding)
	return block
}

func (block *BlockElement) PaddingPercentVertical(padding float64) *BlockElement {
	block.PaddingPercentTop(padding).PaddingPercentBottom(padding)
	return block
}

func (block *BlockElement) PaddingPercentAll(padding float64) *BlockElement {
	block.PaddingPercentHorizontal(padding).PaddingPercentVertical(padding)
	return block
}

func (block *BlockElement) PaddingPercent(top float64, right float64, bottom float64, left float64) *BlockElement {
	block.
		PaddingPercentTop(top).
		PaddingPercentRight(right).
		PaddingPercentBottom(bottom).
		PaddingPercentLeft(left)
	return block
}

// BORDERWIDTH

func (block *BlockElement) BorderWidthTop(borderWidth float64) *BlockElement {
	block.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(borderWidth))
	return block
}

func (block *BlockElement) BorderWidthRight(borderWidth float64) *BlockElement {
	block.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(borderWidth))
	return block
}

func (block *BlockElement) BorderWidthBottom(borderWidth float64) *BlockElement {
	block.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(borderWidth))
	return block
}

func (block *BlockElement) BorderWidthLeft(borderWidth float64) *BlockElement {
	block.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(borderWidth))
	return block
}

func (block *BlockElement) BorderWidthHorizontal(borderWidth float64) *BlockElement {
	block.BorderWidthLeft(borderWidth).BorderWidthRight(borderWidth)
	return block
}

func (block *BlockElement) BorderWidthVertical(borderWidth float64) *BlockElement {
	block.BorderWidthTop(borderWidth).BorderWidthBottom(borderWidth)
	return block
}

func (block *BlockElement) BorderWidthAll(borderWidth float64) *BlockElement {
	block.BorderWidthHorizontal(borderWidth).BorderWidthVertical(borderWidth)
	return block
}

func (block *BlockElement) BorderWidth(top float64, right float64, bottom float64, left float64) *BlockElement {
	block.
		BorderWidthTop(top).
		BorderWidthRight(right).
		BorderWidthBottom(bottom).
		BorderWidthLeft(left)
	return block
}

// BACKGROUND

func (block *BlockElement) BackgroundColor(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.backgroundColor = getRgba(red, green, blue, alpha...)
	return block
}

// BORDER COLOR

func (block *BlockElement) BorderColorTop(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.borderColor[EdgeTop] = getRgba(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorRight(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.borderColor[EdgeRight] = getRgba(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorBottom(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.borderColor[EdgeBottom] = getRgba(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorLeft(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.borderColor[EdgeLeft] = getRgba(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorHorizontal(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.
		BorderColorLeft(red, green, blue, alpha...).
		BorderColorRight(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorVertical(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.
		BorderColorTop(red, green, blue, alpha...).
		BorderColorBottom(red, green, blue, alpha...)
	return block
}

func (block *BlockElement) BorderColorAll(red int, green int, blue int, alpha ...float64) *BlockElement {
	block.
		BorderColorHorizontal(red, green, blue, alpha...).
		BorderColorVertical(red, green, blue, alpha...)
	return block
}
