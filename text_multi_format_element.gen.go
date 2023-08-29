// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import (
	"github.com/kjk/flex"
)

// MARGIN

func (textMultiFormat *TextMultiFormatElement) MarginTop(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(margin))
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginRight(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(margin))
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginBottom(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(margin))
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginLeft(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(margin))
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginHorizontal(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginLeft(margin).MarginRight(margin)
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginVertical(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginTop(margin).MarginBottom(margin)
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) MarginAll(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginHorizontal(margin).MarginVertical(margin)
	return textMultiFormat
}
func (textMultiFormat *TextMultiFormatElement) Margin(top float64, right float64, bottom float64, left float64) *TextMultiFormatElement {
	textMultiFormat.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)
	return textMultiFormat
}

// MARGIN PERCENT

func (textMultiFormat *TextMultiFormatElement) MarginPercentTop(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMarginPercent(flex.EdgeTop, float32(margin))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentRight(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMarginPercent(flex.EdgeRight, float32(margin))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentBottom(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMarginPercent(flex.EdgeBottom, float32(margin))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentLeft(margin float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMarginPercent(flex.EdgeLeft, float32(margin))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentHorizontal(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginPercentLeft(margin).MarginPercentRight(margin)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentVertical(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginPercentTop(margin).MarginPercentBottom(margin)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MarginPercentAll(margin float64) *TextMultiFormatElement {
	textMultiFormat.MarginPercentHorizontal(margin).MarginPercentVertical(margin)
	return textMultiFormat
}

// PADDING

// For Padding

func (textMultiFormat *TextMultiFormatElement) PaddingTop(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPadding(flex.EdgeTop, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingRight(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPadding(flex.EdgeRight, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingBottom(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPadding(flex.EdgeBottom, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingLeft(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPadding(flex.EdgeLeft, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingHorizontal(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingLeft(padding).PaddingRight(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingVertical(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingTop(padding).PaddingBottom(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingAll(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingHorizontal(padding).PaddingVertical(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) Padding(top float64, right float64, bottom float64, left float64) *TextMultiFormatElement {
	textMultiFormat.
		PaddingTop(top).
		PaddingRight(right).
		PaddingBottom(bottom).
		PaddingLeft(left)
	return textMultiFormat
}

// PADDING PERCENT

func (textMultiFormat *TextMultiFormatElement) PaddingPercentTop(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPaddingPercent(flex.EdgeTop, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentRight(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPaddingPercent(flex.EdgeRight, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentBottom(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPaddingPercent(flex.EdgeBottom, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentLeft(padding float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPaddingPercent(flex.EdgeLeft, float32(padding))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentHorizontal(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingPercentLeft(padding).PaddingPercentRight(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentVertical(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingPercentTop(padding).PaddingPercentBottom(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercentAll(padding float64) *TextMultiFormatElement {
	textMultiFormat.PaddingPercentHorizontal(padding).PaddingPercentVertical(padding)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PaddingPercent(top float64, right float64, bottom float64, left float64) *TextMultiFormatElement {
	textMultiFormat.
		PaddingPercentTop(top).
		PaddingPercentRight(right).
		PaddingPercentBottom(bottom).
		PaddingPercentLeft(left)
	return textMultiFormat
}

// BORDERWIDTH

func (textMultiFormat *TextMultiFormatElement) BorderWidthTop(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(borderWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthRight(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(borderWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthBottom(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(borderWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthLeft(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(borderWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthHorizontal(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.BorderWidthLeft(borderWidth).BorderWidthRight(borderWidth)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthVertical(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.BorderWidthTop(borderWidth).BorderWidthBottom(borderWidth)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidthAll(borderWidth float64) *TextMultiFormatElement {
	textMultiFormat.BorderWidthHorizontal(borderWidth).BorderWidthVertical(borderWidth)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderWidth(top float64, right float64, bottom float64, left float64) *TextMultiFormatElement {
	textMultiFormat.
		BorderWidthTop(top).
		BorderWidthRight(right).
		BorderWidthBottom(bottom).
		BorderWidthLeft(left)
	return textMultiFormat
}

// BACKGROUND

func (textMultiFormat *TextMultiFormatElement) BackgroundColor(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.backgroundColor = getRgba(red, green, blue, alpha...)
	return textMultiFormat
}

// BORDER COLOR

func (textMultiFormat *TextMultiFormatElement) BorderColorTop(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.borderColor[EdgeTop] = getRgba(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorRight(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.borderColor[EdgeRight] = getRgba(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorBottom(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.borderColor[EdgeBottom] = getRgba(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorLeft(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.borderColor[EdgeLeft] = getRgba(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorHorizontal(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.
		BorderColorLeft(red, green, blue, alpha...).
		BorderColorRight(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorVertical(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.
		BorderColorTop(red, green, blue, alpha...).
		BorderColorBottom(red, green, blue, alpha...)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) BorderColorAll(red int, green int, blue int, alpha ...float64) *TextMultiFormatElement {
	textMultiFormat.
		BorderColorHorizontal(red, green, blue, alpha...).
		BorderColorVertical(red, green, blue, alpha...)
	return textMultiFormat
}
