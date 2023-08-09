package pdflexgo

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func setFont(fpdf *gofpdf.Fpdf, fontFamily string, style FontStyle, size float64) {

	_fontFamily, _style := getFontFamilyAndStyle(fontFamily, style)

	fpdf.SetFont(_fontFamily, _style, size)
}

func getFontFamilyAndStyle(fontFamily string, style FontStyle) (string, string) {

	isStandard := true
	switch fontFamily {
	case FontFamilyCourier, FontFamilyHelvetica, FontFamilyArial, FontFamilyTime, FontFamilySymbol:
	default:
		isStandard = false
	}

	if isStandard {
		switch style {
		case FontStyleRegular:
			style = ""
		case FontStyleRegularItalic:
			style = "I"
		case FontStyleBold:
			style = "B"
		case FontStyleBoldItalic:
			style = "BI"
		default:
			isStandard = false
		}
	}

	if !isStandard {
		fontFamily = fmt.Sprintf("%s!%s", fontFamily, style)
		style = ""
	}

	return fontFamily, string(style)
}
