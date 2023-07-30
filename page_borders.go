package pdflexgo

func (page *Page) BorderTopWidth(width float64) *Page {

	page.root.BorderTopWidth(width)

	return page
}

func (page *Page) BorderBottomWidth(width float64) *Page {

	page.root.BorderBottomWidth(width)

	return page
}

func (page *Page) BorderLeftWidth(width float64) *Page {

	page.root.BorderLeftWidth(width)

	return page
}

func (page *Page) BorderRightWidth(width float64) *Page {

	page.root.BorderRightWidth(width)

	return page
}

func (page *Page) BorderWidth(top float64, right float64, bottom float64, left float64) *Page {
	page.BorderTopWidth(top).
		BorderRightWidth(right).
		BorderBottomWidth(bottom).
		BorderLeftWidth(left)

	return page
}

func (page *Page) BorderTopColor(color string) *Page {

	page.root.BorderTopColor(color)

	return page
}

func (page *Page) BorderRightColor(color string) *Page {

	page.root.BorderRightColor(color)

	return page
}

func (page *Page) BorderBottomColor(color string) *Page {

	page.root.BorderBottomColor(color)

	return page
}

func (page *Page) BorderLeftColor(color string) *Page {

	page.root.BorderLeftColor(color)

	return page
}

func (page *Page) BorderColor(top string, right string, bottom string, left string) *Page {
	page.root.
		BorderTopColor(top).
		BorderRightColor(right).
		BorderBottomColor(bottom).
		BorderLeftColor(left)
	return page
}

func (page *Page) BorderVerticalWidth(vertical float64) *Page {
	page.root.
		BorderTopWidth(vertical).
		BorderBottomWidth(vertical)

	return page
}

func (page *Page) BorderHorizontalWidth(horizontal float64) *Page {
	page.root.
		BorderLeftWidth(horizontal).
		BorderRightWidth(horizontal)

	return page
}

func (page *Page) BorderAllWidth(border float64) *Page {
	page.root.
		BorderTopWidth(border).
		BorderRightWidth(border).
		BorderBottomWidth(border).
		BorderLeftWidth(border)

	return page
}

func (page *Page) BorderVerticalColor(color string) *Page {
	page.root.
		BorderTopColor(color).
		BorderBottomColor(color)

	return page
}

func (page *Page) BorderHorizontalColor(color string) *Page {
	page.root.
		BorderLeftColor(color).
		BorderRightColor(color)

	return page
}

func (page *Page) BorderAllColor(color string) *Page {
	page.root.
		BorderTopColor(color).
		BorderRightColor(color).
		BorderBottomColor(color).
		BorderLeftColor(color)

	return page
}
