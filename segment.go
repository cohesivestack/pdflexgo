package pdflexgo

type Segment struct {
	delegated  *BlockElement
	pageNumber int
	pageName   string
}

type SegmentBuilder = func(header *Segment)

func segment() *Segment {
	segment := &Segment{}
	segment.delegated = Block()

	return segment
}
