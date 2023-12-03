package pdflexgo

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/kjk/flex"
)

type PageElement struct {
	layout      *BlockElement
	bodyBuilder SegmentBuilder
	body        *Segment

	// Public attributes
	orientation Orientation
	unit        Unit
	name        string

	overflow PageOverflow
}

func Page() *PageElement {
	page := &PageElement{
		layout: Block().FlexDirectionColumn(),
	}

	page.OverflowNew()
	page.Size(DefaultSize)

	return page
}

func (page *PageElement) createOverflowedPage(overflowedNodes []Node) *PageElement {

	for _, node := range overflowedNodes {
		node.getFlexNode().Parent = nil
		node.markRequiredAsDirty()
	}

	newPage := Page()
	newPage.Name(page.name)
	newPage.Orientation(page.orientation)
	newPage.Unit(page.unit)
	newPage.Width(float64(page.layout.flexNode.LayoutGetWidth()))
	newPage.Height(float64(page.layout.flexNode.LayoutGetHeight()))
	newPage.Body(func(body *Segment) {
		page.body.copySegmentWithoutChildren(body)
		body.Children(overflowedNodes...)
	})

	newPage.layout.getFlexNode().Layout.Margin = page.layout.getFlexNode().Layout.Margin

	layoutNode := page.layout.getFlexNode()
	newLayoutNode := newPage.layout.getFlexNode()
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

	newPage.BackgroundColor(page.layout.backgroundColor.red, page.layout.backgroundColor.green, page.layout.backgroundColor.blue, page.layout.backgroundColor.alpha)

	newPage.BorderColorTop(page.layout.borderColor[EdgeTop].red, page.layout.borderColor[EdgeTop].green, page.layout.borderColor[EdgeTop].blue, page.layout.borderColor[EdgeTop].alpha)
	newPage.BorderColorRight(page.layout.borderColor[EdgeRight].red, page.layout.borderColor[EdgeRight].green, page.layout.borderColor[EdgeRight].blue, page.layout.borderColor[EdgeRight].alpha)
	newPage.BorderColorBottom(page.layout.borderColor[EdgeBottom].red, page.layout.borderColor[EdgeBottom].green, page.layout.borderColor[EdgeBottom].blue, page.layout.borderColor[EdgeBottom].alpha)
	newPage.BorderColorLeft(page.layout.borderColor[EdgeLeft].red, page.layout.borderColor[EdgeLeft].green, page.layout.borderColor[EdgeLeft].blue, page.layout.borderColor[EdgeLeft].alpha)

	return newPage
}

func (page *PageElement) Body(body SegmentBuilder) *PageElement {
	page.bodyBuilder = body
	return page
}

func (page *PageElement) Size(size Size) *PageElement {
	width, height := sizeToWidthHeight(size)

	page.Width(width)
	page.Height(height)

	return page
}

func (page *PageElement) Width(width float64) *PageElement {
	page.layout.Width(width)

	return page
}

func (page *PageElement) Height(height float64) *PageElement {
	page.layout.Height(height)

	return page
}

func (page *PageElement) Name(name string) *PageElement {
	page.name = name
	return page
}

func (page *PageElement) Orientation(orientation Orientation) *PageElement {
	page.orientation = orientation
	return page
}

func (page *PageElement) Unit(unit Unit) *PageElement {
	page.unit = unit
	return page
}

func (page *PageElement) markRequiredAsDirty() {
	for _, child := range page.layout.children {
		child.markRequiredAsDirty()
	}
}

func (page *PageElement) preRender(defaultProps *defaultProps, fpdf *gofpdf.Fpdf) {

	for _, child := range page.layout.children {
		child.preRender(defaultProps, fpdf)
	}
}

