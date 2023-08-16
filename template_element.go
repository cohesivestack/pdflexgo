package pdflexgo

import (
	"github.com/kjk/flex"
)

type TemplateElement struct {
	abstractElement
	backgroundColor rgba
	borderColor     [edgeCount]rgba
}

func (templateElement *TemplateElement) Width(width float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetWidth(float32(width))
	return templateElement
}
func (templateElement *TemplateElement) Height(height float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetHeight(float32(height))
	return templateElement
}

// MARGIN

func (templateElement *TemplateElement) MarginTop(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMargin(flex.EdgeTop, float32(margin))
	return templateElement
}
func (templateElement *TemplateElement) MarginRight(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMargin(flex.EdgeRight, float32(margin))
	return templateElement
}
func (templateElement *TemplateElement) MarginBottom(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMargin(flex.EdgeBottom, float32(margin))
	return templateElement
}
func (templateElement *TemplateElement) MarginLeft(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMargin(flex.EdgeLeft, float32(margin))
	return templateElement
}
func (templateElement *TemplateElement) MarginHorizontal(margin float64) *TemplateElement {
	templateElement.MarginLeft(margin).MarginRight(margin)
	return templateElement
}
func (templateElement *TemplateElement) MarginVertical(margin float64) *TemplateElement {
	templateElement.MarginTop(margin).MarginBottom(margin)
	return templateElement
}
func (templateElement *TemplateElement) MarginAll(margin float64) *TemplateElement {
	templateElement.MarginHorizontal(margin).MarginVertical(margin)
	return templateElement
}
func (templateElement *TemplateElement) Margin(top float64, right float64, bottom float64, left float64) *TemplateElement {
	templateElement.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)
	return templateElement
}

// MARGIN PERCENT

func (templateElement *TemplateElement) MarginPercentTop(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMarginPercent(flex.EdgeTop, float32(margin))
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentRight(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMarginPercent(flex.EdgeRight, float32(margin))
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentBottom(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMarginPercent(flex.EdgeBottom, float32(margin))
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentLeft(margin float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetMarginPercent(flex.EdgeLeft, float32(margin))
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentHorizontal(margin float64) *TemplateElement {
	templateElement.MarginPercentLeft(margin).MarginPercentRight(margin)
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentVertical(margin float64) *TemplateElement {
	templateElement.MarginPercentTop(margin).MarginPercentBottom(margin)
	return templateElement
}

func (templateElement *TemplateElement) MarginPercentAll(margin float64) *TemplateElement {
	templateElement.MarginPercentHorizontal(margin).MarginPercentVertical(margin)
	return templateElement
}

// PADDING

// For Padding

func (templateElement *TemplateElement) PaddingTop(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPadding(flex.EdgeTop, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingRight(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPadding(flex.EdgeRight, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingBottom(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPadding(flex.EdgeBottom, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingLeft(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPadding(flex.EdgeLeft, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingHorizontal(padding float64) *TemplateElement {
	templateElement.PaddingLeft(padding).PaddingRight(padding)
	return templateElement
}

func (templateElement *TemplateElement) PaddingVertical(padding float64) *TemplateElement {
	templateElement.PaddingTop(padding).PaddingBottom(padding)
	return templateElement
}

func (templateElement *TemplateElement) PaddingAll(padding float64) *TemplateElement {
	templateElement.PaddingHorizontal(padding).PaddingVertical(padding)
	return templateElement
}

func (templateElement *TemplateElement) Padding(top float64, right float64, bottom float64, left float64) *TemplateElement {
	templateElement.
		PaddingTop(top).
		PaddingRight(right).
		PaddingBottom(bottom).
		PaddingLeft(left)
	return templateElement
}

// PADDING PERCENT

func (templateElement *TemplateElement) PaddingPercentTop(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPaddingPercent(flex.EdgeTop, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentRight(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPaddingPercent(flex.EdgeRight, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentBottom(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPaddingPercent(flex.EdgeBottom, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentLeft(padding float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetPaddingPercent(flex.EdgeLeft, float32(padding))
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentHorizontal(padding float64) *TemplateElement {
	templateElement.PaddingPercentLeft(padding).PaddingPercentRight(padding)
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentVertical(padding float64) *TemplateElement {
	templateElement.PaddingPercentTop(padding).PaddingPercentBottom(padding)
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercentAll(padding float64) *TemplateElement {
	templateElement.PaddingPercentHorizontal(padding).PaddingPercentVertical(padding)
	return templateElement
}

func (templateElement *TemplateElement) PaddingPercent(top float64, right float64, bottom float64, left float64) *TemplateElement {
	templateElement.
		PaddingPercentTop(top).
		PaddingPercentRight(right).
		PaddingPercentBottom(bottom).
		PaddingPercentLeft(left)
	return templateElement
}

// BORDERWIDTH

func (templateElement *TemplateElement) BorderWidthTop(borderWidth float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetBorder(flex.EdgeTop, float32(borderWidth))
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthRight(borderWidth float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetBorder(flex.EdgeRight, float32(borderWidth))
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthBottom(borderWidth float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetBorder(flex.EdgeBottom, float32(borderWidth))
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthLeft(borderWidth float64) *TemplateElement {
	templateElement.getFlexNode().StyleSetBorder(flex.EdgeLeft, float32(borderWidth))
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthHorizontal(borderWidth float64) *TemplateElement {
	templateElement.BorderWidthLeft(borderWidth).BorderWidthRight(borderWidth)
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthVertical(borderWidth float64) *TemplateElement {
	templateElement.BorderWidthTop(borderWidth).BorderWidthBottom(borderWidth)
	return templateElement
}

func (templateElement *TemplateElement) BorderWidthAll(borderWidth float64) *TemplateElement {
	templateElement.BorderWidthHorizontal(borderWidth).BorderWidthVertical(borderWidth)
	return templateElement
}

func (templateElement *TemplateElement) BorderWidth(top float64, right float64, bottom float64, left float64) *TemplateElement {
	templateElement.
		BorderWidthTop(top).
		BorderWidthRight(right).
		BorderWidthBottom(bottom).
		BorderWidthLeft(left)
	return templateElement
}

// BACKGROUND

func (templateElement *TemplateElement) BackgroundColor(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.backgroundColor = getRgba(red, green, blue, alpha...)
	return templateElement
}

// BORDER COLOR

func (templateElement *TemplateElement) BorderColorTop(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.borderColor[EdgeTop] = getRgba(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorRight(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.borderColor[EdgeRight] = getRgba(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorBottom(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.borderColor[EdgeBottom] = getRgba(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorLeft(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.borderColor[EdgeLeft] = getRgba(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorHorizontal(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.
		BorderColorLeft(red, green, blue, alpha...).
		BorderColorRight(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorVertical(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.
		BorderColorTop(red, green, blue, alpha...).
		BorderColorBottom(red, green, blue, alpha...)
	return templateElement
}

func (templateElement *TemplateElement) BorderColorAll(red int, green int, blue int, alpha ...float64) *TemplateElement {
	templateElement.
		BorderColorHorizontal(red, green, blue, alpha...).
		BorderColorVertical(red, green, blue, alpha...)
	return templateElement
}
