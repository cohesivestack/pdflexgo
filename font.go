package pdflexgo

import "github.com/jung-kurt/gofpdf"

func setFont(fpdf *gofpdf.Fpdf, fontFamily string, style FontStyle, size float64) {

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

	fpdf.SetFont(fontFamily, string(style), size)
}
