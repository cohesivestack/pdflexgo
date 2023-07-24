package pdflexgo

func (block *Block) Width(width float32) *Block {
	block.flexNode.StyleSetWidth(width)
	return block
}

func (block *Block) Height(height float32) *Block {
	block.flexNode.StyleSetHeight(height)
	return block
}
