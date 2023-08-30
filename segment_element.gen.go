// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import (
	"github.com/kjk/flex"
)

// MARGIN

func (segment *Segment) MarginTop(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(margin))
	return segment
}
func (segment *Segment) MarginRight(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(margin))
	return segment
}
func (segment *Segment) MarginBottom(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(margin))
	return segment
}
func (segment *Segment) MarginLeft(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(margin))
	return segment
}
func (segment *Segment) MarginHorizontal(margin float64) *Segment {
	segment.MarginLeft(margin).MarginRight(margin)
	return segment
}
func (segment *Segment) MarginVertical(margin float64) *Segment {
	segment.delegated.MarginTop(margin).MarginBottom(margin)
	return segment
}
func (segment *Segment) MarginAll(margin float64) *Segment {
	segment.delegated.MarginHorizontal(margin).MarginVertical(margin)
	return segment
}
func (segment *Segment) Margin(top float64, right float64, bottom float64, left float64) *Segment {
	segment.delegated.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)
	return segment
}

// MARGIN PERCENT

func (segment *Segment) MarginPercentTop(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMarginPercent(flex.EdgeTop, float32(margin))
	return segment
}

func (segment *Segment) MarginPercentRight(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMarginPercent(flex.EdgeRight, float32(margin))
	return segment
}

func (segment *Segment) MarginPercentBottom(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMarginPercent(flex.EdgeBottom, float32(margin))
	return segment
}

func (segment *Segment) MarginPercentLeft(margin float64) *Segment {
	segment.delegated.getFlexNode().StyleSetMarginPercent(flex.EdgeLeft, float32(margin))
	return segment
}

func (segment *Segment) MarginPercentHorizontal(margin float64) *Segment {
	segment.delegated.MarginPercentLeft(margin).MarginPercentRight(margin)
	return segment
}

func (segment *Segment) MarginPercentVertical(margin float64) *Segment {
	segment.delegated.MarginPercentTop(margin).MarginPercentBottom(margin)
	return segment
}

func (segment *Segment) MarginPercentAll(margin float64) *Segment {
	segment.delegated.MarginPercentHorizontal(margin).MarginPercentVertical(margin)
	return segment
}

// PADDING

// For Padding

func (segment *Segment) PaddingTop(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPadding(flex.EdgeTop, float32(padding))
	return segment
}

func (segment *Segment) PaddingRight(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPadding(flex.EdgeRight, float32(padding))
	return segment
}

func (segment *Segment) PaddingBottom(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPadding(flex.EdgeBottom, float32(padding))
	return segment
}

func (segment *Segment) PaddingLeft(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPadding(flex.EdgeLeft, float32(padding))
	return segment
}

func (segment *Segment) PaddingHorizontal(padding float64) *Segment {
	segment.delegated.PaddingLeft(padding).PaddingRight(padding)
	return segment
}

func (segment *Segment) PaddingVertical(padding float64) *Segment {
	segment.delegated.PaddingTop(padding).PaddingBottom(padding)
	return segment
}

func (segment *Segment) PaddingAll(padding float64) *Segment {
	segment.delegated.PaddingHorizontal(padding).PaddingVertical(padding)
	return segment
}

func (segment *Segment) Padding(top float64, right float64, bottom float64, left float64) *Segment {
	segment.delegated.
		PaddingTop(top).
		PaddingRight(right).
		PaddingBottom(bottom).
		PaddingLeft(left)
	return segment
}

// PADDING PERCENT

func (segment *Segment) PaddingPercentTop(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPaddingPercent(flex.EdgeTop, float32(padding))
	return segment
}

func (segment *Segment) PaddingPercentRight(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPaddingPercent(flex.EdgeRight, float32(padding))
	return segment
}

func (segment *Segment) PaddingPercentBottom(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPaddingPercent(flex.EdgeBottom, float32(padding))
	return segment
}

func (segment *Segment) PaddingPercentLeft(padding float64) *Segment {
	segment.delegated.getFlexNode().StyleSetPaddingPercent(flex.EdgeLeft, float32(padding))
	return segment
}

func (segment *Segment) PaddingPercentHorizontal(padding float64) *Segment {
	segment.delegated.PaddingPercentLeft(padding).PaddingPercentRight(padding)
	return segment
}

func (segment *Segment) PaddingPercentVertical(padding float64) *Segment {
	segment.delegated.PaddingPercentTop(padding).PaddingPercentBottom(padding)
	return segment
}

func (segment *Segment) PaddingPercentAll(padding float64) *Segment {
	segment.delegated.PaddingPercentHorizontal(padding).PaddingPercentVertical(padding)
	return segment
}

func (segment *Segment) PaddingPercent(top float64, right float64, bottom float64, left float64) *Segment {
	segment.delegated.
		PaddingPercentTop(top).
		PaddingPercentRight(right).
		PaddingPercentBottom(bottom).
		PaddingPercentLeft(left)
	return segment
}

// BORDERWIDTH

func (segment *Segment) BorderWidthTop(borderWidth float64) *Segment {
	segment.delegated.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(borderWidth))
	return segment
}

func (segment *Segment) BorderWidthRight(borderWidth float64) *Segment {
	segment.delegated.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(borderWidth))
	return segment
}

func (segment *Segment) BorderWidthBottom(borderWidth float64) *Segment {
	segment.delegated.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(borderWidth))
	return segment
}

func (segment *Segment) BorderWidthLeft(borderWidth float64) *Segment {
	segment.delegated.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(borderWidth))
	return segment
}

func (segment *Segment) BorderWidthHorizontal(borderWidth float64) *Segment {
	segment.delegated.BorderWidthLeft(borderWidth).BorderWidthRight(borderWidth)
	return segment
}

func (segment *Segment) BorderWidthVertical(borderWidth float64) *Segment {
	segment.delegated.BorderWidthTop(borderWidth).BorderWidthBottom(borderWidth)
	return segment
}

func (segment *Segment) BorderWidthAll(borderWidth float64) *Segment {
	segment.delegated.BorderWidthHorizontal(borderWidth).BorderWidthVertical(borderWidth)
	return segment
}

func (segment *Segment) BorderWidth(top float64, right float64, bottom float64, left float64) *Segment {
	segment.delegated.
		BorderWidthTop(top).
		BorderWidthRight(right).
		BorderWidthBottom(bottom).
		BorderWidthLeft(left)
	return segment
}

// BACKGROUND

func (segment *Segment) BackgroundColor(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.backgroundColor = getRgba(red, green, blue, alpha...)
	return segment
}

// BORDER COLOR

func (segment *Segment) BorderColorTop(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.borderColor[EdgeTop] = getRgba(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorRight(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.borderColor[EdgeRight] = getRgba(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorBottom(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.borderColor[EdgeBottom] = getRgba(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorLeft(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.borderColor[EdgeLeft] = getRgba(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorHorizontal(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.
		BorderColorLeft(red, green, blue, alpha...).
		BorderColorRight(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorVertical(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.
		BorderColorTop(red, green, blue, alpha...).
		BorderColorBottom(red, green, blue, alpha...)
	return segment
}

func (segment *Segment) BorderColorAll(red int, green int, blue int, alpha ...float64) *Segment {
	segment.delegated.
		BorderColorHorizontal(red, green, blue, alpha...).
		BorderColorVertical(red, green, blue, alpha...)
	return segment
}