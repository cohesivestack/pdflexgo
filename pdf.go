package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Orientation string
type Unit string
type Size string
type Edge int
type FlexDirection int

const (
	FlexDirectionColumn FlexDirection = iota
	FlexDirectionColumnReverse
	FlexDirectionRow
	FlexDirectionRowReverse
)

const (
	OrientationPortrait  Orientation = "P"
	OrientationLandscape Orientation = "L"

	UnitPt Unit = "pt"
	UnitMm Unit = "mm"
	UnitCm Unit = "cm"
	UnitIn Unit = "in"

	SizeA3      Size = "a3"
	SizeA4      Size = "a4"
	SizeA5      Size = "a5"
	SizeA6      Size = "a6"
	SizeA2      Size = "a2"
	SizeA1      Size = "a1"
	SizeLetter  Size = "letter"
	SizeLegal   Size = "legal"
	SizeTabloid Size = "tabloid"

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

const DefaultOrientation = OrientationPortrait
const DefaultUnit = UnitPt
const DefaultSize = SizeA4
const DefaultMarginInPoints = 72
const DefaultFlexDirection = FlexDirectionRow

type Pdf struct {
	fpdf  *gofpdf.Fpdf
	pages []*Page
}

func NewPdf() *Pdf {

	pdf := &Pdf{
		fpdf: gofpdf.New(string(DefaultOrientation), string(DefaultUnit), string(DefaultSize), ""),
	}

	return pdf
}

func (pdf *Pdf) Pages(pages ...*Page) *Pdf {
	for _, page := range pages {
		pdf.pages = append(pdf.pages, page)
	}
	return pdf
}

func (pdf *Pdf) Render() *Pdf {

	for _, page := range pdf.pages {
		page.render(pdf)
	}

	return pdf
}

func (pdf *Pdf) OutputFileAndClose(filePath string) error {
	return pdf.fpdf.OutputFileAndClose(filePath)
}
