package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
)

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
	pdf.pages = append(pdf.pages, pages...)
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
