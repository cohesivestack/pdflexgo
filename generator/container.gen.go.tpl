// Code generated by PDFlexgo; DO NOT EDIT.
package pdflexgo

import "github.com/kjk/flex"

// CHILDREN

func ({{ .Name }} *{{ .Type }}) Children(children ...Node) *{{ .Type }} {

	for _, child := range children {
		{{ .Name }}{{ .Delegated }}.getFlexNode().InsertChild(child.getFlexNode(), len({{ .Name }}{{ .Delegated }}.getFlexNode().Children))
		{{ .Name }}{{ .Delegated }}.children = append({{ .Name }}{{ .Delegated }}.children, child)
	}

	return {{ .Name }}
}

// DIRECTION

func ({{ .Name }} *{{ .Type }}) Direction(direction Direction) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) DirectionInherit() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.Direction(DirectionInherit)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) DirectionLTR() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.Direction(DirectionLTR)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) DirectionRTL() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.Direction(DirectionRTL)
	return {{ .Name }}
}

// FLEX DIRECTION

func ({{ .Name }} *{{ .Type }}) FlexDirection(direction FlexDirection) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexDirectionColumn() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexDirection(FlexDirectionColumn)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexDirectionColumnReverse() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexDirection(FlexDirectionColumnReverse)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexDirectionRow() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexDirection(FlexDirectionRow)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexDirectionRowReverse() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexDirection(FlexDirectionRowReverse)
	return {{ .Name }}
}

// JustifyContent
func ({{ .Name }} *{{ .Type }}) JustifyContent(justification Justify) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetJustifyContent(flex.Justify(justification))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) JustifyContentStart() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.JustifyContent(JustifyStart)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) JustifyContentCenter() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.JustifyContent(JustifyCenter)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) JustifyContentEnd() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.JustifyContent(JustifyEnd)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) JustifyContentSpaceBetween() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.JustifyContent(JustifySpaceBetween)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) JustifyContentSpaceAround() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.JustifyContent(JustifySpaceAround)
	return {{ .Name }}
}

// ALIGN CONTENT

func ({{ .Name }} *{{ .Type }}) AlignContentType(align Align) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetAlignContent(flex.Align(align))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentAuto() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignAuto)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentStart() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignStart)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentCenter() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignCenter)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentEnd() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignEnd)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentStretch() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignStretch)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentBaseline() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignBaseline)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentSpaceBetween() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignSpaceBetween)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignContentSpaceAround() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignContentType(AlignSpaceAround)
	return {{ .Name }}
}

// ALIGN ITEMS

func ({{ .Name }} *{{ .Type }}) AlignItemsType(align Align) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetAlignItems(flex.Align(align))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsAuto() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignAuto)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsStart() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignStart)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsCenter() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignCenter)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsEnd() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignEnd)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsStretch() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignStretch)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsBaseline() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignBaseline)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsSpaceBetween() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignSpaceBetween)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) AlignItemsSpaceAround() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.AlignItemsType(AlignSpaceAround)
	return {{ .Name }}
}

// FLEX WRAP

func ({{ .Name }} *{{ .Type }}) FlexWrapValue(wrapValue Wrap) *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrapValue))
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexNoWrap() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexWrapValue(WrapNoWrap)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexWrap() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexWrapValue(WrapWrap)
	return {{ .Name }}
}

func ({{ .Name }} *{{ .Type }}) FlexWrapReverse() *{{ .Type }} {
	{{ .Name }}{{ .Delegated }}.FlexWrapValue(WrapWrapReverse)
	return {{ .Name }}
}
