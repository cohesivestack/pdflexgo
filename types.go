package pdflexgo

import "github.com/kjk/flex"

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
type Align string

const (
	AlignStart  Align = "start"
	AlignEnd    Align = "end"
	AlignCenter Align = "center"
)

// Justify
type Justify string

const (
	JustifyStart  Justify = "start"
	JustifyEnd    Justify = "end"
	JustifyCenter Justify = "center"
)

// Direction
type Direction int

const (
	DirectionInherit Direction = iota
	DirectionLTR
	DirectionRTL
)
