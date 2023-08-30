package pdflexgo

import "github.com/kjk/flex"

type Segment struct {
	delegated  *BlockElement
	pageNumber int
	pageName   string
}

type SegmentBuilder = func(header *Segment)

func segment() *Segment {
	segment := &Segment{}
	segment.delegated = Block()

	return segment
}

func (segment *Segment) PageNumber() int {
	return segment.pageNumber
}

func (segment *Segment) PageName() string {
	return segment.pageName
}

func (_segment *Segment) copySegmentWithoutChildren(newSegment *Segment) {

	layoutNode := _segment.delegated.getFlexNode()
	newLayoutNode := newSegment.delegated.getFlexNode()
	for _, edge := range []flex.Edge{flex.EdgeTop, flex.EdgeRight, flex.EdgeBottom, flex.EdgeLeft} {
		margin := layoutNode.StyleGetMargin(edge)
		if margin.Unit == flex.UnitPercent {
			newLayoutNode.StyleSetMarginPercent(edge, margin.Value)
		} else if margin.Unit == flex.UnitPoint {
			newLayoutNode.StyleSetMargin(edge, margin.Value)
		} else if margin.Unit == flex.UnitAuto {
			newLayoutNode.StyleSetMarginAuto(edge)
		}

		padding := layoutNode.StyleGetPadding(edge)
		if padding.Unit == flex.UnitPercent {
			newLayoutNode.StyleSetPaddingPercent(edge, padding.Value)
		} else if padding.Unit == flex.UnitPoint {
			newLayoutNode.StyleSetPadding(edge, padding.Value)
		}

		newLayoutNode.StyleSetBorder(edge, layoutNode.StyleGetBorder(edge))

	}

	newSegment.BackgroundColor(_segment.delegated.backgroundColor.red, _segment.delegated.backgroundColor.green, _segment.delegated.backgroundColor.blue, _segment.delegated.backgroundColor.alpha)

	newSegment.BorderColorTop(_segment.delegated.borderColor[EdgeTop].red, _segment.delegated.borderColor[EdgeTop].green, _segment.delegated.borderColor[EdgeTop].blue, _segment.delegated.borderColor[EdgeTop].alpha)
	newSegment.BorderColorRight(_segment.delegated.borderColor[EdgeRight].red, _segment.delegated.borderColor[EdgeRight].green, _segment.delegated.borderColor[EdgeRight].blue, _segment.delegated.borderColor[EdgeRight].alpha)
	newSegment.BorderColorBottom(_segment.delegated.borderColor[EdgeBottom].red, _segment.delegated.borderColor[EdgeBottom].green, _segment.delegated.borderColor[EdgeBottom].blue, _segment.delegated.borderColor[EdgeBottom].alpha)
	newSegment.BorderColorLeft(_segment.delegated.borderColor[EdgeLeft].red, _segment.delegated.borderColor[EdgeLeft].green, _segment.delegated.borderColor[EdgeLeft].blue, _segment.delegated.borderColor[EdgeLeft].alpha)

	newSegment.Direction(Direction(_segment.delegated.getFlexNode().Layout.Direction))
	newSegment.FlexDirection(FlexDirection(_segment.delegated.getFlexNode().Style.FlexDirection))
	newSegment.JustifyContent(Justify(_segment.delegated.getFlexNode().Style.JustifyContent))
	newSegment.AlignContentType(Align(_segment.delegated.getFlexNode().Style.AlignContent))
	newSegment.AlignItemsType(Align(_segment.delegated.getFlexNode().Style.AlignItems))
	newSegment.FlexWrapValue(Wrap(_segment.delegated.getFlexNode().Style.FlexWrap))

}
