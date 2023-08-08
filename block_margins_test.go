package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestBlockMargin(t *testing.T) {

	block := Block().
		Margin(5, 4, 3, 2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	block = Block().
		MarginTop(5).
		MarginRight(4).
		MarginBottom(3).
		MarginLeft(2)

	assert.Equal(t, float32(5), block.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(4), block.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(3), block.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(2), block.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	block = Block().
		MarginVertical(6).
		MarginHorizontal(5)

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(5), block.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(5), block.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

	block = Block().
		MarginAll(6)

	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeTop).Value)
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeRight).Value)
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeBottom).Value)
	assert.Equal(t, float32(6), block.getFlexNode().StyleGetMargin(flex.EdgeLeft).Value)

}
