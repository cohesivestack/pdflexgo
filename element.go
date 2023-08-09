package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Element interface {
	setFlexNode(*flex.Node)
	getFlexNode() *flex.Node
	render(*Pdf)
	preRender(*defaultProps, *gofpdf.Fpdf)
	setZeroFpdf()
	X() float32
	Y() float32
}

type AbstractElement struct {
	_parent       Element
	preRenderFpdf *gofpdf.Fpdf
	_flexNode     *flex.Node
}

func (elem *AbstractElement) setFlexNode(flexNode *flex.Node) {
	elem._flexNode = flexNode
}

func (elem *AbstractElement) getFlexNode() *flex.Node {
	return elem._flexNode
}

func (elem *AbstractElement) X() float32 {
	return getXParent(elem.getFlexNode())
}

func (elem *AbstractElement) Y() float32 {
	return getYParent(elem.getFlexNode())
}

func (elem *AbstractElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {
	elem.preRenderFpdf = fpdf
}

func (elem *AbstractElement) setZeroFpdf() {
	elem.preRenderFpdf = nil
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
