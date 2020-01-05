package types

import b64_helper "serpent_aes/helper/base64"

type BitsSliceBase64Transformer interface {
	FromLongBase64ToBitSlice(longBase64 string) BitSlice
}

func (bitslice *BitSlice) FromLongBase64ToBitSlice(longBase64 string) BitSlice {
	var sliceLength int = len(longBase64) / 232
	var outputSlice BitSlice = make([]Bitstring, sliceLength)
	var stringToDecode string
	var base64StringDecoded string
	var bitstringDecoded Bitstring

	for sliceIndex := 0; sliceIndex < sliceLength; sliceIndex++ {
		stringToDecode = longBase64[sliceIndex*232 : (sliceIndex+1)*232]
		base64StringDecoded = b64_helper.DecodeString(stringToDecode)
		bitstringDecoded = bitstringDecoded.FromBase64ToBitstring(base64StringDecoded)

		outputSlice[sliceIndex] = bitstringDecoded
	}

	return outputSlice
}
