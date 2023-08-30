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

func (segment *Segment) PageNumber() int {
	return segment.pageNumber
}

func (segment *Segment) PageName() string {
	return segment.pageName
}
