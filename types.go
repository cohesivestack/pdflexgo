package pdflexgo

import "github.com/kjk/flex"

// Font Style
type FontStyle string

const (
	FontStyleThin             FontStyle = "Thin"
	FontStyleExtraLight       FontStyle = "ExtraLight"
	FontStyleLight            FontStyle = "Light"
	FontStyleRegular          FontStyle = "Regular"
	FontStyleMedium           FontStyle = "Medium"
	FontStyleSemiBold         FontStyle = "SemiBold"
	FontStyleBold             FontStyle = "Bold"
	FontStyleExtraBold        FontStyle = "ExtraBold"
	FontStyleBlack            FontStyle = "Black"
	FontStyleThinItalic       FontStyle = "ThinItalic"
	FontStyleExtraLightItalic FontStyle = "ExtraLightItalic"
	FontStyleLightItalic      FontStyle = "LightItalic"
	FontStyleRegularItalic    FontStyle = "RegularItalic"
	FontStyleMediumItalic     FontStyle = "MediumItalic"
	FontStyleSemiBoldItalic   FontStyle = "SemiBoldItalic"
	FontStyleBoldItalic       FontStyle = "BoldItalic"
	FontStyleExtraBoldItalic  FontStyle = "ExtraBoldItalic"
	FontStyleBlackItalic      FontStyle = "BlackItalic"
)

const DefaultFontStyle = FontStyleRegular

// Font Family

const (
	FontFamilyCourier   = "Courier"
	FontFamilyHelvetica = "Helvetica"
	FontFamilyArial     = "Arial"
	FontFamilyTime      = "Times"
	FontFamilySymbol    = "Symbol"
)

const DefaultFontFamily = FontFamilyHelvetica

// Orientation
type Orientation string

const (
	OrientationPortrait  Orientation = "P"
	OrientationLandscape Orientation = "L"
)

const DefaultOrientation = OrientationPortrait

// Unit
type Unit string

const (
	UnitPt Unit = "pt"
	UnitMm Unit = "mm"
	UnitCm Unit = "cm"
	UnitIn Unit = "in"
)

const DefaultUnit = UnitPt

// Size
type Size string

const (
	SizeA3      Size = "a3"
	SizeA4      Size = "a4"
	SizeA5      Size = "a5"
	SizeA6      Size = "a6"
	SizeA2      Size = "a2"
	SizeA1      Size = "a1"
	SizeLetter  Size = "letter"
	SizeLegal   Size = "legal"
	SizeTabloid Size = "tabloid"
)

const DefaultSize = SizeA4

// FlexDirection
type FlexDirection int

const (
	FlexDirectionColumn FlexDirection = iota
	FlexDirectionColumnReverse
	FlexDirectionRow
	FlexDirectionRowReverse
)

const DefaultFlexDirection = FlexDirectionRow

// Edge
type Edge int

const (
	EdgeLeft       Edge = Edge(flex.EdgeLeft)
	EdgeTop        Edge = Edge(flex.EdgeTop)
	EdgeRight      Edge = Edge(flex.EdgeRight)
	EdgeBottom     Edge = Edge(flex.EdgeBottom)
	EdgeStart      Edge = Edge(flex.EdgeStart)
	EdgeEnd        Edge = Edge(flex.EdgeEnd)
	EdgeHorizontal Edge = Edge(flex.EdgeHorizontal)
	EdgeVertical   Edge = Edge(flex.EdgeVertical)
	EdgeAll        Edge = Edge(flex.EdgeAll)
)

// Align
type Align int

const (
	AlignAuto Align = iota
	AlignStart
	AlignCenter
	AlignEnd
	AlignStretch
	AlignBaseline
	AlignSpaceBetween
	AlignSpaceAround
)

// Justify
type Justify int

const (
	JustifyStart Justify = iota
	JustifyCenter
	JustifyEnd
	JustifySpaceBetween
	JustifySpaceAround
)

// Direction
type Direction int

const (
	DirectionInherit Direction = iota
	DirectionLTR
	DirectionRTL
)

// Wrap
type Wrap int

const (
	WrapNoWrap Wrap = iota
	WrapWrap
	WrapWrapReverse
)

// PositionType
type PositionType int

const (
	PositionTypeRelative PositionType = iota
	PositionTypeAbsolute
)

// Overflow
type Overflow int

const (
	OverflowVisible Overflow = iota
	OverflowHidden
	OverflowScroll
)

// PageOverflow
type PageOverflow int

const (
	PageOverflowNew PageOverflow = iota
	PageOverflowHidden
)

// Display
type Display int

const (
	DisplayFlex Display = iota
	DisplayNone
)
