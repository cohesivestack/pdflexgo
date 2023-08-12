package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Element interface {
	getFlexNode() *flex.Node
	markRequiredAsDirty()
	render(*Pdf)
	preRender(*defaultProps, *gofpdf.Fpdf)
	x() float32
	y() float32
}
