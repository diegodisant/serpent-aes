package types

import "encoding/base64"

type BitstringBase64Transformer interface {
	FromBitstringToBase64(decodedBitstring Bitstring) string
	FromBase64ToBitstring(encodedString string) Bitstring
}

func (bitstring *Bitstring) FromBitstringToBase64(decodedBitstring Bitstring) string {
	var encodedString string
	var bitstringBuffer []byte = decodedBitstring.FromBitstringToByteSlice()

	encodedString = base64.StdEncoding.EncodeToString(bitstringBuffer)

	return encodedString
}

func (bitstring *Bitstring) FromBase64ToBitstring(encodedString string) Bitstring {
	var decodedBitstring Bitstring
	var bitstringByteBuffer []byte

	bitstringByteBuffer, _ = base64.StdEncoding.DecodeString(encodedString)
	decodedBitstring = bitstring.FromByteSliceToBitstring(bitstringByteBuffer)

	return decodedBitstring
}