func (page *PageElement) render(pdf *Pdf, pageNumber int, overflowedContinuation bool) (nextPageWithOverflowed *PageElement) {

	// Set Layout
	var headerSegment *Header
	if pdf.header != nil {
		headerSegment = header()
		headerSegment.FlexDirectionColumn()
		headerSegment.delegated.Width(float64(page.layout.getFlexNode().LayoutGetWidth()))
		headerSegment.delegated.FlexNone()
		headerSegment.pageNumber = pageNumber
		headerSegment.pageName = page.name
		pdf.header(headerSegment)
		page.layout.Children(headerSegment.delegated)
	}

	if page.bodyBuilder != nil {
		page.body = segment().FlexDirectionColumn()
		page.body.delegated.Width(float64(page.layout.getFlexNode().LayoutGetWidth()))
		page.body.delegated.FlexAuto()
		page.body.pageNumber = pageNumber
		page.body.pageName = page.name
		page.bodyBuilder(page.body)
		page.layout.Children(page.body.delegated)
	}

	var footerSegment *Footer
	if pdf.footer != nil {
		footerSegment = footer()
		footerSegment.FlexDirectionColumn()
		footerSegment.delegated.Width(float64(page.layout.getFlexNode().LayoutGetWidth()))
		footerSegment.delegated.FlexNone()
		footerSegment.pageNumber = pageNumber
		footerSegment.pageName = page.name
		pdf.footer(footerSegment)
		page.layout.Children(footerSegment.delegated)
	}

	// Initialize fpdf
	initializeFpdf := func(fpdf *gofpdf.Fpdf) {
		fpdf.AddPageFormat(
			string(page.orientation),
			gofpdf.SizeType{
				Wd: float64(page.layout.getFlexNode().StyleGetWidth().Value),
				Ht: float64(page.layout.getFlexNode().StyleGetHeight().Value)})
		fpdf.SetFont("Arial", "", DefaultFontSize)
	}

	fpdf := pdf.fpdf

	initializeFpdf(fpdf)

	// Prerender is used to calculate the size of the elements with the
	// CalculateLayout process
	fpdfTemp := gofpdf.New(string(DefaultOrientation), string(DefaultUnit), string(DefaultSize), "")
	initializeFpdf(fpdfTemp)

	for _, fontLoaded := range pdf.fontsLoaded {
		_family, _style := getFontFamilyAndStyle(fontLoaded.fontFamily, fontLoaded.style)
		fpdfTemp.AddUTF8Font(_family, _style, fontLoaded.filePath)
	}

	page.preRender(pdf.defaultProps, fpdfTemp)

	// Calculate Flex nodes
	flex.CalculateLayout(page.layout.getFlexNode(), flex.Undefined, flex.Undefined, page.layout.getFlexNode().Style.Direction)

	page.markRequiredAsDirty()

	flex.CalculateLayout(page.layout.getFlexNode(), flex.Undefined, flex.Undefined, page.layout.getFlexNode().Style.Direction)

	if page.body != nil && page.body.delegated.getFlexNode().Style.FlexDirection == flex.FlexDirectionColumn && page.body.delegated.getFlexNode().Layout.HadOverflow {
		bodyHeight := page.body.delegated.getFlexNode().LayoutGetHeight()
		nodesVerticalSpace := 0.0
		for i, node := range page.body.delegated.children {
			if node.getFlexNode().Style.PositionType != flex.PositionTypeAbsolute {
				nodesVerticalSpace += float64(node.getFlexNode().LayoutGetHeight() + node.getFlexNode().LayoutGetMargin(flex.EdgeTop) + node.getFlexNode().LayoutGetMargin(flex.EdgeBottom))
				if i > 0 && nodesVerticalSpace > float64(bodyHeight) {
					overflowedNodes := page.body.delegated.children[i:]
					page.body.delegated.children = page.body.delegated.children[:i]
					nextPageWithOverflowed = page.createOverflowedPage(overflowedNodes)
					break
				}
			}
		}
	}

	page.layout.renderElement(pdf)

	if headerSegment != nil && (!overflowedContinuation || !headerSegment.skipRenderAfterOverflow) {
		headerSegment.delegated.render(pdf)
	}

	if page.body != nil {
		page.body.delegated.render(pdf)
	}

	if footerSegment != nil && (nextPageWithOverflowed == nil || !footerSegment.skipRenderAfterOverflow) {
		footerSegment.delegated.render(pdf)
	}

	return
}

func (page *PageElement) OverflowNew() *PageElement {
	page.overflow = PageOverflowNew
	return page
}

func (page *PageElement) Overflow(overflow PageOverflow) *PageElement {
	page.overflow = overflow
	return page
}

func (page *PageElement) OverflowHidden() *PageElement {
	page.overflow = PageOverflowHidden
	return page
}

func sizeToWidthHeight(size Size) (float64, float64) {

	switch size {
	case SizeA3:
		return 841.89, 1190.55
	case SizeA4:
		return 595.28, 841.89
	case SizeA5:
		return 420.94, 595.28
	case SizeA6:
		return 297.64, 420.94
	case SizeA2:
		return 1190.55, 1683.78
	case SizeA1:
		return 1683.78, 2383.94
	case SizeLetter:
		return 612, 792
	case SizeLegal:
		return 612, 1008
	case SizeTabloid:
		return 792, 1224
	}

	log.Fatal(fmt.Printf("standard size %s is not supported", size))

	return 0, 0
}
