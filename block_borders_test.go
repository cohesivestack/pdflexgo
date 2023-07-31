package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestBlockBorderWidth(t *testing.T) {

	block := NewBlock().
		BorderWidth(5, 4, 3, 2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = NewBlock().
		BorderTopWidth(5).
		BorderRightWidth(4).
		BorderBottomWidth(3).
		BorderLeftWidth(2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = NewBlock().
		BorderVerticalWidth(7).
		BorderHorizontalWidth(6)

	assert.Equal(t, float32(7), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = NewBlock().
		BorderAllWidth(6)

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
