package pdflexgo

func (page *Page) MarginTop(margin float64) *Page {

	page.root.MarginTop(margin)

	return page
}

func (page *Page) MarginBottom(margin float64) *Page {

	page.root.MarginBottom(margin)

	return page
}

func (page *Page) MarginLeft(margin float64) *Page {

	page.root.MarginLeft(margin)

	return page
}

func (page *Page) MarginRight(margin float64) *Page {

	page.root.MarginRight(margin)

	return page
}

func (page *Page) Margin(margin float64) *Page {
	page.root.
		MarginTop(margin).
		MarginRight(margin).
		MarginBottom(margin).
		MarginLeft(margin)

	return page
}
