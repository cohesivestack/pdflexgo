package pdflexgo

type Footer struct {
	*Segment
	skipRenderAfterOverflow bool
}

type FooterBuilder = func(footer *Footer)

func footer() *Footer {
	return &Footer{
		Segment: segment(),
	}
}

func (footer *Footer) SkipRenderAfterOverflow(skip bool) *Footer {
	footer.skipRenderAfterOverflow = skip
	return footer
}
