package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageMarginWidth(t *testing.T) {

	page := NewPage().
		Border(5, 5, 5, 5)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderTop(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))

	page = NewPage().
		BorderRight(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))

	page = NewPage().
		BorderBottom(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	page = NewPage().
		BorderLeft(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderVertical(7)

	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	page = NewPage().
		BorderHorizontal(7)

	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(7), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	page = NewPage().
		BorderAll(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
