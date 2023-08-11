package pdflexgo

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type textMultiFormatPart struct {
	content    string
	size       float64
	color      string
	fontStyle  FontStyle
	fontFamily string
}

type TextMultiFormatElement struct {
	AbstractElement

	lineHeight *float64

	size       float64
	color      string
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
	if element.color == "" {
		element.color = defaultProps.fontColor
	}
	if element.size == -1 {
		element.size = defaultProps.fontSize
	}

	for _, part := range element.parts {
		if part.fontFamily == "" {
			part.fontFamily = defaultProps.fontFamily
		}
		if part.fontStyle == "" {
			part.fontStyle = defaultProps.fontStyle
		}
		if part.color == "" {
			part.color = defaultProps.fontColor
		}
		if part.size == -1 {
			part.size = defaultProps.fontSize
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

		if element._flexNode.Parent.StyleGetFlexShrink() == 0 && element._flexNode.Parent.StyleGetFlexGrow() == 0 && element._flexNode.Parent.Style.FlexBasis.Value != element._flexNode.Parent.Style.FlexBasis.Value && fpdf.GetY() == 0 && fpdf.GetX() < float64(width) {
			width = float32(fpdf.GetX() + 7)
		}

		newHeight := fpdf.GetY() + *element.lineHeight

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	element.getFlexNode().SetMeasureFunc(measureFunc)
}

func (element *TextMultiFormatElement) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	fpdf.SetXY(
		float64(element.X()),
		float64(element.Y()))

	pageWidth, _ := fpdf.GetPageSize()
	marginRight := pageWidth - float64(element.X()+element._flexNode.LayoutGetWidth())
	if marginRight < 0 {
		marginRight = 0
	}
	fpdf.SetMargins(float64(element.X()), float64(element.Y()), marginRight)

	for _, part := range element.parts {
		r, g, b, err := hexToRGB(part.color)
		if err != nil {
			log.Fatal(err)
		}
		fpdf.SetTextColor(r, g, b)
		setFont(fpdf, part.fontFamily, part.fontStyle, part.size)
		fpdf.Write(*element.lineHeight, part.content)
	}
}

func (element *TextMultiFormatElement) addPart() {
	element.parts = append(element.parts, &textMultiFormatPart{})
}

func (element *TextMultiFormatElement) Size(size float64) *TextMultiFormatElement {
	element.parts[len(element.parts)-1].size = size
	return element
}

func (element *TextMultiFormatElement) Color(color string) *TextMultiFormatElement {
	element.parts[len(element.parts)-1].color = color
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
