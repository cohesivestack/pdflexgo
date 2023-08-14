package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type textMultiFormatPart struct {
	content    string
	size       float64
	color      rgba
	fontStyle  FontStyle
	fontFamily string
}

type TextMultiFormatElement struct {
	abstractElement

	lineHeight *float64

	size       float64
	color      rgba
	fontStyle  FontStyle
	fontFamily string

	parts []*textMultiFormatPart
}

func (element *TextMultiFormatElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {
	if element.fontFamily == "" {
		element.fontFamily = defaultProps.fontFamily
	}
	if element.fontStyle == "" {
		element.fontStyle = defaultProps.fontStyle
	}
	if element.size == -1 {
		element.size = defaultProps.fontSize
	}
	if equalColor(element.color, DefaultFontColor) {
		element.color = defaultProps.fontColor
	}

	for _, part := range element.parts {
		if part.fontFamily == "" {
			part.fontFamily = defaultProps.fontFamily
		}
		if part.fontStyle == "" {
			part.fontStyle = defaultProps.fontStyle
		}
		if part.size == -1 {
			part.size = defaultProps.fontSize
		}
		if equalColor(part.color, DefaultFontColor) {
			part.color = defaultProps.fontColor
		}
	}

	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {

		fpdf.SetXY(0, 0)
		pageWidth, _ := fpdf.GetPageSize()
		marginRight := pageWidth - float64(width)
		if marginRight < 0 {
			marginRight = 0
		}
		fpdf.SetMargins(0, 0, marginRight)

		if element.lineHeight == nil {
			lineHeight := 0.0
			element.lineHeight = &lineHeight
			for _, part := range element.parts {
				setFont(fpdf, part.fontFamily, part.fontStyle, part.size)
				_, fontHeight := fpdf.GetFontSize()
				if fontHeight > lineHeight {
					lineHeight = fontHeight
				}
			}
		}

		for i := range element.parts {
			part := element.parts[i]
			setFont(fpdf, part.fontFamily, part.fontStyle, part.size)
			fpdf.Write(*element.lineHeight, part.content)
		}

		newHeight := fpdf.GetY() + *element.lineHeight

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	element.getFlexNode().SetMeasureFunc(measureFunc)
}

func (block *TextMultiFormatElement) markRequiredAsDirty() {
}

func (element *TextMultiFormatElement) render(pdf *Pdf) {
	element.renderElement(pdf)

	fpdf := pdf.fpdf

	fpdf.SetXY(
		float64(element.x()),
		float64(element.y()))

	pageWidth, _ := fpdf.GetPageSize()
	marginRight := pageWidth - float64(element.x()+element.flexNode.LayoutGetWidth())
	if marginRight < 0 {
		marginRight = 0
	}
	fpdf.SetMargins(float64(element.x()), float64(element.y()), marginRight)

	for _, part := range element.parts {

		fpdf.SetTextColor(part.color.red, part.color.green, part.color.blue)
		fpdf.SetAlpha(part.color.alpha, "")

		setFont(fpdf, part.fontFamily, part.fontStyle, part.size)
		fpdf.Write(*element.lineHeight, part.content)
		fpdf.SetAlpha(1, "")
	}
}

func (element *TextMultiFormatElement) addPart() {
	element.parts = append(element.parts, &textMultiFormatPart{})
}

func (element *TextMultiFormatElement) Size(size float64) *TextMultiFormatElement {
	element.parts[len(element.parts)-1].size = size
	return element
}

func (element *TextMultiFormatElement) Thin() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleThin
	return element
}
func (element *TextMultiFormatElement) ExtraLight() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleExtraLight
	return element
}
func (element *TextMultiFormatElement) Light() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleLight
	return element
}
func (element *TextMultiFormatElement) Regular() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleRegular
	return element
}
func (element *TextMultiFormatElement) Medium() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleMedium
	return element
}
func (element *TextMultiFormatElement) SemiBold() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleSemiBold
	return element
}
func (element *TextMultiFormatElement) Bold() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleBold
	return element
}
func (element *TextMultiFormatElement) ExtraBold() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleExtraBold
	return element
}
func (element *TextMultiFormatElement) Black() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleBlack
	return element
}
func (element *TextMultiFormatElement) ThinItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleThinItalic
	return element
}
func (element *TextMultiFormatElement) ExtraLightItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleExtraLightItalic
	return element
}
func (element *TextMultiFormatElement) LightItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleLightItalic
	return element
}
func (element *TextMultiFormatElement) RegularItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleRegularItalic
	return element
}
func (element *TextMultiFormatElement) MediumItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleMediumItalic
	return element
}
func (element *TextMultiFormatElement) SemiBoldItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleSemiBoldItalic
	return element
}
func (element *TextMultiFormatElement) BoldItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleBoldItalic
	return element
}
func (element *TextMultiFormatElement) ExtraBoldItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleExtraBoldItalic
	return element
}
func (element *TextMultiFormatElement) BlackItalic() *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontStyle = FontStyleBlackItalic
	return element
}

func (element *TextMultiFormatElement) FontFamily(family string) *TextMultiFormatElement {
	element.parts[len(element.parts)-1].fontFamily = family
	return element
}

func (element *TextMultiFormatElement) Content(text string) *TextMultiFormatElement {
	part := element.parts[len(element.parts)-1]
	part.content = text

	element.addPart()
	return element
}
