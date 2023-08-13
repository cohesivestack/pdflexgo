package pdflexgo

import "github.com/kjk/flex"

type abstractFlexElement struct {
	flexNode *flex.Node
}

func (element *abstractFlexElement) getFlexNode() *flex.Node {
	return element.flexNode
}

func (element *abstractFlexElement) x() float32 {
	return getXParent(element.getFlexNode())
}

func (element *abstractFlexElement) y() float32 {
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
