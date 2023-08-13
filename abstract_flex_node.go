package pdflexgo

import "github.com/kjk/flex"

type abstractFlexNode struct {
	flexNode *flex.Node
}

func (elem *abstractFlexNode) getFlexNode() *flex.Node {
	return elem.flexNode
}

func (elem *abstractFlexNode) x() float32 {
	return getXParent(elem.getFlexNode())
}

func (elem *abstractFlexNode) y() float32 {
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
