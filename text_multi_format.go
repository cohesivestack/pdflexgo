package pdflexgo

import (
	"github.com/kjk/flex"
)

type textMultiFormatPart struct {
	content string
	size    float64
}

type TextMultiFormatElement struct {
	AbstractElement

	lineHeight *float64
	parts      []*textMultiFormatPart
}

type TextMultiFormatCreator struct {
	lineHeight *float64
}

func TextMultiFormat() *TextMultiFormatCreator {
	return &TextMultiFormatCreator{}
}

func (creator *TextMultiFormatCreator) LineHeight(height float64) *TextMultiFormatCreator {
	creator.lineHeight = &height
	return creator
}

func (constructor *TextMultiFormatCreator) Create() *TextMultiFormatElement {

	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	element := &TextMultiFormatElement{
		lineHeight: constructor.lineHeight,
		parts: []*textMultiFormatPart{
			{
				size: DefaultFontSize,
			},
		},
	}

	element.AbstractElement.setFlexNode(node)
	element._flexNode.StyleSetMargin(flex.EdgeAll, 0)
	element._flexNode.StyleSetPadding(flex.EdgeAll, 0)
	element._flexNode.StyleSetHeightAuto()
	element._flexNode.StyleSetWidthAuto()

	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {
		fpdf := element.preRenderPdf

		fpdf.SetXY(0, 0)
		pageWidth, _ := fpdf.GetPageSize()
		fpdf.SetMargins(0, 0, pageWidth-float64(width))

		if element.lineHeight == nil {
			lineHeight := 0.0
			element.lineHeight = &lineHeight
			for _, part := range element.parts {
				fpdf.SetFontSize(float64(part.size))
				_, fontHeight := fpdf.GetFontSize()
				if fontHeight > *element.lineHeight {
					element.lineHeight = &fontHeight
				}
			}
		}
		for _, part := range element.parts {
			fpdf.SetFontSize(float64(part.size))
			fpdf.Write(*element.lineHeight, part.content)
		}

		newHeight := fpdf.GetY() + *element.lineHeight

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	node.SetMeasureFunc(measureFunc)

	return element
}

func (element *TextMultiFormatElement) Size(size float64) *TextMultiFormatElement {
	element.parts[len(element.parts)-1].size = size
	return element
}

func (element *TextMultiFormatElement) Content(text string) *TextMultiFormatElement {
	part := element.parts[len(element.parts)-1]
	part.content = text

	element.parts = append(element.parts, &textMultiFormatPart{
		size: part.size,
	})
	return element
}

func (element *TextMultiFormatElement) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	fpdf.SetXY(
		float64(element.X()),
		float64(element.Y()))

	pageWidth, _ := fpdf.GetPageSize()
	fpdf.SetMargins(float64(element.X()), float64(element.Y()), pageWidth-float64(element.X()+element._flexNode.LayoutGetWidth()))

	for _, part := range element.parts {
		fpdf.SetFontSize(part.size)
		fpdf.Write(*element.lineHeight, part.content)
	}
}
