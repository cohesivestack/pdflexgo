package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorders(t *testing.T) {

	NewPdf().Pages(
		NewPage().
			MarginLeft(10).
			MarginRight(20).
			MarginTop(30).
			MarginBottom(40).
			BorderColor("#ff0000").
			BorderWidth(1),
	).Render()

	assert.True(t, true)

}
