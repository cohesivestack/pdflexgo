package pdflexgo

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type Element interface {
	setFlexNode(*flex.Node)
	getFlexNode() *flex.Node
	markDirty()
	render(*Pdf)
	preRender(*defaultProps, *gofpdf.Fpdf)
	X() float32
	Y() float32
}

type AbstractElement struct {
	_parent   Element
	_flexNode *flex.Node
}

func (elem *AbstractElement) setFlexNode(flexNode *flex.Node) {
	elem._flexNode = flexNode
}

func (elem *AbstractElement) getFlexNode() *flex.Node {
	return elem._flexNode
}

func (elem *AbstractElement) markDirty() {
	//elem._flexNode.MarkDirty()
}

func (elem *AbstractElement) X() float32 {
	return getXParent(elem.getFlexNode())
}

func (elem *AbstractElement) Y() float32 {
	return getYParent(elem.getFlexNode())
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
