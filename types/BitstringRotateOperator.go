package types

type BitstringRotateOperator interface {
	RotateLeft(bitsToRotate int) Bitstring
	RotateRight(bitsToRotate int) Bitstring
}

func (bitstring *Bitstring) RotateLeft(bitsToRotate int) Bitstring {
	var byteSlice []byte = bitstring.FromBitstringToByteSlice()
	var byteSliceLength int = len(byteSlice)
	var temporarySlice = make([]byte, byteSliceLength)
	var bitsRotated int

	for byteIndex := 0; byteIndex < byteSliceLength; byteIndex++ {
		if byteIndex < bitsToRotate {
			bitsRotated = byteSliceLength - bitsToRotate + byteIndex
		} else if byteIndex == bitsToRotate {
			bitsRotated = 0
		} else if byteIndex > bitsToRotate {
			bitsRotated = byteIndex - bitsToRotate
		}

		temporarySlice[byteIndex] = byteSlice[bitsRotated]
	}

	return Bitstring(temporarySlice)
}

func (bitstring *Bitstring) RotateRight(bitsToRotate int) Bitstring {
	var byteSlice []byte = bitstring.FromBitstringToByteSlice()
	var byteSliceLength int = len(byteSlice)
	var temporarySlice = make([]byte, byteSliceLength)
	var bitsRotated int

	for byteIndex := 0; byteIndex < byteSliceLength; byteIndex++ {
		if (byteIndex + bitsToRotate) < byteSliceLength {
			bitsRotated = byteIndex + bitsToRotate
		} else if (byteIndex + bitsToRotate) == byteSliceLength {
			bitsRotated = 0
		} else if (byteIndex + bitsToRotate) > byteSliceLength {
			bitsRotated = bitsToRotate + byteIndex - byteSliceLength
		}

		temporarySlice[byteIndex] = byteSlice[bitsRotated]
	}

	return Bitstring(temporarySlice)
}
