package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestBlockBorderWidth(t *testing.T) {

	block := Block().
		BorderWidth(5, 4, 3, 2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = Block().
		BorderWidthTop(5).
		BorderWidthRight(4).
		BorderWidthBottom(3).
		BorderWidthLeft(2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = Block().
		BorderWidthVertical(7).
		BorderWidthHorizontal(6)

	assert.Equal(t, float32(7), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(7), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

	block = Block().
		BorderWidthAll(6)

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
