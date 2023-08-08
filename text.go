package pdflexgo

import (
	"github.com/kjk/flex"
)

type Text struct {
	AbstractElement

	content string
	size    float64
}

func (text *Text) Content(content string) *Text {
	text.content = content
	return text
}

func (text *Text) Size(size float64) *Text {
	text.size = size
	return text
}

func NewText() *Text {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	text := &Text{
		size: DefaultFontSize,
	}

	text.AbstractElement.setFlexNode(node)
	text._flexNode.StyleSetMargin(flex.EdgeAll, 0)
	text._flexNode.StyleSetPadding(flex.EdgeAll, 0)
	text._flexNode.StyleSetHeightAuto()
	text._flexNode.StyleSetWidthAuto()

	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {
		fpdf := text.preRenderPdf

		fpdf.SetFontSize(float64(text.size))
		_, fontSize := fpdf.GetFontSize()
		fpdf.SetXY(0, 0)
		pageWidth, _ := fpdf.GetPageSize()
		fpdf.SetMargins(0, 0, pageWidth-float64(width))
		fpdf.Write(fontSize, text.content)
		newHeight := fpdf.GetY() + fontSize

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	node.SetMeasureFunc(measureFunc)

	return text
}

func (text *Text) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	fpdf.SetFontSize(float64(text.size))
	_, fontSize := fpdf.GetFontSize()
	fpdf.SetXY(
		float64(text.X()),
		float64(text.Y()))
	pageWidth, _ := fpdf.GetPageSize()
	fpdf.SetMargins(float64(text.X()), float64(text.Y()), pageWidth-float64(text.X()+text._flexNode.LayoutGetWidth()))
	fpdf.Write(fontSize, text.content)
}
