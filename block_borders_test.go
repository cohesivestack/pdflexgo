package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorders(t *testing.T) {

	NewPdf().Children(
		NewBlock().
			BorderColor("#ff0000").
			BorderWidth(1).
			Width(100).
			Height(100),
	).Build()

	assert.True(t, true)

}
