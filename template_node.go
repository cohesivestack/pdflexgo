package pdflexgo

import "github.com/kjk/flex"

func (templateNode *TemplateElement) WidthPercent(width float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetWidthPercent(float32(width))
	return templateNode
}

func (templateNode *TemplateElement) HeightPercent(height float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetHeightPercent(float32(height))

	return templateNode
}

// ALIGN SELF

func (templateNode *TemplateElement) AlignSelfType(align Align) *TemplateElement {
	templateNode.getFlexNode().StyleSetAlignSelf(flex.Align(align))
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfAuto() *TemplateElement {
	templateNode.AlignSelfType(AlignAuto)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfStart() *TemplateElement {
	templateNode.AlignSelfType(AlignStart)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfCenter() *TemplateElement {
	templateNode.AlignSelfType(AlignCenter)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfEnd() *TemplateElement {
	templateNode.AlignSelfType(AlignEnd)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfStretch() *TemplateElement {
	templateNode.AlignSelfType(AlignStretch)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfBaseline() *TemplateElement {
	templateNode.AlignSelfType(AlignBaseline)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfSpaceBetween() *TemplateElement {
	templateNode.AlignSelfType(AlignSpaceBetween)
	return templateNode
}

func (templateNode *TemplateElement) AlignSelfSpaceAround() *TemplateElement {
	templateNode.AlignSelfType(AlignSpaceAround)
	return templateNode
}

// FLEX FUNCTIONS

func (templateNode *TemplateElement) FlexGrow(flexGrow float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexGrow(float32(flexGrow))
	return templateNode
}

func (templateNode *TemplateElement) FlexShrink(flexShrink float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexShrink(float32(flexShrink))
	return templateNode
}

func (templateNode *TemplateElement) FlexBasis(flexBasis float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexBasis(float32(flexBasis))
	return templateNode
}

func (templateNode *TemplateElement) FlexBasisPercent(flexBasisPercent float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexBasisPercent(float32(flexBasisPercent))
	return templateNode
}

func (templateNode *TemplateElement) MinWidth(minWidth float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMinWidth(float32(minWidth))
	return templateNode
}

func (templateNode *TemplateElement) MinWidthPercent(minWidthPercent float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMinWidthPercent(float32(minWidthPercent))
	return templateNode
}

func (templateNode *TemplateElement) MinHeight(minHeight float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMinHeight(float32(minHeight))
	return templateNode
}

func (templateNode *TemplateElement) MinHeightPercent(minHeightPercent float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMinHeightPercent(float32(minHeightPercent))
	return templateNode
}

func (templateNode *TemplateElement) MaxWidth(maxWidth float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMaxWidth(float32(maxWidth))
	return templateNode
}

func (templateNode *TemplateElement) MaxWidthPercent(maxWidthPercent float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMaxWidthPercent(float32(maxWidthPercent))
	return templateNode
}

func (templateNode *TemplateElement) MaxHeight(maxHeight float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMaxHeight(float32(maxHeight))
	return templateNode
}

func (templateNode *TemplateElement) MaxHeightPercent(maxHeightPercent float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetMaxHeightPercent(float32(maxHeightPercent))
	return templateNode
}

// ASPECT RATIO

func (templateNode *TemplateElement) AspectRatio(ratio float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetAspectRatio(float32(ratio))
	return templateNode
}

// POSITION TYPE

func (templateNode *TemplateElement) PositionType(_type PositionType) *TemplateElement {
	templateNode.getFlexNode().StyleSetPositionType(flex.PositionType(_type))

	return templateNode
}

func (templateNode *TemplateElement) PositionTypeRelative() *TemplateElement {
	templateNode.PositionType(PositionTypeRelative)

	return templateNode
}

func (templateNode *TemplateElement) PositionTypeAbsolute() *TemplateElement {
	templateNode.PositionType(PositionTypeAbsolute)

	return templateNode
}

// POSITION  TYPE

func (templateNode *TemplateElement) Position(_type PositionType) *TemplateElement {
	templateNode.getFlexNode().StyleSetPositionType(flex.PositionType(_type))
	return templateNode
}

// POSITION

func (templateNode *TemplateElement) PositionTop(position float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetPosition(flex.EdgeTop, float32(position))
	return templateNode
}

func (templateNode *TemplateElement) PositionRight(position float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetPosition(flex.EdgeRight, float32(position))
	return templateNode
}

func (templateNode *TemplateElement) PositionBottom(position float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetPosition(flex.EdgeBottom, float32(position))
	return templateNode
}

func (templateNode *TemplateElement) PositionLeft(position float64) *TemplateElement {
	templateNode.getFlexNode().StyleSetPosition(flex.EdgeLeft, float32(position))
	return templateNode
}

func (templateNode *TemplateElement) PositionAll(position float64) *TemplateElement {
	templateNode.
		PositionTop(position).
		PositionRight(position).
		PositionBottom(position).
		PositionLeft(position)
	return templateNode
}
