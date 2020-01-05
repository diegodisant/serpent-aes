package types

// Interface PadOperator is used to handle
// variant segment size when is less than 128
// fill with zero in right side 01 -> 00000001
type BitstringPadOperator interface {
	// Method pad fills with zero when current block
	// size is less than 128 bit
	Pad(byteBufferSize int)
}

func (bitstring *Bitstring) Pad(byteBufferSize int) {
	var byteBuffer = make([]byte, byteBufferSize)

	for byteIndex := 0; byteIndex < byteBufferSize; byteIndex++ {
		byteBuffer[byteIndex] = 48
	}

	*bitstring = Bitstring(byteBuffer) + *bitstring
}
