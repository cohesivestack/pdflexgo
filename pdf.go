package pdflexgo

//go:generate go run generator/main.go

import (
	"io"

	"github.com/jung-kurt/gofpdf"
)

type defaultProps struct {
	fontFamily string
	fontStyle  FontStyle
	fontSize   float64
	fontColor  rgba
}

type Pdf struct {
	fpdf  *gofpdf.Fpdf
	pages []*PageElement

	fontsLoaded  []FontLoadInformation
	defaultProps *defaultProps

	header HeaderBuilder
	footer FooterBuilder
}

func NewPdf() *Pdf {

	pdf := &Pdf{
		fpdf: gofpdf.New(string(DefaultOrientation), string(DefaultUnit), string(DefaultSize), ""),

		defaultProps: &defaultProps{
			fontStyle:  DefaultFontStyle,
			fontFamily: DefaultFontFamily,
			fontSize:   DefaultFontSize,
			fontColor:  DefaultFontColor,
		},
	}

	pdf.fpdf.SetAutoPageBreak(false, 0)

	return pdf
}

func (pdf *Pdf) Header(header HeaderBuilder) *Pdf {
	pdf.header = header
	return pdf
}

func (pdf *Pdf) Footer(footer FooterBuilder) *Pdf {
	pdf.footer = footer
	return pdf
}

func (pdf *Pdf) Pages(pages ...*PageElement) *Pdf {
	pdf.pages = append(pdf.pages, pages...)
	return pdf
}

func renderPageRecursive(page *PageElement, pdf *Pdf, pageNumber int, overflowedContinuation bool) int {
	nextPageOverflowed := page.render(pdf, pageNumber, overflowedContinuation)
	if nextPageOverflowed != nil {
		// Recursive call with overflowed nodes and incremented page number
		return renderPageRecursive(nextPageOverflowed, pdf, pageNumber+1, true)
	}
	return pageNumber + 1
}

func (pdf *Pdf) Render() *Pdf {

	pageNumber := 1
	for _, page := range pdf.pages {
		pageNumber = renderPageRecursive(page, pdf, pageNumber, false)
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

	pdf.fontsLoaded = append(pdf.fontsLoaded, FontLoadInformation{
		fontFamily: family,
		style:      style,
		filePath:   filePath,
	})

	_family, _style := getFontFamilyAndStyle(family, style)

	pdf.fpdf.AddUTF8Font(_family, _style, filePath)
	return pdf
}

func (pdf *Pdf) AddFontFromBytes(family string, style FontStyle, bytes []byte) *Pdf {
	_family, _style := getFontFamilyAndStyle(family, style)

	pdf.fpdf.AddUTF8FontFromBytes(_family, _style, bytes)
	return pdf
}

func (pdf *Pdf) FontThin() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleThin
	return pdf
}
func (pdf *Pdf) FontExtraLight() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleExtraLight
	return pdf
}
func (pdf *Pdf) FontLight() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleLight
	return pdf
}
func (pdf *Pdf) FontRegular() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleRegular
	return pdf
}
func (pdf *Pdf) FontMedium() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleMedium
	return pdf
}
func (pdf *Pdf) FontSemiBold() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleSemiBold
	return pdf
}
func (pdf *Pdf) FontBold() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleBold
	return pdf
}
func (pdf *Pdf) FontExtraBold() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleExtraBold
	return pdf
}
func (pdf *Pdf) FontBlack() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleBlack
	return pdf
}
func (pdf *Pdf) FontThinItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleThinItalic
	return pdf
}
func (pdf *Pdf) FontExtraLightItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleExtraLightItalic
	return pdf
}
func (pdf *Pdf) FontLightItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleLightItalic
	return pdf
}
func (pdf *Pdf) FontRegularItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleRegularItalic
	return pdf
}
func (pdf *Pdf) FontMediumItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleMediumItalic
	return pdf
}
func (pdf *Pdf) FontSemiBoldItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleSemiBoldItalic
	return pdf
}
func (pdf *Pdf) FontBoldItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleBoldItalic
	return pdf
}
func (pdf *Pdf) FontExtraBoldItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleExtraBoldItalic
	return pdf
}
func (pdf *Pdf) FontBlackItalic() *Pdf {
	pdf.defaultProps.fontStyle = FontStyleBlackItalic
	return pdf
}

func (pdf *Pdf) FontFamily(family string) *Pdf {
	pdf.defaultProps.fontFamily = family
	return pdf
}

func (pdf *Pdf) FontSize(size float64) *Pdf {
	pdf.defaultProps.fontSize = size
	return pdf
}

func (pdf *Pdf) FontColor(red int, green int, blue int, alpha ...float64) *Pdf {
	pdf.defaultProps.fontColor = getRgba(red, green, blue, alpha...)
	return pdf
}
