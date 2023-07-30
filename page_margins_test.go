package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageMargin(t *testing.T) {

	page := NewPage().
		Margin(5, 4, 3, 2)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), page.root.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = NewPage().
		MarginTop(5).
		MarginRight(4).
		MarginBottom(3).
		MarginLeft(2)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), page.root.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = NewPage().
		MarginVertical(6).
		MarginHorizontal(5)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = NewPage().
		MarginAll(6)

	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(6), page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

}
