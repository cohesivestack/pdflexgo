package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type ImageElement struct {
	abstractElement
	filePath string
}

func Image() *ImageElement {
	config := flex.NewConfig()
	config.UseWebDefaults = true
	node := flex.NewNodeWithConfig(config)

	image := &ImageElement{}

	image.abstractElement.flexNode = node

	return image
}

func (image *ImageElement) FromFile(filePath string) *ImageElement {

	image.filePath = filePath

	return image
}

func (image *ImageElement) markRequiredAsDirty() {

}

func (image *ImageElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {
	var measureFunc = func(node *flex.Node, width float32, widthMode flex.MeasureMode, height float32, heightMode flex.MeasureMode) flex.Size {

		fpdf.SetXY(0, 0)
		fpdf.Image(image.filePath, 0, 0, float64(width), 0, true, "", 0, "")
		newHeight := fpdf.GetY()

		return flex.Size{Width: width, Height: float32(newHeight)}
	}

	image.getFlexNode().SetMeasureFunc(measureFunc)
}

func (image *ImageElement) render(pdf *Pdf) {
	image.renderElement(pdf)

	fpdf := pdf.fpdf

	if image.filePath != "" {
		fpdf.Image(
			image.filePath,
			float64(image.x()),
			float64(image.y()),
			float64(image.getFlexNode().LayoutGetWidth()),
			0,
			false, "", 0, "")
	}

}
