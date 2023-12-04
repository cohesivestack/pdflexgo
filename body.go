package pdflexgo

import "github.com/kjk/flex"

type Body struct {
	delegated  *BlockElement
	pageNumber int
}

type BodyBuilder = func(body *Body)

func body() *Body {
	body := &Body{}
	body.delegated = Block()

	return body
}

func (body *Body) PageNumber() int {
	return body.pageNumber
}

func (_body *Body) copyBodyWithoutChildren(newBody *Body) {

	layoutNode := _body.delegated.getFlexNode()
	newLayoutNode := newBody.delegated.getFlexNode()
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

	newBody.BackgroundColor(_body.delegated.backgroundColor.red, _body.delegated.backgroundColor.green, _body.delegated.backgroundColor.blue, _body.delegated.backgroundColor.alpha)

	newBody.BorderColorTop(_body.delegated.borderColor[EdgeTop].red, _body.delegated.borderColor[EdgeTop].green, _body.delegated.borderColor[EdgeTop].blue, _body.delegated.borderColor[EdgeTop].alpha)
	newBody.BorderColorRight(_body.delegated.borderColor[EdgeRight].red, _body.delegated.borderColor[EdgeRight].green, _body.delegated.borderColor[EdgeRight].blue, _body.delegated.borderColor[EdgeRight].alpha)
	newBody.BorderColorBottom(_body.delegated.borderColor[EdgeBottom].red, _body.delegated.borderColor[EdgeBottom].green, _body.delegated.borderColor[EdgeBottom].blue, _body.delegated.borderColor[EdgeBottom].alpha)
	newBody.BorderColorLeft(_body.delegated.borderColor[EdgeLeft].red, _body.delegated.borderColor[EdgeLeft].green, _body.delegated.borderColor[EdgeLeft].blue, _body.delegated.borderColor[EdgeLeft].alpha)

	newBody.Direction(Direction(_body.delegated.getFlexNode().Layout.Direction))
	newBody.FlexDirection(FlexDirection(_body.delegated.getFlexNode().Style.FlexDirection))
	newBody.JustifyContent(Justify(_body.delegated.getFlexNode().Style.JustifyContent))
	newBody.AlignContentType(Align(_body.delegated.getFlexNode().Style.AlignContent))
	newBody.AlignItemsType(Align(_body.delegated.getFlexNode().Style.AlignItems))
	newBody.FlexWrapValue(Wrap(_body.delegated.getFlexNode().Style.FlexWrap))

}
