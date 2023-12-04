package pdflexgo

type Header struct {
	delegated               *BlockElement
	skipRenderAfterOverflow bool
	pageNumber              int
	pageName                string
}

type HeaderBuilder = func(header *Header)

func header() *Header {
	header := &Header{}
	header.delegated = Block()

	return header
}

func (header *Header) PageNumber() int {
	return header.pageNumber
}

func (header *Header) PageName() string {
	return header.pageName
}

func (header *Header) SkipRenderAfterOverflow(skip bool) *Header {
	header.skipRenderAfterOverflow = skip
	return header
}
