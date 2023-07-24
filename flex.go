package pdflexgo

import (
	"github.com/kjk/flex"
)

func Test() {

	config := flex.NewConfig()

	root := flex.NewNodeWithConfig(config)

	root.StyleSetBorder(flex.EdgeAll, 5)

}
