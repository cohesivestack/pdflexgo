package pdflexgo

type Footer struct {
	delegated               *BlockElement
	skipRenderAfterOverflow bool
	pageNumber              int
	pageName                string
}

type FooterBuilder = func(footer *Footer)

func footer() *Footer {
	footer := &Footer{}
	footer.delegated = Block()

	return footer
}

func (footer *Footer) PageNumber() int {
	return footer.pageNumber
}

func (footer *Footer) PageName() string {
	return footer.pageName
}

func (footer *Footer) SkipRenderAfterOverflow(skip bool) *Footer {
	footer.skipRenderAfterOverflow = skip
	return footer
}
