package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageMargin(t *testing.T) {

	page := Page().
		Margin(5, 4, 3, 2)

	assert.Equal(t, float32(5), page.body.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), page.body.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), page.body.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), page.body.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = Page().
		MarginTop(5).
		MarginRight(4).
		MarginBottom(3).
		MarginLeft(2)

	assert.Equal(t, float32(5), page.body.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), page.body.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), page.body.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), page.body.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = Page().
		MarginVertical(6).
		MarginHorizontal(5)

	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(5), page.body.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(5), page.body.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	page = Page().
		MarginAll(6)

	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(6), page.body.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

}
