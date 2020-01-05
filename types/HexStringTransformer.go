package types

type HexStringTransformer interface {
	FromHexStringToBitstring() (bitstringTransformed Bitstring)
}

func (hexstring HexString) FromHexStringToBitstring() (bitstringTransformed Bitstring) {
	var hexStringLength int = len(hexstring) - 1
	var hexByteBuffer = make([]byte, hexStringLength)
	var reverseHexIndex int = 0

	//reverse hex byte buffer buffer acde to edca
	for hexIndex := hexStringLength; hexIndex >= 0; hexIndex-- {
		hexByteBuffer[hexIndex] = hexstring[reverseHexIndex]
		reverseHexIndex++
	}

	for hexIndex := 0; hexIndex < hexStringLength; hexIndex++ {
		bitstringTransformed += bitstringTransformed.FromHexToBitstringOctect(HexString(hexByteBuffer[hexIndex]))
	}

	return bitstringTransformed
}
