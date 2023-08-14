// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import "github.com/kjk/flex"

func (block *BlockElement) WidthPercent(width float64) *BlockElement {
	block.getFlexNode().StyleSetWidthPercent(float32(width))
	return block
}

func (block *BlockElement) HeightPercent(height float64) *BlockElement {
	block.getFlexNode().StyleSetHeightPercent(float32(height))

	return block
}

func (block *BlockElement) WidthAuto() *BlockElement {
	block.getFlexNode().StyleSetWidthAuto()
	return block
}

func (block *BlockElement) HeightAuto() *BlockElement {
	block.getFlexNode().StyleSetHeightAuto()

	return block
}

// ALIGN SELF

func (block *BlockElement) AlignSelfType(align Align) *BlockElement {
	block.getFlexNode().StyleSetAlignSelf(flex.Align(align))
	return block
}

func (block *BlockElement) AlignSelfAuto() *BlockElement {
	block.AlignSelfType(AlignAuto)
	return block
}

func (block *BlockElement) AlignSelfStart() *BlockElement {
	block.AlignSelfType(AlignStart)
	return block
}

func (block *BlockElement) AlignSelfCenter() *BlockElement {
	block.AlignSelfType(AlignCenter)
	return block
}

func (block *BlockElement) AlignSelfEnd() *BlockElement {
	block.AlignSelfType(AlignEnd)
	return block
}

func (block *BlockElement) AlignSelfStretch() *BlockElement {
	block.AlignSelfType(AlignStretch)
	return block
}

func (block *BlockElement) AlignSelfBaseline() *BlockElement {
	block.AlignSelfType(AlignBaseline)
	return block
}

func (block *BlockElement) AlignSelfSpaceBetween() *BlockElement {
	block.AlignSelfType(AlignSpaceBetween)
	return block
}

func (block *BlockElement) AlignSelfSpaceAround() *BlockElement {
	block.AlignSelfType(AlignSpaceAround)
	return block
}

// FLEX FUNCTIONS

func (block *BlockElement) FlexGrow(flexGrow float64) *BlockElement {
	block.getFlexNode().StyleSetFlexGrow(float32(flexGrow))
	return block
}

func (block *BlockElement) FlexShrink(flexShrink float64) *BlockElement {
	block.getFlexNode().StyleSetFlexShrink(float32(flexShrink))
	return block
}

func (block *BlockElement) FlexBasis(flexBasis float64) *BlockElement {
	block.getFlexNode().StyleSetFlexBasis(float32(flexBasis))
	return block
}

func (block *BlockElement) FlexBasisPercent(flexBasisPercent float64) *BlockElement {
	block.getFlexNode().StyleSetFlexBasisPercent(float32(flexBasisPercent))
	return block
}

func (block *BlockElement) FlexBasisAuto() *BlockElement {
	flex.NodeStyleSetFlexBasisAuto(block.getFlexNode())
	return block
}


func (block *BlockElement) FlexAuto() *BlockElement {
	return block.
		FlexGrow(1).
		FlexShrink(1).
		FlexBasisAuto()
}

func (block *BlockElement) FlexNone() *BlockElement {
	return block.
		FlexGrow(0).
		FlexShrink(0).
		FlexBasisAuto()
}

func (block *BlockElement) MinWidth(minWidth float64) *BlockElement {
	block.getFlexNode().StyleSetMinWidth(float32(minWidth))
	return block
}

func (block *BlockElement) MinWidthPercent(minWidthPercent float64) *BlockElement {
	block.getFlexNode().StyleSetMinWidthPercent(float32(minWidthPercent))
	return block
}

func (block *BlockElement) MinHeight(minHeight float64) *BlockElement {
	block.getFlexNode().StyleSetMinHeight(float32(minHeight))
	return block
}

func (block *BlockElement) MinHeightPercent(minHeightPercent float64) *BlockElement {
	block.getFlexNode().StyleSetMinHeightPercent(float32(minHeightPercent))
	return block
}

func (block *BlockElement) MaxWidth(maxWidth float64) *BlockElement {
	block.getFlexNode().StyleSetMaxWidth(float32(maxWidth))
	return block
}

func (block *BlockElement) MaxWidthPercent(maxWidthPercent float64) *BlockElement {
	block.getFlexNode().StyleSetMaxWidthPercent(float32(maxWidthPercent))
	return block
}

func (block *BlockElement) MaxHeight(maxHeight float64) *BlockElement {
	block.getFlexNode().StyleSetMaxHeight(float32(maxHeight))
	return block
}

func (block *BlockElement) MaxHeightPercent(maxHeightPercent float64) *BlockElement {
	block.getFlexNode().StyleSetMaxHeightPercent(float32(maxHeightPercent))
	return block
}

// ASPECT RATIO

func (block *BlockElement) AspectRatio(ratio float64) *BlockElement {
	block.getFlexNode().StyleSetAspectRatio(float32(ratio))
	return block
}

// POSITION TYPE

func (block *BlockElement) PositionType(_type PositionType) *BlockElement {
	block.getFlexNode().StyleSetPositionType(flex.PositionType(_type))

	return block
}

func (block *BlockElement) PositionTypeRelative() *BlockElement {
	block.PositionType(PositionTypeRelative)

	return block
}

func (block *BlockElement) PositionTypeAbsolute() *BlockElement {
	block.PositionType(PositionTypeAbsolute)

	return block
}

// POSITION  TYPE

func (block *BlockElement) Position(_type PositionType) *BlockElement {
	block.getFlexNode().StyleSetPositionType(flex.PositionType(_type))
	return block
}

// POSITION

func (block *BlockElement) PositionTop(position float64) *BlockElement {
	block.getFlexNode().StyleSetPosition(flex.EdgeTop, float32(position))
	return block
}

func (block *BlockElement) PositionRight(position float64) *BlockElement {
	block.getFlexNode().StyleSetPosition(flex.EdgeRight, float32(position))
	return block
}

func (block *BlockElement) PositionBottom(position float64) *BlockElement {
	block.getFlexNode().StyleSetPosition(flex.EdgeBottom, float32(position))
	return block
}

func (block *BlockElement) PositionLeft(position float64) *BlockElement {
	block.getFlexNode().StyleSetPosition(flex.EdgeLeft, float32(position))
	return block
}

func (block *BlockElement) PositionAll(position float64) *BlockElement {
	block.
		PositionTop(position).
		PositionRight(position).
		PositionBottom(position).
		PositionLeft(position)
	return block
}
