package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

var defaultBackgroundColor = rgba{0, 0, 0, 0}

type TextElement struct {
	abstractElement

	lineHeight         float64
	lineHeightAssigned bool
	firstRequestMade   bool

	content         string
	size            float64
	fontStyle       FontStyle
	fontFamily      string
	color           rgba
	backgroundColor rgba
}

func Text() *TextElement {
	text := &TextElement{
		size: -1,
	}

	text.abstractElement.initialize()

	return text
}

func (text *TextElement) Content(content string) *TextElement {
	text.content = content
	return text
}

func (text *TextElement) Size(size float64) *TextElement {
	text.size = size
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
	text.lineHeight = height
	text.lineHeightAssigned = true
	return text
}

func (text *TextElement) FontFamily(family string) *TextElement {
	text.fontFamily = family
	return text
}

func (elem *TextElement) markRequiredAsDirty() {
	elem.flexNode.MarkDirty()
}

func (text *TextElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	if text.fontFamily == "" {
		text.fontFamily = defaultProps.fontFamily
	}
	if text.fontStyle == "" {
		text.fontStyle = defaultProps.fontStyle
	}
	if text.size == -1 {
		text.size = defaultProps.fontSize
	}
	if equalColor(text.color, DefaultFontColor) {
		text.color = defaultProps.fontColor
	}

	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {

		setFont(fpdf, text.fontFamily, text.fontStyle, text.size)

		_, fontSize := fpdf.GetFontSize()

		if !text.lineHeightAssigned {
			text.lineHeight = fontSize
		}

		if !text.firstRequestMade {
			width = float32(fpdf.GetStringWidth(text.content))
			height = float32(text.lineHeight)
			text.firstRequestMade = true
		} else {
			fpdf.SetXY(0, 0)
			pageWidth, _ := fpdf.GetPageSize()
			marginRight := pageWidth - float64(text.flexNode.LayoutGetWidth()-text.flexNode.LayoutGetPadding(flex.EdgeLeft)+text.flexNode.LayoutGetPadding(flex.EdgeRight))
			if marginRight < 0 {
				marginRight = 0
			}
			fpdf.SetMargins(0, 0, marginRight)
			fpdf.SetCellMargin(0)
			fpdf.Write(text.lineHeight, text.content)
			height = float32(fpdf.GetY() + text.lineHeight)
			width = text.flexNode.LayoutGetWidth() - (text.flexNode.LayoutGetPadding(flex.EdgeLeft) + text.flexNode.LayoutGetPadding(flex.EdgeRight))
		}

		return flex.Size{Width: width, Height: float32(height)}
	}

	text.getFlexNode().SetMeasureFunc(measureFunc)
}

func (text *TextElement) render(pdf *Pdf) {
	text.renderElement(pdf)

	fpdf := pdf.fpdf

	setFont(fpdf, text.fontFamily, text.fontStyle, text.size)

	fpdf.SetXY(
		float64(text.x()),
		float64(text.y()))

	if !equalColor(text.backgroundColor, defaultBackgroundColor) {
		pdf.fpdf.SetFillColor(text.backgroundColor.red, text.backgroundColor.green, text.backgroundColor.blue)
		pdf.fpdf.SetAlpha(text.backgroundColor.alpha, "")
		pdf.fpdf.Rect(
			float64(text.x()+text.getFlexNode().LayoutGetBorder(flex.EdgeLeft)),
			float64(text.y()+text.getFlexNode().LayoutGetBorder(flex.EdgeTop)),
			float64(text.getFlexNode().LayoutGetWidth()-(text.getFlexNode().LayoutGetBorder(flex.EdgeLeft)+text.getFlexNode().LayoutGetBorder(flex.EdgeRight))),
			float64(text.getFlexNode().LayoutGetHeight()-(text.getFlexNode().LayoutGetBorder(flex.EdgeTop)+text.getFlexNode().LayoutGetBorder(flex.EdgeBottom))), "F")
		pdf.fpdf.SetAlpha(1.0, "")

	}

	pageWidth, _ := fpdf.GetPageSize()
	marginRight := pageWidth - float64(text.x()+text.flexNode.LayoutGetWidth()+text.flexNode.LayoutGetPadding(flex.EdgeLeft)+text.flexNode.LayoutGetPadding(flex.EdgeRight))
	if marginRight < 0 {
		marginRight = 0
	}
	fpdf.SetCellMargin(0)
	fpdf.SetXY(float64(text.x()+text.flexNode.LayoutGetPadding(flex.EdgeLeft)), float64(text.y()+text.flexNode.LayoutGetPadding(flex.EdgeTop)))
	fpdf.SetMargins(float64(text.x()+text.flexNode.LayoutGetPadding(flex.EdgeLeft)), float64(text.y()+text.flexNode.LayoutGetPadding(flex.EdgeTop)), marginRight)

	fpdf.SetTextColor(text.color.red, text.color.green, text.color.blue)
	fpdf.SetAlpha(text.color.alpha, "")

	fpdf.Write(text.lineHeight, text.content)
	fpdf.SetAlpha(1, "")
}
