package pdflexgo

import (
	"github.com/kjk/flex"
)

type abstractElement struct {
	flexNode *flex.Node

	backgroundColor rgba
	borderColor     [edgeCount]rgba
}

func (element *abstractElement) initialize() {
	config := flex.NewConfig()
	config.UseWebDefaults = true
	element.flexNode = flex.NewNodeWithConfig(config)
	element.borderColor = [edgeCount]rgba{{alpha: 1}, {alpha: 1}, {alpha: 1}, {alpha: 1}}
}

func (element *abstractElement) getFlexNode() *flex.Node {
	return element.flexNode
}

func (element *abstractElement) x() float32 {
	return getXParent(element.getFlexNode())
}

func (element *abstractElement) y() float32 {
	return getYParent(element.getFlexNode())
}

func getYParent(flexNode *flex.Node) float32 {
	if flexNode.Parent != nil {
		return flexNode.LayoutGetTop() + getYParent(flexNode.Parent)
	}
	return flexNode.LayoutGetTop()
}

func getXParent(flexNode *flex.Node) float32 {
	if flexNode.Parent != nil {
		return flexNode.LayoutGetLeft() + getXParent(flexNode.Parent)
	}
	return flexNode.LayoutGetLeft()
}

func (element *abstractElement) renderElement(pdf *Pdf) {
	element.renderBorders(pdf).renderBackground(pdf)
}

func (element *abstractElement) renderBorders(pdf *Pdf) *abstractElement {

	element.
		renderBorderTop(pdf).
		renderBorderRight(pdf).
		renderBorderBottom(pdf).
		renderBorderLeft(pdf)

	return element
}

func (element *abstractElement) renderBorderTop(pdf *Pdf) *abstractElement {

	if element.getFlexNode().LayoutGetBorder(flex.EdgeTop) > 0 {
		pdf.fpdf.SetDrawColor(element.borderColor[EdgeTop].red, element.borderColor[EdgeTop].green, element.borderColor[EdgeTop].blue)
		pdf.fpdf.SetAlpha(element.borderColor[EdgeTop].alpha, "")

		width := float64(element.getFlexNode().LayoutGetBorder(flex.EdgeTop))
		pdf.fpdf.SetLineWidth(width)
		pdf.fpdf.Line(
			float64(element.x()),
			float64(element.y())+(width/2),
			float64(element.x()+element.getFlexNode().LayoutGetWidth()),
			float64(element.y())+(width/2))

		pdf.fpdf.SetAlpha(1.0, "")
	}

	return element
}

func (element *abstractElement) renderBorderRight(pdf *Pdf) *abstractElement {

	if element.getFlexNode().LayoutGetBorder(flex.EdgeRight) > 0 {
		pdf.fpdf.SetDrawColor(element.borderColor[EdgeRight].red, element.borderColor[EdgeRight].green, element.borderColor[EdgeRight].blue)
		pdf.fpdf.SetAlpha(element.borderColor[EdgeRight].alpha, "")

		width := float64(element.getFlexNode().LayoutGetBorder(flex.EdgeRight))
		pdf.fpdf.SetLineWidth(width)

		pdf.fpdf.Line(
			float64(element.x()+element.getFlexNode().LayoutGetWidth())-(width/2),
			float64(element.y()),
			float64(element.x()+element.getFlexNode().LayoutGetWidth())-(width/2),
			float64(element.y()+element.getFlexNode().LayoutGetHeight()))

		pdf.fpdf.SetAlpha(1.0, "")
	}

	return element
}

func (element *abstractElement) renderBorderBottom(pdf *Pdf) *abstractElement {

	if element.getFlexNode().LayoutGetBorder(flex.EdgeBottom) > 0 {
		pdf.fpdf.SetDrawColor(element.borderColor[EdgeBottom].red, element.borderColor[EdgeBottom].green, element.borderColor[EdgeBottom].blue)
		pdf.fpdf.SetAlpha(element.borderColor[EdgeBottom].alpha, "")

		width := float64(element.getFlexNode().LayoutGetBorder(flex.EdgeBottom))
		pdf.fpdf.SetLineWidth(width)
		pdf.fpdf.Line(
			float64(element.x()),
			float64(element.y()+element.getFlexNode().LayoutGetHeight())-(width/2),
			float64(element.x()+element.getFlexNode().LayoutGetWidth()),
			float64(element.y()+element.getFlexNode().LayoutGetHeight())-(width/2))

		pdf.fpdf.SetAlpha(1.0, "")
	}

	return element
}

func (element *abstractElement) renderBorderLeft(pdf *Pdf) *abstractElement {

	if element.getFlexNode().LayoutGetBorder(flex.EdgeLeft) > 0 {
		pdf.fpdf.SetDrawColor(element.borderColor[EdgeLeft].red, element.borderColor[EdgeLeft].green, element.borderColor[EdgeLeft].blue)
		pdf.fpdf.SetAlpha(element.borderColor[EdgeLeft].alpha, "")

		width := float64(element.getFlexNode().LayoutGetBorder(flex.EdgeLeft))
		pdf.fpdf.SetLineWidth(width)
		pdf.fpdf.Line(
			float64(element.x())+(width/2),
			float64(element.y()),
			float64(element.x())+(width/2),
			float64(element.y()+element.getFlexNode().LayoutGetHeight()))

		pdf.fpdf.SetAlpha(1.0, "")
	}

	return element
}

func (element *abstractElement) renderBackground(pdf *Pdf) *abstractElement {

	if !equalColor(element.backgroundColor, defaultBackgroundColor) {
		pdf.fpdf.SetFillColor(element.backgroundColor.red, element.backgroundColor.green, element.backgroundColor.blue)
		pdf.fpdf.SetAlpha(element.backgroundColor.alpha, "")
		pdf.fpdf.Rect(
			float64(element.x()+element.getFlexNode().LayoutGetBorder(flex.EdgeLeft)),
			float64(element.y()+element.getFlexNode().LayoutGetBorder(flex.EdgeTop)),
			float64(element.getFlexNode().LayoutGetWidth()-(element.getFlexNode().LayoutGetBorder(flex.EdgeLeft)+element.getFlexNode().LayoutGetBorder(flex.EdgeRight))),
			float64(element.getFlexNode().LayoutGetHeight()-(element.getFlexNode().LayoutGetBorder(flex.EdgeTop)+element.getFlexNode().LayoutGetBorder(flex.EdgeBottom))), "F")
		pdf.fpdf.SetAlpha(1.0, "")
	}

	return element
}
