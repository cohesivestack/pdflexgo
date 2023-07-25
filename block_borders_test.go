package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorders(t *testing.T) {

	NewPdf().Pages(
		NewPage().Children(
			NewBlock().
				BorderColor("#ff0000").
				BorderWidth(1).
				Width(100).
				Height(800),
		),
	).Render()

	assert.True(t, true)

}
