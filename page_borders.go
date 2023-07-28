package pdflexgo

func (page *Page) BorderTop(width float64) *Page {

	page.root.BorderTop(width)

	return page
}

func (page *Page) BorderBottom(width float64) *Page {

	page.root.BorderBottom(width)

	return page
}

func (page *Page) BorderLeft(width float64) *Page {

	page.root.BorderLeft(width)

	return page
}

func (page *Page) BorderRight(width float64) *Page {

	page.root.BorderRight(width)

	return page
}

func (page *Page) Border(top float64, right float64, bottom float64, left float64) *Page {
	page.BorderTop(top).
		BorderRight(right).
		BorderBottom(bottom).
		BorderLeft(left)

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

func (page *Page) BorderVertical(vertical float64) *Page {
	page.root.
		BorderTop(vertical).
		BorderBottom(vertical)

	return page
}

func (page *Page) BorderHorizontal(horizontal float64) *Page {
	page.root.
		BorderLeft(horizontal).
		BorderRight(horizontal)

	return page
}

func (page *Page) BorderAll(border float64) *Page {
	page.root.
		BorderTop(border).
		BorderRight(border).
		BorderBottom(border).
		BorderLeft(border)

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
