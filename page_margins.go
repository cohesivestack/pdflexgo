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

func (page *Page) Margin(top float64, right float64, bottom float64, left float64) *Page {
	page.root.
		MarginTop(top).
		MarginRight(right).
		MarginBottom(bottom).
		MarginLeft(left)

	return page
}

func (page *Page) MarginVertical(vertical float64) *Page {
	page.root.
		MarginTop(vertical).
		MarginBottom(vertical)

	return page
}

func (page *Page) MarginHorizontal(horizontal float64) *Page {
	page.root.
		MarginLeft(horizontal).
		MarginRight(horizontal)

	return page
}

func (page *Page) MarginAll(margin float64) *Page {
	page.root.
		MarginTop(margin).
		MarginRight(margin).
		MarginBottom(margin).
		MarginLeft(margin)

	return page
}
