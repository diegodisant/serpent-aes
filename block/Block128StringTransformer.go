package block

type Block128StringTransformer interface {
	String() string
}

// @deprecated find a way to convert to string
// func (block *Block128) String() string {
// 	var segment Segment
// 	var segmentByteIndex int = 0
//
// 	for blockByteIndex := 0; blockByteIndex < BLOCK_128B_BYTES_SIZE; blockByteIndex++ {
// 		block.walkByte(blockByteIndex, &segment, &segmentByteIndex)
// 	}
//
// 	return string(segment[0:])
// }
//
// func (block *Block128) walkByte(
// 	blockByteIndex int,
// 	segment *Segment,
// 	segmentByteIndex *int) {
// 	for bitIndex := 0; bitIndex < 8; bitIndex++ {
// 		block.shiftBit(blockByteIndex, &bitIndex, segment, segmentByteIndex, CHAR_SHIFTS[2])
// 	}
// }
//
// func (block *Block128) shiftBit(
// 	blockByteIndex int,
// 	bitIndex *int,
// 	segment *Segment,
// 	segmentByteIndex *int,
// 	shiftSize uint) {
// 	var zeroIndexTrimmer uintptr = uintptr(1)
//
// 	for block[blockByteIndex] > 1 {
// 		(*segment)[*segmentByteIndex] = ASCII_PRINTABLE_CHARS[uintptr(block[blockByteIndex])&zeroIndexTrimmer]
// 		block[blockByteIndex] >>= shiftSize
//
// 		(*bitIndex)++
// 		(*segmentByteIndex)++
// 	}
//
// 	if *bitIndex < 7 && block[blockByteIndex] == 1 {
// 		(*segment)[*segmentByteIndex] = ASCII_PRINTABLE_CHARS[uintptr(block[blockByteIndex])]
// 		block[blockByteIndex]--
// 	} else {
// 		(*segment)[*segmentByteIndex] = ASCII_PRINTABLE_CHARS[uintptr(block[blockByteIndex])]
// 	}
//
// 	(*segmentByteIndex)++
// }
