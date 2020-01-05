package block

// Interface SplitOperator is used to split one 128 bit block
// into small segments
type Block128SplitOperator interface {
	QuadSplit() BlockSlice
}

func (block *Block128) QuadSplit() BlockSlice {
	var block32 Block32

	return BlockSlice{
		block32.FromUInt32ToBlock32(block[0:3]),
		block32.FromUInt32ToBlock32(block[4:7]),
		block32.FromUInt32ToBlock32(block[8:11]),
		block32.FromUInt32ToBlock32(block[12:15]),
	}
}
