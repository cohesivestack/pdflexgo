package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageBorderWidth(t *testing.T) {

	page := Page().
		BorderWidth(5, 4, 3, 2)

	assert.Equal(t, float32(5), page.layout.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), page.layout.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), page.layout.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), page.layout.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = Page().
		BorderWidthTop(5).
		BorderWidthRight(4).
		BorderWidthBottom(3).
		BorderWidthLeft(2)

	assert.Equal(t, float32(5), page.layout.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), page.layout.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), page.layout.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), page.layout.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = Page().
		BorderWidthVertical(7).
		BorderWidthHorizontal(6)

	assert.Equal(t, float32(7), page.layout.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), page.layout.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = Page().
		BorderWidthAll(6)

	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), page.layout.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
