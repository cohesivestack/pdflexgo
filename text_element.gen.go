// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import (
	"github.com/kjk/flex"
)

// MARGIN

func (text *TextElement) MarginTop(margin float64) *TextElement {
	text.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(margin))
	return text
}
func (text *TextElement) MarginRight(margin float64) *TextElement {
	text.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(margin))
	return text
}
func (text *TextElement) MarginBottom(margin float64) *TextElement {
	text.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(margin))
	return text
}
func (text *TextElement) MarginLeft(margin float64) *TextElement {
	text.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(margin))
	return text
}
func (text *TextElement) MarginHorizontal(margin float64) *TextElement {
	text.MarginLeft(margin).MarginRight(margin)
	return text
}
func (text *TextElement) MarginVertical(margin float64) *TextElement {
	text.MarginTop(margin).MarginBottom(margin)
	return text
}
func (text *TextElement) MarginAll(margin float64) *TextElement {
	text.MarginHorizontal(margin).MarginVertical(margin)
	return text
}
func (text *TextElement) Margin(top float64, right float64, bottom float64, left float64) *TextElement {
	text.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)
	return text
}

// MARGIN PERCENT

func (text *TextElement) MarginPercentTop(margin float64) *TextElement {
	text.getFlexNode().StyleSetMarginPercent(flex.EdgeTop, float32(margin))
	return text
}

func (text *TextElement) MarginPercentRight(margin float64) *TextElement {
	text.getFlexNode().StyleSetMarginPercent(flex.EdgeRight, float32(margin))
	return text
}

func (text *TextElement) MarginPercentBottom(margin float64) *TextElement {
	text.getFlexNode().StyleSetMarginPercent(flex.EdgeBottom, float32(margin))
	return text
}

func (text *TextElement) MarginPercentLeft(margin float64) *TextElement {
	text.getFlexNode().StyleSetMarginPercent(flex.EdgeLeft, float32(margin))
	return text
}

func (text *TextElement) MarginPercentHorizontal(margin float64) *TextElement {
	text.MarginPercentLeft(margin).MarginPercentRight(margin)
	return text
}

func (text *TextElement) MarginPercentVertical(margin float64) *TextElement {
	text.MarginPercentTop(margin).MarginPercentBottom(margin)
	return text
}

func (text *TextElement) MarginPercentAll(margin float64) *TextElement {
	text.MarginPercentHorizontal(margin).MarginPercentVertical(margin)
	return text
}

// PADDING

// For Padding

func (text *TextElement) PaddingTop(padding float64) *TextElement {
	text.getFlexNode().StyleSetPadding(flex.EdgeTop, float32(padding))
	return text
}

func (text *TextElement) PaddingRight(padding float64) *TextElement {
	text.getFlexNode().StyleSetPadding(flex.EdgeRight, float32(padding))
	return text
}

func (text *TextElement) PaddingBottom(padding float64) *TextElement {
	text.getFlexNode().StyleSetPadding(flex.EdgeBottom, float32(padding))
	return text
}

func (text *TextElement) PaddingLeft(padding float64) *TextElement {
	text.getFlexNode().StyleSetPadding(flex.EdgeLeft, float32(padding))
	return text
}

func (text *TextElement) PaddingHorizontal(padding float64) *TextElement {
	text.PaddingLeft(padding).PaddingRight(padding)
	return text
}

func (text *TextElement) PaddingVertical(padding float64) *TextElement {
	text.PaddingTop(padding).PaddingBottom(padding)
	return text
}

func (text *TextElement) PaddingAll(padding float64) *TextElement {
	text.PaddingHorizontal(padding).PaddingVertical(padding)
	return text
}

func (text *TextElement) Padding(top float64, right float64, bottom float64, left float64) *TextElement {
	text.
		PaddingTop(top).
		PaddingRight(right).
		PaddingBottom(bottom).
		PaddingLeft(left)
	return text
}

// PADDING PERCENT

func (text *TextElement) PaddingPercentTop(padding float64) *TextElement {
	text.getFlexNode().StyleSetPaddingPercent(flex.EdgeTop, float32(padding))
	return text
}

