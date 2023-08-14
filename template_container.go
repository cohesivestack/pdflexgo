package pdflexgo

import "github.com/kjk/flex"

// DIRECTION

func (templateNode *TemplateElement) Direction(direction Direction) *TemplateElement {
	templateNode.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return templateNode
}

func (templateNode *TemplateElement) DirectionInherit() *TemplateElement {
	return templateNode.Direction(DirectionInherit)
}

func (templateNode *TemplateElement) DirectionLTR() *TemplateElement {
	return templateNode.Direction(DirectionLTR)
}

func (templateNode *TemplateElement) DirectionRTL() *TemplateElement {
	return templateNode.Direction(DirectionRTL)
}

// FLEX DIRECTION

func (templateNode *TemplateElement) FlexDirection(direction FlexDirection) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return templateNode
}

func (templateNode *TemplateElement) FlexDirectionColumn() *TemplateElement {
	return templateNode.FlexDirection(FlexDirectionColumn)
}

func (templateNode *TemplateElement) FlexDirectionColumnReverse() *TemplateElement {
	return templateNode.FlexDirection(FlexDirectionColumnReverse)
}

func (templateNode *TemplateElement) FlexDirectionRow() *TemplateElement {
	return templateNode.FlexDirection(FlexDirectionRow)
}

func (templateNode *TemplateElement) FlexDirectionRowReverse() *TemplateElement {
	return templateNode.FlexDirection(FlexDirectionRowReverse)
}

// JustifyContent
func (templateNode *TemplateElement) Justify(justification Justify) *TemplateElement {
	templateNode.getFlexNode().StyleSetJustifyContent(flex.Justify(justification))
	return templateNode
}

func (templateNode *TemplateElement) JustifyStart() *TemplateElement {
	return templateNode.Justify(JustifyStart)
}

func (templateNode *TemplateElement) JustifyCenter() *TemplateElement {
	return templateNode.Justify(JustifyCenter)
}

func (templateNode *TemplateElement) JustifyEnd() *TemplateElement {
	return templateNode.Justify(JustifyEnd)
}

func (templateNode *TemplateElement) JustifySpaceBetween() *TemplateElement {
	return templateNode.Justify(JustifySpaceBetween)
}

func (templateNode *TemplateElement) JustifySpaceAround() *TemplateElement {
	return templateNode.Justify(JustifySpaceAround)
}

// ALIGN CONTENT

func (templateNode *TemplateElement) AlignContentType(align Align) *TemplateElement {
	templateNode.getFlexNode().StyleSetAlignContent(flex.Align(align))
	return templateNode
}

func (templateNode *TemplateElement) AlignContentAuto() *TemplateElement {
	templateNode.AlignContentType(AlignAuto)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentStart() *TemplateElement {
	templateNode.AlignContentType(AlignStart)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentCenter() *TemplateElement {
	templateNode.AlignContentType(AlignCenter)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentEnd() *TemplateElement {
	templateNode.AlignContentType(AlignEnd)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentStretch() *TemplateElement {
	templateNode.AlignContentType(AlignStretch)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentBaseline() *TemplateElement {
	templateNode.AlignContentType(AlignBaseline)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentSpaceBetween() *TemplateElement {
	templateNode.AlignContentType(AlignSpaceBetween)
	return templateNode
}

func (templateNode *TemplateElement) AlignContentSpaceAround() *TemplateElement {
	templateNode.AlignContentType(AlignSpaceAround)
	return templateNode
}

// ALIGN ITEMS

func (templateNode *TemplateElement) AlignItemsType(align Align) *TemplateElement {
	templateNode.getFlexNode().StyleSetAlignItems(flex.Align(align))
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsAuto() *TemplateElement {
	templateNode.AlignItemsType(AlignAuto)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsStart() *TemplateElement {
	templateNode.AlignItemsType(AlignStart)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsCenter() *TemplateElement {
	templateNode.AlignItemsType(AlignCenter)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsEnd() *TemplateElement {
	templateNode.AlignItemsType(AlignEnd)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsStretch() *TemplateElement {
	templateNode.AlignItemsType(AlignStretch)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsBaseline() *TemplateElement {
	templateNode.AlignItemsType(AlignBaseline)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsSpaceBetween() *TemplateElement {
	templateNode.AlignItemsType(AlignSpaceBetween)
	return templateNode
}

func (templateNode *TemplateElement) AlignItemsSpaceAround() *TemplateElement {
	templateNode.AlignItemsType(AlignSpaceAround)
	return templateNode
}

// FLEX WRAP

func (templateNode *TemplateElement) FlexWrapValue(wrapValue Wrap) *TemplateElement {
	templateNode.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrapValue))
	return templateNode
}

func (templateNode *TemplateElement) FlexNoWrap() *TemplateElement {
	return templateNode.FlexWrapValue(WrapNoWrap)
}

func (templateNode *TemplateElement) FlexWrap() *TemplateElement {
	return templateNode.FlexWrapValue(WrapWrap)
}

func (templateNode *TemplateElement) FlexWrapReverse() *TemplateElement {
	return templateNode.FlexWrapValue(WrapWrapReverse)
}
