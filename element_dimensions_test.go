package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockDimensions(t *testing.T) {

	pdf := NewPdf()

	block := Block().Width(400).Height(200)
	pdf.Pages(Page().Children(block)).Render()

	assert.Equal(t, float32(400), block.getFlexNode().LayoutGetWidth())
	assert.Equal(t, float32(200), block.getFlexNode().LayoutGetHeight())
}

func TestBlockDimensionsWithMargin(t *testing.T) {

	pdf := NewPdf()

	block := Block().Width(400).Height(200).Margin(10, 20, 30, 40)
	pdf.Pages(Page().Children(block)).Render()

	assert.Equal(t, float32(400), block.getFlexNode().LayoutGetWidth())
	assert.Equal(t, float32(200), block.getFlexNode().LayoutGetHeight())
}

func TestBlockDimensionsWithPadding(t *testing.T) {

	pdf := NewPdf()

	block := Block().Width(400).Height(200).Padding(10, 20, 30, 40)
	pdf.Pages(Page().Children(block)).Render()

	assert.Equal(t, float32(400), block.getFlexNode().LayoutGetWidth())
	assert.Equal(t, float32(200), block.getFlexNode().LayoutGetHeight())
}

func TestBlockDimensionsWithBorder(t *testing.T) {

	pdf := NewPdf()

	block := Block().Width(400).Height(200).BorderWidth(10, 20, 30, 40)
	pdf.Pages(Page().Children(block)).Render()

	assert.Equal(t, float32(400), block.getFlexNode().LayoutGetWidth())
	assert.Equal(t, float32(200), block.getFlexNode().LayoutGetHeight())
}
