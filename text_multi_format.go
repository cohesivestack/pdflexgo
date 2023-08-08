package pdflexgo

import (
	"log"
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

func (element *TextMultiFormatElement) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	fpdf.SetXY(
		float64(element.X()),
		float64(element.Y()))

	pageWidth, _ := fpdf.GetPageSize()
	fpdf.SetMargins(float64(element.X()), float64(element.Y()), pageWidth-float64(element.X()+element._flexNode.LayoutGetWidth()))

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
	element.parts = append(element.parts, &textMultiFormatPart{
		size:       element.size,
		color:      element.color,
		fontFamily: element.fontFamily,
		fontStyle:  element.fontStyle,
	})
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

func (element *TextMultiFormatElement) Content(text string) *TextMultiFormatElement {
	part := element.parts[len(element.parts)-1]
	part.content = text

	element.addPart()
	return element
}
