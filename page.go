package pdflexgo

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type PageElement struct {
	abstractElement

	children []Node

	// Public attributes
	orientation Orientation
	unit        Unit
}

func Page() *PageElement {
	page := &PageElement{}

	page.abstractElement.initialize()

	page.Size(DefaultSize)
	page.FlexDirection(FlexDirectionColumn)

	return page
}

func (page *PageElement) Size(size Size) *PageElement {
	width, height := sizeToWidthHeight(size)

	page.Width(width)
	page.Height(height)

	return page
}

func (page *PageElement) Orientation(orientation Orientation) *PageElement {
	page.orientation = orientation
	return page
}

func (page *PageElement) Unit(unit Unit) *PageElement {
	page.unit = unit
	return page
}

func (page *PageElement) markRequiredAsDirty() {
	for _, child := range page.children {
		child.markRequiredAsDirty()
	}
}

func (page *PageElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	for _, child := range page.children {
		child.preRender(defaultProps, fpdf)
	}
}

func (page *PageElement) render(pdf *Pdf) {

	initializeFpdf := func(fpdf *gofpdf.Fpdf) {
		fpdf.AddPageFormat(
			string(page.orientation),
			gofpdf.SizeType{
				Wd: float64(page.getFlexNode().LayoutGetWidth()),
				Ht: float64(page.getFlexNode().LayoutGetHeight())})
		fpdf.SetFont("Arial", "", DefaultFontSize)
	}

	fpdf := pdf.fpdf

	initializeFpdf(fpdf)

	// Prerender is used to calculate the size of the elements with the
	// CalculateLayout process
	fpdfTemp := gofpdf.New(string(DefaultOrientation), string(DefaultUnit), string(DefaultSize), "")
	initializeFpdf(fpdfTemp)

	for _, fontLoaded := range pdf.fontsLoaded {
		_family, _style := getFontFamilyAndStyle(fontLoaded.fontFamily, fontLoaded.style)
		fpdfTemp.AddUTF8Font(_family, _style, fontLoaded.filePath)
	}

	page.preRender(pdf.defaultProps, fpdfTemp)

	// Calculate Flex nodes
	flex.CalculateLayout(page.getFlexNode(), flex.Undefined, flex.Undefined, page.getFlexNode().Style.Direction)

	page.markRequiredAsDirty()

	flex.CalculateLayout(page.getFlexNode(), flex.Undefined, flex.Undefined, page.getFlexNode().Style.Direction)

	page.renderElement(pdf)
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
