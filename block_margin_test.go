package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageMargin(t *testing.T) {

	value := flex.Value{Value: 5, Unit: 1}
	page := NewPage().
		Margin(5, 5, 5, 5)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeTop))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeRight))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft))

	value.Value = 6

	page = NewPage().
		MarginTop(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeTop))

	page = NewPage().
		MarginRight(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeRight))

	page = NewPage().
		MarginBottom(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom))

	page = NewPage().
		MarginLeft(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft))

	page = NewPage().
		MarginVertical(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeTop))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom))

	page = NewPage().
		MarginHorizontal(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeRight))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft))

	page = NewPage().
		MarginAll(6)

	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeTop))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeRight))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeBottom))
	assert.Equal(t, value, page.root.getFlexNode().StyleGetMargin(flex.EdgeLeft))

}
