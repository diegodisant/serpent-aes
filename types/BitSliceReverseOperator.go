package types

type BitSliceReverseOperator interface {
	Reverse() []Bitstring
}

func (bitslice BitSlice) Reverse() []Bitstring {
	var sliceLength int = len(bitslice) - 1
	var bitstringBuffer = make([]Bitstring, sliceLength)

	for bitstringIndex := 0; bitstringIndex <= sliceLength; bitstringIndex++ {
		bitstringBuffer[bitstringIndex] = Bitstring(bitslice[sliceLength-bitstringIndex])
	}

	return bitstringBuffer
}
