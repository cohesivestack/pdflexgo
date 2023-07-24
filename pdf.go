package pdflexgo

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Pdf struct {
	fpdf *gofpdf.Fpdf
	root *Block
}

func NewPdf() *Pdf {
	pdf := &Pdf{
		fpdf: gofpdf.New("P", "mm", "A4", ""),
	}

	pdf.root = NewBlock()
	pdf.root.setPdf(pdf)

	return pdf
}

func (pdf *Pdf) Children(children ...*Block) *Pdf {

	for _, block := range children {
		block.setPdf(pdf)
	}

	pdf.root.Children(children...)

	return pdf
}

func (pdf *Pdf) Build() *Pdf {

	pdf.fpdf.AddPage()

	pdf.root.flexNode.StyleSetMargin(flex.EdgeAll, 10)

	flex.CalculateLayout(pdf.root.flexNode, flex.Undefined, flex.Undefined, flex.DirectionLTR)

	pdf.root.build()

	err := pdf.fpdf.OutputFileAndClose("tmp/hello.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	return pdf
}
