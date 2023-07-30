package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageMarginWidth(t *testing.T) {

	page := NewPage().
		BorderWidth(5, 5, 5, 5)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderTopWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))

	page = NewPage().
		BorderRightWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))

	page = NewPage().
		BorderBottomWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	page = NewPage().
		BorderLeftWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderVerticalWidth(7)

	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	page = NewPage().
		BorderHorizontalWidth(7)

	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderAllWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
