package pdflexgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorders(t *testing.T) {

	NewPdf().Pages(
		NewPage().
			MarginWidth(10).
			BorderColor("#ff0000").
			BorderWidth(1).
			Children(
				NewBlock().
					BorderColor("#0000ff").
					BorderWidth(1).
					Width(100).
					Height(100),
			),
	).Render()

	assert.True(t, true)

}
