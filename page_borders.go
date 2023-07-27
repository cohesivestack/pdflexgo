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

func (page *Page) BorderWidth(width float64) *Page {
	page.BorderTopWidth(width).
		BorderRightWidth(width).
		BorderBottomWidth(width).
		BorderLeftWidth(width)

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

func (page *Page) BorderColor(color string) *Page {

	page.root.BorderColor(color)

	return page
}
