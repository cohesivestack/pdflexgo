package pdflexgo

type Body struct {
	*Segment
}

type BodyBuilder = func(body *Body)

func body() *Body {
	return &Body{
		Segment: segment(),
	}
}
