package pdflexgo

import "github.com/kjk/flex"

type Element interface {
	setFlexNode(*flex.Node)
	getFlexNode() *flex.Node
	render(*Pdf)
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

func (elem *AbstractElement) X() float32 {
	x := elem._flexNode.LayoutGetLeft()
	if elem._flexNode.Parent != nil {
		x += elem._flexNode.Parent.LayoutGetLeft()
	}

	return x
}

func (elem *AbstractElement) Y() float32 {
	y := elem._flexNode.LayoutGetTop()
	if elem._flexNode.Parent != nil {
		y += elem._flexNode.Parent.LayoutGetTop()
	}

	return y
}
