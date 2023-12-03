package pdflexgo

type Header struct {
	*Segment
	skipRenderAfterOverflow bool
}

type HeaderBuilder = func(header *Header)

func header() *Header {
	return &Header{
		Segment: segment(),
	}
}

func (header *Header) SkipRenderAfterOverflow(skip bool) *Header {
	header.skipRenderAfterOverflow = skip
	return header
}
