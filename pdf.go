package pdflexgo

import (
	"io"

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

func (pdf *Pdf) Output(writer io.Writer) error {
	return pdf.fpdf.Output(writer)
}

func (pdf *Pdf) AddFont(family string, style FontStyle, filePath string) *Pdf {

	_family, _style := getFontFamilyAndStyle(family, style)

	pdf.fpdf.AddUTF8Font(_family, _style, filePath)
	return pdf
}

func (pdf *Pdf) AddFontFromBytes(family string, style FontStyle, bytes []byte) *Pdf {
	_family, _style := getFontFamilyAndStyle(family, style)

	pdf.fpdf.AddUTF8FontFromBytes(_family, _style, bytes)
	return pdf
}
