package pdflexgo

import "github.com/kjk/flex"

type TextMultiFormatCreator struct {
	element *TextMultiFormatElement
}

func TextMultiFormat() *TextMultiFormatCreator {

	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	element := &TextMultiFormatElement{
		lineHeight: nil,

		size:       DefaultFontSize,
		color:      DefaultFontColor,
		fontStyle:  DefaultFontStyle,
		fontFamily: DefaultFontFamily,

		parts: []*textMultiFormatPart{},
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
				setFont(fpdf, part.fontFamily, part.fontStyle, part.size)
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

	return &TextMultiFormatCreator{
		element: element,
	}
}

func (constructor *TextMultiFormatCreator) Create() *TextMultiFormatElement {
	constructor.element.addPart()
	return constructor.element
}

func (creator *TextMultiFormatCreator) LineHeight(height float64) *TextMultiFormatCreator {
	creator.element.lineHeight = &height
	return creator
}

func (creator *TextMultiFormatCreator) Size(size float64) *TextMultiFormatCreator {
	creator.element.size = size
	return creator
}

func (creator *TextMultiFormatCreator) Color(color string) *TextMultiFormatCreator {
	creator.element.color = color
	return creator
}

func (creator *TextMultiFormatCreator) Thin() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleThin
	return creator
}
func (creator *TextMultiFormatCreator) ExtraLight() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleExtraLight
	return creator
}
func (creator *TextMultiFormatCreator) Light() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleLight
	return creator
}
func (creator *TextMultiFormatCreator) Regular() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleRegular
	return creator
}
func (creator *TextMultiFormatCreator) Medium() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleMedium
	return creator
}
func (creator *TextMultiFormatCreator) SemiBold() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleSemiBold
	return creator
}
func (creator *TextMultiFormatCreator) Bold() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleBold
	return creator
}
func (creator *TextMultiFormatCreator) ExtraBold() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleExtraBold
	return creator
}
func (creator *TextMultiFormatCreator) Black() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleBlack
	return creator
}
func (creator *TextMultiFormatCreator) ThinItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleThinItalic
	return creator
}
func (creator *TextMultiFormatCreator) ExtraLightItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleExtraLightItalic
	return creator
}
func (creator *TextMultiFormatCreator) LightItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleLightItalic
	return creator
}
func (creator *TextMultiFormatCreator) RegularItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleRegularItalic
	return creator
}
func (creator *TextMultiFormatCreator) MediumItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleMediumItalic
	return creator
}
func (creator *TextMultiFormatCreator) SemiBoldItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleSemiBoldItalic
	return creator
}
func (creator *TextMultiFormatCreator) BoldItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleBoldItalic
	return creator
}
func (creator *TextMultiFormatCreator) ExtraBoldItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleExtraBoldItalic
	return creator
}
func (creator *TextMultiFormatCreator) BlackItalic() *TextMultiFormatCreator {
	creator.element.fontStyle = FontStyleBlackItalic
	return creator
}
