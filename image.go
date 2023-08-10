package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type ImageElement struct {
	AbstractElement
	filePath string
}

func Image() *ImageElement {
	config := flex.NewConfig()
	node := flex.NewNodeWithConfig(config)

	image := &ImageElement{}

	image.AbstractElement.setFlexNode(node)
	image._flexNode.StyleSetMargin(flex.EdgeAll, 0)
	image._flexNode.StyleSetPadding(flex.EdgeAll, 0)
	image._flexNode.StyleSetHeightAuto()
	image._flexNode.StyleSetWidthAuto()

	return image
}

func (image *ImageElement) FromFile(filePath string) *ImageElement {

	image.filePath = filePath

	return image
}

func (image *ImageElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {
	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {

		fpdf.SetXY(0, 0)
		fpdf.ImageOptions(image.filePath, 0, 0, float64(width), 0, true, gofpdf.ImageOptions{}, 0, "")
		newHeight := fpdf.GetY()

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	image.getFlexNode().SetMeasureFunc(measureFunc)
}

func (image *ImageElement) render(pdf *Pdf) {
	fpdf := pdf.fpdf

	if image.filePath != "" {
		fpdf.ImageOptions(
			image.filePath,
			float64(image.X()),
			float64(image.Y()),
			float64(image.getFlexNode().LayoutGetWidth()),
			0,
			true, gofpdf.ImageOptions{}, 0, "")
	}

}