func (text *TextElement) PaddingPercentRight(padding float64) *TextElement {
	text.getFlexNode().StyleSetPaddingPercent(flex.EdgeRight, float32(padding))
	return text
}

func (text *TextElement) PaddingPercentBottom(padding float64) *TextElement {
	text.getFlexNode().StyleSetPaddingPercent(flex.EdgeBottom, float32(padding))
	return text
}

func (text *TextElement) PaddingPercentLeft(padding float64) *TextElement {
	text.getFlexNode().StyleSetPaddingPercent(flex.EdgeLeft, float32(padding))
	return text
}

func (text *TextElement) PaddingPercentHorizontal(padding float64) *TextElement {
	text.PaddingPercentLeft(padding).PaddingPercentRight(padding)
	return text
}

func (text *TextElement) PaddingPercentVertical(padding float64) *TextElement {
	text.PaddingPercentTop(padding).PaddingPercentBottom(padding)
	return text
}

func (text *TextElement) PaddingPercentAll(padding float64) *TextElement {
	text.PaddingPercentHorizontal(padding).PaddingPercentVertical(padding)
	return text
}

func (text *TextElement) PaddingPercent(top float64, right float64, bottom float64, left float64) *TextElement {
	text.
		PaddingPercentTop(top).
		PaddingPercentRight(right).
		PaddingPercentBottom(bottom).
		PaddingPercentLeft(left)
	return text
}

// BORDERWIDTH

func (text *TextElement) BorderWidthTop(borderWidth float64) *TextElement {
	text.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(borderWidth))
	return text
}

func (text *TextElement) BorderWidthRight(borderWidth float64) *TextElement {
	text.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(borderWidth))
	return text
}

func (text *TextElement) BorderWidthBottom(borderWidth float64) *TextElement {
	text.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(borderWidth))
	return text
}

func (text *TextElement) BorderWidthLeft(borderWidth float64) *TextElement {
	text.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(borderWidth))
	return text
}

func (text *TextElement) BorderWidthHorizontal(borderWidth float64) *TextElement {
	text.BorderWidthLeft(borderWidth).BorderWidthRight(borderWidth)
	return text
}

func (text *TextElement) BorderWidthVertical(borderWidth float64) *TextElement {
	text.BorderWidthTop(borderWidth).BorderWidthBottom(borderWidth)
	return text
}

func (text *TextElement) BorderWidthAll(borderWidth float64) *TextElement {
	text.BorderWidthHorizontal(borderWidth).BorderWidthVertical(borderWidth)
	return text
}

func (text *TextElement) BorderWidth(top float64, right float64, bottom float64, left float64) *TextElement {
	text.
		BorderWidthTop(top).
		BorderWidthRight(right).
		BorderWidthBottom(bottom).
		BorderWidthLeft(left)
	return text
}

// BACKGROUND

func (text *TextElement) BackgroundColor(red int, green int, blue int, alpha ...float64) *TextElement {
	text.backgroundColor = getRgba(red, green, blue, alpha...)
	return text
}

// BORDER COLOR

func (text *TextElement) BorderColorTop(red int, green int, blue int, alpha ...float64) *TextElement {
	text.borderColor[EdgeTop] = getRgba(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorRight(red int, green int, blue int, alpha ...float64) *TextElement {
	text.borderColor[EdgeRight] = getRgba(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorBottom(red int, green int, blue int, alpha ...float64) *TextElement {
	text.borderColor[EdgeBottom] = getRgba(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorLeft(red int, green int, blue int, alpha ...float64) *TextElement {
	text.borderColor[EdgeLeft] = getRgba(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorHorizontal(red int, green int, blue int, alpha ...float64) *TextElement {
	text.
		BorderColorLeft(red, green, blue, alpha...).
		BorderColorRight(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorVertical(red int, green int, blue int, alpha ...float64) *TextElement {
	text.
		BorderColorTop(red, green, blue, alpha...).
		BorderColorBottom(red, green, blue, alpha...)
	return text
}

func (text *TextElement) BorderColorAll(red int, green int, blue int, alpha ...float64) *TextElement {
	text.
		BorderColorHorizontal(red, green, blue, alpha...).
		BorderColorVertical(red, green, blue, alpha...)
	return text
}
