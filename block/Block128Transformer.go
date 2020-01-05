package block

import t "serpent_aes/types"

type Block128Transformer interface {
	FromUInt64ToBlock128(quadWord uint64) (block128OutputTransformed Block128)
	FromBitstringToBlock128(bitstring t.Bitstring) Block128
}

func (block *Block128) FromUInt64ToBlock128(quadWord uint64) (block128OutputTransformed Block128) {
	for bitIndex := 0; bitIndex < 8; bitIndex++ {
		block128OutputTransformed[bitIndex] = byte(quadWord & 0xff)
		quadWord >>= 8
	}

	return block128OutputTransformed
}

func (block *Block128) FromBitstringToBlock128(bitstring t.Bitstring) Block128 {
	// If bitstring lenght is less than 128 buffer size
	// then pad the bitstring and add zeroes 48 in ASCII
	if len(bitstring) < BLOCK_128_SIZE {
		bitstring.Pad(BLOCK_128_SIZE - len(bitstring))
	}

	var blockTransformed Block128
	var temporarySlice = make([]byte, 8)
	var temporaryByte byte = 0
	var byteIndex int = 0

	for bitIndex := 0; bitIndex < BLOCK_128_SIZE; bitIndex += 8 {
		block.avoidSpecialCharacters(&temporarySlice, bitstring, bitIndex)
		block.transformByte(temporarySlice, &temporaryByte)
		blockTransformed[byteIndex] = temporaryByte
		byteIndex++
	}

	return blockTransformed
}

func (block *Block128) avoidSpecialCharacters(
	slice *[]byte,
	bitstring t.Bitstring,
	bitIndex int) {
	for bsBitIndex, character := range bitstring[bitIndex : bitIndex+8] {
		if character == 49 {
			(*slice)[bsBitIndex] = 1
		} else {
			(*slice)[bsBitIndex] = 0
		}
	}
}

func (block *Block128) transformByte(
	slice []byte,
	currentByte *byte) {
	*currentByte = 0

	for bitIndex := 7; bitIndex >= 0; bitIndex-- {
		if bitIndex == 7 && slice[bitIndex] == 0 {
			continue
		} else if slice[bitIndex] == 1 {
			*currentByte += 1 << uint(bitIndex)
		}
	}
}
