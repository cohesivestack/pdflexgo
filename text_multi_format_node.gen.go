// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import "github.com/kjk/flex"

func (textMultiFormat *TextMultiFormatElement) WidthPercent(width float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetWidthPercent(float32(width))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) HeightPercent(height float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetHeightPercent(float32(height))

	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) WidthAuto() *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetWidthAuto()
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) HeightAuto() *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetHeightAuto()

	return textMultiFormat
}

// ALIGN SELF

func (textMultiFormat *TextMultiFormatElement) AlignSelfType(align Align) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetAlignSelf(flex.Align(align))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfAuto() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignAuto)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfStart() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignStart)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfCenter() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignCenter)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfEnd() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignEnd)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfStretch() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignStretch)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfBaseline() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignBaseline)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfSpaceBetween() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignSpaceBetween)
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) AlignSelfSpaceAround() *TextMultiFormatElement {
	textMultiFormat.AlignSelfType(AlignSpaceAround)
	return textMultiFormat
}

// FLEX FUNCTIONS

func (textMultiFormat *TextMultiFormatElement) FlexGrow(flexGrow float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetFlexGrow(float32(flexGrow))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) FlexShrink(flexShrink float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetFlexShrink(float32(flexShrink))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) FlexBasis(flexBasis float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetFlexBasis(float32(flexBasis))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) FlexBasisPercent(flexBasisPercent float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetFlexBasisPercent(float32(flexBasisPercent))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) FlexBasisAuto() *TextMultiFormatElement {
	flex.NodeStyleSetFlexBasisAuto(textMultiFormat.getFlexNode())
	return textMultiFormat
}


func (textMultiFormat *TextMultiFormatElement) FlexAuto() *TextMultiFormatElement {
	return textMultiFormat.
		FlexGrow(1).
		FlexShrink(1).
		FlexBasisAuto()
}

func (textMultiFormat *TextMultiFormatElement) FlexNone() *TextMultiFormatElement {
	return textMultiFormat.
		FlexGrow(0).
		FlexShrink(0).
		FlexBasisAuto()
}

func (textMultiFormat *TextMultiFormatElement) MinWidth(minWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMinWidth(float32(minWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MinWidthPercent(minWidthPercent float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMinWidthPercent(float32(minWidthPercent))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MinHeight(minHeight float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMinHeight(float32(minHeight))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MinHeightPercent(minHeightPercent float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMinHeightPercent(float32(minHeightPercent))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MaxWidth(maxWidth float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMaxWidth(float32(maxWidth))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MaxWidthPercent(maxWidthPercent float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMaxWidthPercent(float32(maxWidthPercent))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MaxHeight(maxHeight float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMaxHeight(float32(maxHeight))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) MaxHeightPercent(maxHeightPercent float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetMaxHeightPercent(float32(maxHeightPercent))
	return textMultiFormat
}

// ASPECT RATIO

func (textMultiFormat *TextMultiFormatElement) AspectRatio(ratio float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetAspectRatio(float32(ratio))
	return textMultiFormat
}

// POSITION TYPE

func (textMultiFormat *TextMultiFormatElement) PositionType(_type PositionType) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPositionType(flex.PositionType(_type))

	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionTypeRelative() *TextMultiFormatElement {
	textMultiFormat.PositionType(PositionTypeRelative)

	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionTypeAbsolute() *TextMultiFormatElement {
	textMultiFormat.PositionType(PositionTypeAbsolute)

	return textMultiFormat
}

// POSITION  TYPE

func (textMultiFormat *TextMultiFormatElement) Position(_type PositionType) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPositionType(flex.PositionType(_type))
	return textMultiFormat
}

// POSITION

func (textMultiFormat *TextMultiFormatElement) PositionTop(position float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPosition(flex.EdgeTop, float32(position))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionRight(position float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPosition(flex.EdgeRight, float32(position))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionBottom(position float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPosition(flex.EdgeBottom, float32(position))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionLeft(position float64) *TextMultiFormatElement {
	textMultiFormat.getFlexNode().StyleSetPosition(flex.EdgeLeft, float32(position))
	return textMultiFormat
}

func (textMultiFormat *TextMultiFormatElement) PositionAll(position float64) *TextMultiFormatElement {
	textMultiFormat.
		PositionTop(position).
		PositionRight(position).
		PositionBottom(position).
		PositionLeft(position)
	return textMultiFormat
}
