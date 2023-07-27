package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorders(t *testing.T) {

	NewPdf().Pages(
		NewPage().
			Margin(10).
			BorderColor("#ff0000").
			BorderWidth(1),
	).Render()

	assert.True(t, true)

}
