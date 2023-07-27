package pdflexgo

func (page *Page) MarginTopWidth(width float64) *Page {

	page.root.MarginTop(width)

	return page
}

func (page *Page) MarginBottomWidth(width float64) *Page {

	page.root.MarginBottom(width)

	return page
}

func (page *Page) MarginLeftWidth(width float64) *Page {

	page.root.MarginLeft(width)

	return page
}

func (page *Page) MarginRightWidth(width float64) *Page {

	page.root.MarginRight(width)

	return page
}

func (page *Page) MarginWidth(width float64) *Page {
	page.root.
		MarginTop(width).
		MarginRight(width).
		MarginBottom(width).
		MarginLeft(width)

	return page
}
