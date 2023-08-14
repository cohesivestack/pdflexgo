package pdflexgo

type rgba struct {
	red   int
	green int
	blue  int
	alpha float64
}

var DefaultFontColor = rgba{0, 0, 0, 1.0}
var DefaultBackgroundColor = rgba{0, 0, 0, 0}

func getRgba(red int, green int, blue int, alpha ...float64) rgba {
	_alpha := 1.0
	if len(alpha) > 0 && alpha[0] < 1 {
		_alpha = alpha[0]
	}
	return rgba{red, green, blue, _alpha}
}

func equalColor(colorA rgba, colorB rgba) bool {
	return colorA.red == colorB.red &&
		colorA.green == colorB.green &&
		colorA.blue == colorB.blue &&
		colorA.alpha == colorB.alpha
}
