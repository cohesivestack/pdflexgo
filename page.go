package pdflexgo

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Page struct {
	root *Block

	// Public attributes
	orientation Orientation
	unit        Unit
	width       float64
	height      float64
}

func NewPage() *Page {
	page := &Page{
		root: NewBlock(),
	}

	page.width, page.height = sizeToWidthHeight(DefaultSize)

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

func (page *Page) Children(children ...Element) *Page {
	page.root.Children(children...)

	return page
}

func (page *Page) render(pdf *Pdf) {
	pdf.fpdf.AddPageFormat(
		string(page.orientation),
		gofpdf.SizeType{Wd: page.width, Ht: page.height})

	flex.CalculateLayout(page.root.getFlexNode(), flex.Undefined, flex.Undefined, flex.DirectionLTR)

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
