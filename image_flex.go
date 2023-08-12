package pdflexgo

import (
	"github.com/kjk/flex"
)

func (image *ImageElement) Width(width float64) *ImageElement {
	image.getFlexNode().StyleSetWidth(float32(width))
	return image
}

func (image *ImageElement) WidthPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetWidthPercent(float32(percent))
	return image
}

func (image *ImageElement) Height(height float64) *ImageElement {
	image.getFlexNode().StyleSetHeight(float32(height))
	return image
}

func (image *ImageElement) HeightPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetHeightPercent(float32(percent))
	return image
}

func (image *ImageElement) PositionType(position PositionType) *ImageElement {
	image.getFlexNode().StyleSetPositionType(flex.PositionType(position))
	return image
}

func (image *ImageElement) GetPositionType() PositionType {
	return PositionType(image.getFlexNode().Style.PositionType)
}

func (image *ImageElement) Direction(direction Direction) *ImageElement {
	image.getFlexNode().StyleSetDirection(flex.Direction(direction))
	return image
}

func (image *ImageElement) GetDirection() Direction {
	return Direction(image.getFlexNode().Style.Direction)
}

func (image *ImageElement) FlexDirection(direction FlexDirection) *ImageElement {
	image.getFlexNode().StyleSetFlexDirection(flex.FlexDirection(direction))
	return image
}

func (image *ImageElement) GetFlexDirection() FlexDirection {
	return FlexDirection(image.getFlexNode().Style.FlexDirection)
}

func (image *ImageElement) JustifyContent(justifyContent Justify) *ImageElement {
	image.getFlexNode().StyleSetJustifyContent(flex.Justify(justifyContent))
	return image
}

func (image *ImageElement) GetJustifyContent() Justify {
	return Justify(image.getFlexNode().Style.JustifyContent)
}

func (image *ImageElement) AlignContent(alignContent flex.Align) *ImageElement {
	image.getFlexNode().StyleSetAlignContent(flex.Align(alignContent))
	return image
}

func (image *ImageElement) GetAlignContent() flex.Align {
	return flex.Align(image.getFlexNode().Style.AlignContent)
}

func (image *ImageElement) AlignItems(alignItems flex.Align) *ImageElement {
	image.getFlexNode().StyleSetAlignItems(flex.Align(alignItems))
	return image
}

func (image *ImageElement) GetAlignItems() flex.Align {
	return flex.Align(image.getFlexNode().Style.AlignItems)
}

func (image *ImageElement) AlignSelf(alignSelf flex.Align) *ImageElement {
	image.getFlexNode().StyleSetAlignSelf(flex.Align(alignSelf))
	return image
}

func (image *ImageElement) GetAlignSelf() flex.Align {
	return flex.Align(image.getFlexNode().Style.AlignSelf)
}

func (image *ImageElement) FlexWrap(wrap Wrap) *ImageElement {
	image.getFlexNode().StyleSetFlexWrap(flex.Wrap(wrap))
	return image
}

func (image *ImageElement) GetFlexWrap() Wrap {
	return Wrap(image.getFlexNode().Style.FlexWrap)
}

func (image *ImageElement) Overflow(overflow Overflow) *ImageElement {
	image.getFlexNode().StyleSetOverflow(flex.Overflow(overflow))
	return image
}

func (image *ImageElement) GetOverflow() Overflow {
	return Overflow(image.getFlexNode().Style.Overflow)
}

func (image *ImageElement) Display(overflow Display) *ImageElement {
	image.getFlexNode().StyleSetDisplay(flex.Display(overflow))
	return image
}

func (image *ImageElement) GetDisplay() Display {
	return Display(image.getFlexNode().Style.Display)
}

func (image *ImageElement) Flex(flex float64) *ImageElement {
	image.getFlexNode().StyleSetFlex(float32(flex))
	return image
}

func (image *ImageElement) GetFlex() float64 {
	return float64(image.getFlexNode().Style.Flex)
}

func (image *ImageElement) FlexGrow(grow float64) *ImageElement {
	image.getFlexNode().StyleSetFlexGrow(float32(grow))
	return image
}

func (image *ImageElement) GetFlexGrow() float64 {
	return float64(image.getFlexNode().StyleGetFlexGrow())
}

func (image *ImageElement) FlexShrink(shrink float64) *ImageElement {
	image.getFlexNode().StyleSetFlexShrink(float32(shrink))
	return image
}

func (image *ImageElement) GetFlexShrink() float64 {
	return float64(image.getFlexNode().StyleGetFlexShrink())
}

func (image *ImageElement) FlexBasis(basis float64) *ImageElement {
	image.getFlexNode().StyleSetFlexBasis(float32(basis))
	return image
}

func (image *ImageElement) FlexBasisPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetFlexBasisPercent(float32(percent))
	return image
}

func (image *ImageElement) MinWidth(minWidth float64) *ImageElement {
	image.getFlexNode().StyleSetMinWidth(float32(minWidth))
	return image
}

func (image *ImageElement) MinWidthPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetMinWidthPercent(float32(percent))
	return image
}

func (image *ImageElement) MinHeight(minHeight float64) *ImageElement {
	image.getFlexNode().StyleSetMinHeight(float32(minHeight))
	return image
}

func (image *ImageElement) MinHeightPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetMinHeightPercent(float32(percent))
	return image
}

func (image *ImageElement) MaxWidth(maxWidth float64) *ImageElement {
	image.getFlexNode().StyleSetMaxWidth(float32(maxWidth))
	return image
}

func (image *ImageElement) MaxWidthPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetMaxWidthPercent(float32(percent))
	return image
}

func (image *ImageElement) MaxHeight(maxHeight float64) *ImageElement {
	image.getFlexNode().StyleSetMaxHeight(float32(maxHeight))
	return image
}

func (image *ImageElement) MaxHeightPercent(percent float64) *ImageElement {
	image.getFlexNode().StyleSetMaxHeightPercent(float32(percent))
	return image
}

func (image *ImageElement) AspectRatio(aspectRatio float64) *ImageElement {
	image.getFlexNode().StyleSetAspectRatio(float32(aspectRatio))
	return image
}
