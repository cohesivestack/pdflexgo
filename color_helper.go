package pdflexgo

import (
	"fmt"
	"strconv"
)

func hexToRGB(hexColor string) (int, int, int, error) {
	if len(hexColor) != 7 || hexColor[0] != '#' {
		return 0, 0, 0, fmt.Errorf("invalid color format")
	}

	r, err := strconv.ParseInt(hexColor[1:3], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}

	g, err := strconv.ParseInt(hexColor[3:5], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}

	b, err := strconv.ParseInt(hexColor[5:7], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(r), int(g), int(b), nil
}
