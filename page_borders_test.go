package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageBorderWidth(t *testing.T) {

	page := NewPage().
		BorderWidth(5, 4, 3, 2)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderTopWidth(5).
		BorderRightWidth(4).
		BorderBottomWidth(3).
		BorderLeftWidth(2)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderVerticalWidth(7).
		BorderHorizontalWidth(6)

	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderAllWidth(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
