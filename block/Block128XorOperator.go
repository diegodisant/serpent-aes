package block

// Interface XorOperator implemented as a Xor
// operator for blocks
type Block128XorOperator interface {
	// Method SegmentWithBlock128SizeXor applies xor operation
	// to all blocks that segment contains
	SegmentWithBlock128SizeXor(segment []Block128) *Block128

	// Method Block128SizeXor applies xor operation to one block
	// by using current block currentBlock XOR targetBlock
	Block128SizeXor(octaWord Block128)
}

func (block *Block128) SegmentWithBlock128SizeXor(segment []Block128) *Block128 {
	var segmentWithXorApplied *Block128 = &segment[0]

	for segmentIndex := 1; segmentIndex < len(segment); segmentIndex++ {
		segmentWithXorApplied.Block128SizeXor(segment[segmentIndex])
	}

	return segmentWithXorApplied
}

func (block *Block128) Block128SizeXor(octaWord Block128) {
	for byteIndex := 0; byteIndex < BLOCK_128B_BYTES_SIZE; byteIndex++ {
		block[byteIndex] ^= octaWord[byteIndex]
	}
}
