// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import "github.com/kjk/flex"

// CHILDREN

func (header *Header) Children(children ...Node) *Header {

	for _, child := range children {
		header.delegated.getFlexNode().InsertChild(child.getFlexNode(), len(header.delegated.getFlexNode().Children))
		header.delegated.children = append(header.delegated.children, child)
	}

	return header
}

// DIRECTION

func (header *Header) Direction(direction Direction) *Header {
	header.delegated.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return header
}

func (header *Header) DirectionInherit() *Header {
	header.delegated.Direction(DirectionInherit)
	return header
}

func (header *Header) DirectionLTR() *Header {
	header.delegated.Direction(DirectionLTR)
	return header
}

func (header *Header) DirectionRTL() *Header {
	header.delegated.Direction(DirectionRTL)
	return header
}

// FLEX DIRECTION

func (header *Header) FlexDirection(direction FlexDirection) *Header {
	header.delegated.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return header
}

func (header *Header) FlexDirectionColumn() *Header {
	header.delegated.FlexDirection(FlexDirectionColumn)
	return header
}

func (header *Header) FlexDirectionColumnReverse() *Header {
	header.delegated.FlexDirection(FlexDirectionColumnReverse)
	return header
}

func (header *Header) FlexDirectionRow() *Header {
	header.delegated.FlexDirection(FlexDirectionRow)
	return header
}

func (header *Header) FlexDirectionRowReverse() *Header {
	header.delegated.FlexDirection(FlexDirectionRowReverse)
	return header
}

// JustifyContent
func (header *Header) JustifyContent(justification Justify) *Header {
	header.delegated.getFlexNode().StyleSetJustifyContent(flex.Justify(justification))
	return header
}

func (header *Header) JustifyContentStart() *Header {
	header.delegated.JustifyContent(JustifyStart)
	return header
}

func (header *Header) JustifyContentCenter() *Header {
	header.delegated.JustifyContent(JustifyCenter)
	return header
}

func (header *Header) JustifyContentEnd() *Header {
	header.delegated.JustifyContent(JustifyEnd)
	return header
}

func (header *Header) JustifyContentSpaceBetween() *Header {
	header.delegated.JustifyContent(JustifySpaceBetween)
	return header
}

func (header *Header) JustifyContentSpaceAround() *Header {
	header.delegated.JustifyContent(JustifySpaceAround)
	return header
}

// ALIGN CONTENT

func (header *Header) AlignContentType(align Align) *Header {
	header.delegated.getFlexNode().StyleSetAlignContent(flex.Align(align))
	return header
}

func (header *Header) AlignContentAuto() *Header {
	header.delegated.AlignContentType(AlignAuto)
	return header
}

func (header *Header) AlignContentStart() *Header {
	header.delegated.AlignContentType(AlignStart)
	return header
}

func (header *Header) AlignContentCenter() *Header {
	header.delegated.AlignContentType(AlignCenter)
	return header
}

func (header *Header) AlignContentEnd() *Header {
	header.delegated.AlignContentType(AlignEnd)
	return header
}

func (header *Header) AlignContentStretch() *Header {
	header.delegated.AlignContentType(AlignStretch)
	return header
}

func (header *Header) AlignContentBaseline() *Header {
	header.delegated.AlignContentType(AlignBaseline)
	return header
}

func (header *Header) AlignContentSpaceBetween() *Header {
	header.delegated.AlignContentType(AlignSpaceBetween)
	return header
}

func (header *Header) AlignContentSpaceAround() *Header {
	header.delegated.AlignContentType(AlignSpaceAround)
	return header
}

// ALIGN ITEMS

func (header *Header) AlignItemsType(align Align) *Header {
	header.delegated.getFlexNode().StyleSetAlignItems(flex.Align(align))
	return header
}

func (header *Header) AlignItemsAuto() *Header {
	header.delegated.AlignItemsType(AlignAuto)
	return header
}

func (header *Header) AlignItemsStart() *Header {
	header.delegated.AlignItemsType(AlignStart)
	return header
}

func (header *Header) AlignItemsCenter() *Header {
	header.delegated.AlignItemsType(AlignCenter)
	return header
}

func (header *Header) AlignItemsEnd() *Header {
	header.delegated.AlignItemsType(AlignEnd)
	return header
}

func (header *Header) AlignItemsStretch() *Header {
	header.delegated.AlignItemsType(AlignStretch)
	return header
}

func (header *Header) AlignItemsBaseline() *Header {
	header.delegated.AlignItemsType(AlignBaseline)
	return header
}

func (header *Header) AlignItemsSpaceBetween() *Header {
	header.delegated.AlignItemsType(AlignSpaceBetween)
	return header
}

func (header *Header) AlignItemsSpaceAround() *Header {
	header.delegated.AlignItemsType(AlignSpaceAround)
	return header
}

// FLEX WRAP

func (header *Header) FlexWrapValue(wrapValue Wrap) *Header {
	header.delegated.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrapValue))
	return header
}

func (header *Header) FlexNoWrap() *Header {
	header.delegated.FlexWrapValue(WrapNoWrap)
	return header
}

func (header *Header) FlexWrap() *Header {
	header.delegated.FlexWrapValue(WrapWrap)
	return header
}

func (header *Header) FlexWrapReverse() *Header {
	header.delegated.FlexWrapValue(WrapWrapReverse)
	return header
}
