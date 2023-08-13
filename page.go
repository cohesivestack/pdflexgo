package pdflexgo

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Page struct {
	root *BlockElement

	// Public attributes
	orientation Orientation
	unit        Unit
	width       float64
	height      float64
}

func NewPage() *Page {
	page := &Page{
		root: Block(),
	}

	page.Size(DefaultSize)
	page.FlexDirection(FlexDirectionColumn)

	return page
}

func (page *Page) Size(size Size) *Page {
	page.width, page.height = sizeToWidthHeight(size)

	return page
}

func (page *Page) Orientation(orientation Orientation) *Page {
	page.orientation = orientation
	return page
}

func (page *Page) Unit(unit Unit) *Page {
	page.unit = unit
	return page
}

func (page *Page) Width(width float64) *Page {
	page.width = width
	page.root.Width(width)
	return page
}

func (page *Page) Height(height float64) *Page {
	page.height = height
	page.root.Height(height)
	return page
}

func (page *Page) FlexDirection(direction FlexDirection) *Page {
	page.root.FlexDirection(direction)
	return page
}

func (page *Page) JustifyContent(justify Justify) *Page {
	page.root.JustifyContent(justify)
	return page
}

func (page *Page) BackgroundColor(color string) *Page {
	page.root.backgrondColor = color
	return page
}

func (page *Page) Children(children ...Node) *Page {
	page.root.Children(children...)

	return page
}

func (page *Page) render(pdf *Pdf) {

	initializeFpdf := func(fpdf *gofpdf.Fpdf) {
		fpdf.AddPageFormat(
			string(page.orientation),
			gofpdf.SizeType{Wd: page.width, Ht: page.height})
		fpdf.SetFont("Arial", "", DefaultFontSize)
	}

	fpdf := pdf.fpdf

	initializeFpdf(fpdf)

	page.root.Width(page.width - (page.root.GetMarginLeft() + page.root.GetMarginRight()))
	page.root.Height(page.height - (page.root.GetMarginTop() + page.root.GetMarginBottom()))

	// Prerender is used to calculate the size of the elements with the
	// CalculateLayout process
	fpdfTemp := gofpdf.New(string(DefaultOrientation), string(DefaultUnit), string(DefaultSize), "")
	initializeFpdf(fpdfTemp)

	for _, fontLoaded := range pdf.fontsLoaded {
		_family, _style := getFontFamilyAndStyle(fontLoaded.fontFamily, fontLoaded.style)
		fpdfTemp.AddUTF8Font(_family, _style, fontLoaded.filePath)
	}

	page.root.preRender(pdf.defaultProps, fpdfTemp)

	// Calculate Flex nodes
	flex.CalculateLayout(page.root.getFlexNode(), float32(page.width), float32(page.height), flex.DirectionLTR)

	page.root.markRequiredAsDirty()

	flex.CalculateLayout(page.root.getFlexNode(), float32(page.width), float32(page.height), flex.DirectionLTR)

	page.root.render(pdf)
}

func sizeToWidthHeight(size Size) (float64, float64) {

	switch size {
	case SizeA3:
		return 841.89, 1190.55
	case SizeA4:
		return 595.28, 841.89
	case SizeA5:
		return 420.94, 595.28
	case SizeA6:
		return 297.64, 420.94
	case SizeA2:
		return 1190.55, 1683.78
	case SizeA1:
		return 1683.78, 2383.94
	case SizeLetter:
		return 612, 792
	case SizeLegal:
		return 612, 1008
	case SizeTabloid:
		return 792, 1224
	}

	log.Fatal(fmt.Printf("standard size %s is not supported", size))

	return 0, 0
}
