package pdflexgo

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type TextElement struct {
	AbstractElement

	lineHeight *float64

	content    string
	size       float64
	color      string
	fontStyle  FontStyle
	fontFamily string
}

func (text *TextElement) Content(content string) *TextElement {
	text.content = content
	return text
}

func (text *TextElement) Size(size float64) *TextElement {
	text.size = size
	return text
}

func (text *TextElement) Color(color string) *TextElement {
	text.color = color
	return text
}

func (text *TextElement) Thin() *TextElement {
	text.fontStyle = FontStyleThin
	return text
}
func (text *TextElement) ExtraLight() *TextElement {
	text.fontStyle = FontStyleExtraLight
	return text
}
func (text *TextElement) Light() *TextElement {
	text.fontStyle = FontStyleLight
	return text
}
func (text *TextElement) Regular() *TextElement {
	text.fontStyle = FontStyleRegular
	return text
}
func (text *TextElement) Medium() *TextElement {
	text.fontStyle = FontStyleMedium
	return text
}
func (text *TextElement) SemiBold() *TextElement {
	text.fontStyle = FontStyleSemiBold
	return text
}
func (text *TextElement) Bold() *TextElement {
	text.fontStyle = FontStyleBold
	return text
}
func (text *TextElement) ExtraBold() *TextElement {
	text.fontStyle = FontStyleExtraBold
	return text
}
func (text *TextElement) Black() *TextElement {
	text.fontStyle = FontStyleBlack
	return text
}
func (text *TextElement) ThinItalic() *TextElement {
	text.fontStyle = FontStyleThinItalic
	return text
}
func (text *TextElement) ExtraLightItalic() *TextElement {
	text.fontStyle = FontStyleExtraLightItalic
	return text
}
func (text *TextElement) LightItalic() *TextElement {
	text.fontStyle = FontStyleLightItalic
	return text
}
func (text *TextElement) RegularItalic() *TextElement {
	text.fontStyle = FontStyleRegularItalic
	return text
}
func (text *TextElement) MediumItalic() *TextElement {
	text.fontStyle = FontStyleMediumItalic
	return text
}
func (text *TextElement) SemiBoldItalic() *TextElement {
	text.fontStyle = FontStyleSemiBoldItalic
	return text
}
func (text *TextElement) BoldItalic() *TextElement {
	text.fontStyle = FontStyleBoldItalic
	return text
}
func (text *TextElement) ExtraBoldItalic() *TextElement {
	text.fontStyle = FontStyleExtraBoldItalic
	return text
}
func (text *TextElement) BlackItalic() *TextElement {
	text.fontStyle = FontStyleBlackItalic
	return text
}

func (text *TextElement) LineHeight(height float64) *TextElement {
	text.lineHeight = &height
	return text
}

func (text *TextElement) FontFamily(family string) *TextElement {
	text.fontFamily = family
	return text
}

func Text() *TextElement {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	text := &TextElement{
		size: -1,
	}

	text.AbstractElement.setFlexNode(node)
	text._flexNode.StyleSetMargin(flex.EdgeAll, 0)
	text._flexNode.StyleSetPadding(flex.EdgeAll, 0)
	text._flexNode.StyleSetHeightAuto()
	text._flexNode.StyleSetWidthAuto()

	return text
}

func (text *TextElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	if text.fontFamily == "" {
		text.fontFamily = defaultProps.fontFamily
	}
	if text.fontStyle == "" {
		text.fontStyle = defaultProps.fontStyle
	}
	if text.color == "" {
		text.color = defaultProps.fontColor
	}
	if text.size == -1 {
		text.size = defaultProps.fontSize
	}

	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {
		setFont(fpdf, text.fontFamily, text.fontStyle, text.size)

		_, fontSize := fpdf.GetFontSize()

		if text.lineHeight == nil {
			text.lineHeight = &fontSize
		}

		fpdf.SetXY(0, 0)
		pageWidth, _ := fpdf.GetPageSize()
		marginRight := pageWidth - float64(width)
		if marginRight < 0 {
			marginRight = 0
		}
		fpdf.SetMargins(0, 0, marginRight)
		fpdf.Write(fontSize, text.content)
		newHeight := fpdf.GetY() + fontSize

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	text.getFlexNode().SetMeasureFunc(measureFunc)
}

func (text *TextElement) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	r, g, b, err := hexToRGB(text.color)
	if err != nil {
		log.Fatal(err)
	}
	fpdf.SetTextColor(r, g, b)

	setFont(fpdf, text.fontFamily, text.fontStyle, text.size)

	_, fontSize := fpdf.GetFontSize()

	fpdf.SetXY(
		float64(text.X()),
		float64(text.Y()))
	pageWidth, _ := fpdf.GetPageSize()
	marginRight := pageWidth - float64(text.X()+text._flexNode.LayoutGetWidth())
	if marginRight < 0 {
		marginRight = 0
	}
	fpdf.SetMargins(float64(text.X()), float64(text.Y()), marginRight)
	fpdf.Write(fontSize, text.content)
}
