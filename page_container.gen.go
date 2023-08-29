// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import "github.com/kjk/flex"

// CHILDREN

func (page *PageElement) Children(children ...Node) *PageElement {

	for _, child := range children {
		page.body.getFlexNode().InsertChild(child.getFlexNode(), len(page.body.getFlexNode().Children))
		page.body.children = append(page.body.children, child)
	}

	return page
}

// DIRECTION

func (page *PageElement) Direction(direction Direction) *PageElement {
	page.body.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return page
}

func (page *PageElement) DirectionInherit() *PageElement {
	page.body.Direction(DirectionInherit)
	return page
}

func (page *PageElement) DirectionLTR() *PageElement {
	page.body.Direction(DirectionLTR)
	return page
}

func (page *PageElement) DirectionRTL() *PageElement {
	page.body.Direction(DirectionRTL)
	return page
}

// FLEX DIRECTION

func (page *PageElement) FlexDirection(direction FlexDirection) *PageElement {
	page.body.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return page
}

func (page *PageElement) FlexDirectionColumn() *PageElement {
	page.body.FlexDirection(FlexDirectionColumn)
	return page
}

func (page *PageElement) FlexDirectionColumnReverse() *PageElement {
	page.body.FlexDirection(FlexDirectionColumnReverse)
	return page
}

func (page *PageElement) FlexDirectionRow() *PageElement {
	page.body.FlexDirection(FlexDirectionRow)
	return page
}

func (page *PageElement) FlexDirectionRowReverse() *PageElement {
	page.body.FlexDirection(FlexDirectionRowReverse)
	return page
}

// JustifyContent
func (page *PageElement) JustifyContent(justification Justify) *PageElement {
	page.body.getFlexNode().StyleSetJustifyContent(flex.Justify(justification))
	return page
}

func (page *PageElement) JustifyContentStart() *PageElement {
	page.body.JustifyContent(JustifyStart)
	return page
}

func (page *PageElement) JustifyContentCenter() *PageElement {
	page.body.JustifyContent(JustifyCenter)
	return page
}

func (page *PageElement) JustifyContentEnd() *PageElement {
	page.body.JustifyContent(JustifyEnd)
	return page
}

func (page *PageElement) JustifyContentSpaceBetween() *PageElement {
	page.body.JustifyContent(JustifySpaceBetween)
	return page
}

func (page *PageElement) JustifyContentSpaceAround() *PageElement {
	page.body.JustifyContent(JustifySpaceAround)
	return page
}

// ALIGN CONTENT

func (page *PageElement) AlignContentType(align Align) *PageElement {
	page.body.getFlexNode().StyleSetAlignContent(flex.Align(align))
	return page
}

func (page *PageElement) AlignContentAuto() *PageElement {
	page.body.AlignContentType(AlignAuto)
	return page
}

func (page *PageElement) AlignContentStart() *PageElement {
	page.body.AlignContentType(AlignStart)
	return page
}

func (page *PageElement) AlignContentCenter() *PageElement {
	page.body.AlignContentType(AlignCenter)
	return page
}

func (page *PageElement) AlignContentEnd() *PageElement {
	page.body.AlignContentType(AlignEnd)
	return page
}

func (page *PageElement) AlignContentStretch() *PageElement {
	page.body.AlignContentType(AlignStretch)
	return page
}

func (page *PageElement) AlignContentBaseline() *PageElement {
	page.body.AlignContentType(AlignBaseline)
	return page
}

func (page *PageElement) AlignContentSpaceBetween() *PageElement {
	page.body.AlignContentType(AlignSpaceBetween)
	return page
}

func (page *PageElement) AlignContentSpaceAround() *PageElement {
	page.body.AlignContentType(AlignSpaceAround)
	return page
}

// ALIGN ITEMS

func (page *PageElement) AlignItemsType(align Align) *PageElement {
	page.body.getFlexNode().StyleSetAlignItems(flex.Align(align))
	return page
}

func (page *PageElement) AlignItemsAuto() *PageElement {
	page.body.AlignItemsType(AlignAuto)
	return page
}

func (page *PageElement) AlignItemsStart() *PageElement {
	page.body.AlignItemsType(AlignStart)
	return page
}

func (page *PageElement) AlignItemsCenter() *PageElement {
	page.body.AlignItemsType(AlignCenter)
	return page
}

func (page *PageElement) AlignItemsEnd() *PageElement {
	page.body.AlignItemsType(AlignEnd)
	return page
}

func (page *PageElement) AlignItemsStretch() *PageElement {
	page.body.AlignItemsType(AlignStretch)
	return page
}

func (page *PageElement) AlignItemsBaseline() *PageElement {
	page.body.AlignItemsType(AlignBaseline)
	return page
}

func (page *PageElement) AlignItemsSpaceBetween() *PageElement {
	page.body.AlignItemsType(AlignSpaceBetween)
	return page
}

func (page *PageElement) AlignItemsSpaceAround() *PageElement {
	page.body.AlignItemsType(AlignSpaceAround)
	return page
}

// FLEX WRAP

func (page *PageElement) FlexWrapValue(wrapValue Wrap) *PageElement {
	page.body.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrapValue))
	return page
}

func (page *PageElement) FlexNoWrap() *PageElement {
	page.body.FlexWrapValue(WrapNoWrap)
	return page
}

func (page *PageElement) FlexWrap() *PageElement {
	page.body.FlexWrapValue(WrapWrap)
	return page
}

func (page *PageElement) FlexWrapReverse() *PageElement {
	page.body.FlexWrapValue(WrapWrapReverse)
	return page
}
