package pdflexgo

import (
	"testing"

	"github.com/kjk/flex"
	"github.com/stretchr/testify/assert"
)

func TestPageBorderWidth(t *testing.T) {

	page := NewPage().
		BorderWidth(5)

	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeTop))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeRight))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeBottom))
	assert.Equal(t, float32(5), page.root.getFlexNode().StyleGetBorder(flex.EdgeLeft))

}
