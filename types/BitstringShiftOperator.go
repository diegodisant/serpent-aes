package types

type BitstringShiftOperator interface {
	ShiftLeft(bitsToRotate int) Bitstring
	ShiftRight(bitsToRotate int) Bitstring
}

func (bitstring *Bitstring) ShiftLeft(bitsToRotate int) Bitstring {
	var byteSlice []byte = bitstring.FromBitstringToByteSlice()
	var byteSliceLength int = len(byteSlice)
	var temporarySlice = make([]byte, byteSliceLength)

	if bitsToRotate < 0 {
		return bitstring.ShiftRight(bitsToRotate - bitsToRotate*2)
	}

	for byteIndex := 0; byteIndex < byteSliceLength; byteIndex++ {
		if byteIndex < bitsToRotate {
			temporarySlice[byteIndex] = '0'
		} else if byteIndex == bitsToRotate {
			temporarySlice[bitsToRotate] = byteSlice[byteIndex-bitsToRotate]
		} else {
			temporarySlice[byteIndex] = byteSlice[byteIndex-bitsToRotate]
		}
	}

	return Bitstring(temporarySlice)
}

func (bitstring *Bitstring) ShiftRight(bitsToRotate int) Bitstring {
	var byteSlice []byte = bitstring.FromBitstringToByteSlice()
	var byteSliceLength int = len(byteSlice)
	var temporarySlice = make([]byte, byteSliceLength)

	for byteIndex := 0; byteIndex < byteSliceLength; byteIndex++ {
		if byteIndex <= byteSliceLength {
			temporarySlice[byteIndex] = byteSlice[bitsToRotate+byteIndex]
		} else if (byteIndex + bitsToRotate) < byteSliceLength {
			temporarySlice[byteIndex] = byteSlice[byteIndex+bitsToRotate]
		} else {
			temporarySlice[byteIndex] = '0'
		}
	}

	return Bitstring(temporarySlice)
}
