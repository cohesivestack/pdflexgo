package pdflexgo

import (
	"github.com/kjk/flex"
)

func (text *TextElement) Width(width float64) *TextElement {
	text.getFlexNode().StyleSetWidth(float32(width))
	return text
}

func (text *TextElement) WidthPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetWidthPercent(float32(percent))
	return text
}

func (text *TextElement) Height(height float64) *TextElement {
	text.getFlexNode().StyleSetHeight(float32(height))
	return text
}

func (text *TextElement) HeightPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetHeightPercent(float32(percent))
	return text
}

func (text *TextElement) PositionType(position PositionType) *TextElement {
	text.getFlexNode().StyleSetPositionType(flex.PositionType(position))
	return text
}

func (text *TextElement) GetPositionType() PositionType {
	return PositionType(text.getFlexNode().Style.PositionType)
}

func (text *TextElement) Direction(direction Direction) *TextElement {
	text.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return text
}

func (text *TextElement) GetDirection() Direction {
	return Direction(text.getFlexNode().Style.Direction)
}

func (text *TextElement) FlexDirection(direction FlexDirection) *TextElement {
	text.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return text
}

func (text *TextElement) GetFlexDirection() FlexDirection {
	return FlexDirection(text.getFlexNode().Style.FlexDirection)
}

func (text *TextElement) JustifyContent(justifyContent Justify) *TextElement {
	text.getFlexNode().StyleSetJustifyContent(flex.Justify(justifyContent))
	return text
}

func (text *TextElement) GetJustifyContent() Justify {
	return Justify(text.getFlexNode().Style.JustifyContent)
}

func (text *TextElement) AlignContent(alignContent flex.Align) *TextElement {
	text.getFlexNode().StyleSetAlignContent(flex.Align(alignContent))
	return text
}

func (text *TextElement) GetAlignContent() flex.Align {
	return flex.Align(text.getFlexNode().Style.AlignContent)
}

func (text *TextElement) AlignItems(alignItems flex.Align) *TextElement {
	text.getFlexNode().StyleSetAlignItems(flex.Align(alignItems))
	return text
}

func (text *TextElement) GetAlignItems() flex.Align {
	return flex.Align(text.getFlexNode().Style.AlignItems)
}

func (text *TextElement) AlignSelf(alignSelf flex.Align) *TextElement {
	text.getFlexNode().StyleSetAlignSelf(flex.Align(alignSelf))
	return text
}

func (text *TextElement) GetAlignSelf() flex.Align {
	return flex.Align(text.getFlexNode().Style.AlignSelf)
}

func (text *TextElement) FlexWrap(wrap Wrap) *TextElement {
	text.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrap))
	return text
}

func (text *TextElement) GetFlexWrap() Wrap {
	return Wrap(text.getFlexNode().Style.FlexWrap)
}

func (text *TextElement) Overflow(overflow Overflow) *TextElement {
	text.getFlexNode().StyleSetOverflow(flex.Overflow(overflow))
	return text
}

func (text *TextElement) GetOverflow() Overflow {
	return Overflow(text.getFlexNode().Style.Overflow)
}

func (text *TextElement) Display(overflow Display) *TextElement {
	text.getFlexNode().StyleSetDisplay(flex.Display(overflow))
	return text
}

func (text *TextElement) GetDisplay() Display {
	return Display(text.getFlexNode().Style.Display)
}

func (text *TextElement) Flex(flex float64) *TextElement {
	text.getFlexNode().StyleSetFlex(float32(flex))
	return text
}

func (text *TextElement) GetFlex() float64 {
	return float64(text.getFlexNode().Style.Flex)
}

func (text *TextElement) FlexGrow(grow float64) *TextElement {
	text.getFlexNode().StyleSetFlexGrow(float32(grow))
	return text
}

func (text *TextElement) GetFlexGrow() float64 {
	return float64(text.getFlexNode().StyleGetFlexGrow())
}

func (text *TextElement) FlexShrink(shrink float64) *TextElement {
	text.getFlexNode().StyleSetFlexShrink(float32(shrink))
	return text
}

func (text *TextElement) GetFlexShrink() float64 {
	return float64(text.getFlexNode().StyleGetFlexShrink())
}

func (text *TextElement) FlexBasis(basis float64) *TextElement {
	text.getFlexNode().StyleSetFlexBasis(float32(basis))
	return text
}

func (text *TextElement) FlexBasisPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetFlexBasisPercent(float32(percent))
	return text
}

func (text *TextElement) MinWidth(minWidth float64) *TextElement {
	text.getFlexNode().StyleSetMinWidth(float32(minWidth))
	return text
}

func (text *TextElement) MinWidthPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetMinWidthPercent(float32(percent))
	return text
}

func (text *TextElement) MinHeight(minHeight float64) *TextElement {
	text.getFlexNode().StyleSetMinHeight(float32(minHeight))
	return text
}

func (text *TextElement) MinHeightPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetMinHeightPercent(float32(percent))
	return text
}

func (text *TextElement) MaxWidth(maxWidth float64) *TextElement {
	text.getFlexNode().StyleSetMaxWidth(float32(maxWidth))
	return text
}

func (text *TextElement) MaxWidthPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetMaxWidthPercent(float32(percent))
	return text
}

func (text *TextElement) MaxHeight(maxHeight float64) *TextElement {
	text.getFlexNode().StyleSetMaxHeight(float32(maxHeight))
	return text
}

func (text *TextElement) MaxHeightPercent(percent float64) *TextElement {
	text.getFlexNode().StyleSetMaxHeightPercent(float32(percent))
	return text
}

func (text *TextElement) AspectRatio(aspectRatio float64) *TextElement {
	text.getFlexNode().StyleSetAspectRatio(float32(aspectRatio))
	return text
}

func (element *TextElement) BackgroundColor(color string) *TextElement {
	element.backgroundColor = color

	return element
}
