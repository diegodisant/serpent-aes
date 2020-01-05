package types

import "fmt"

type BitstringBinaryByteSliceTransformer interface {
	FromBinaryByteSliceToBitstring(binaryContent []byte) Bitstring
	FromBitSliceToByteSlice(bitSlice BitSlice) []byte
}

func (bitstring *Bitstring) FromBinaryByteSliceToBitstring(binaryContent []byte) Bitstring {
	var bytesSize int = len(binaryContent)
	var byteAsInt int64 = 0
	var convertedBitstring Bitstring

	for byteIndex := 0; byteIndex < bytesSize; byteIndex++ {
		byteAsInt = int64(binaryContent[byteIndex])
		convertedBitstring += bitstring.FromIntToBitstring(byteAsInt, 8)
	}

	return convertedBitstring
}

func (bitstring *Bitstring) FromBitSliceToByteSlice(bitSlice BitSlice) []byte {
	//var convertedByte byte
	var byteSliceSize int = len(bitSlice)
	var currentBitstring Bitstring
	//var binaryByteAsString Bitstring
	//var byteDecimalValue uint8
	var finalByteBufferSize int = byteSliceSize * 16

	var outputByteBuffer []byte = make([]byte, finalByteBufferSize)

	fmt.Println(bitSlice)

	for sliceIndex := 0; sliceIndex < byteSliceSize; sliceIndex++ {
		currentBitstring = bitSlice[sliceIndex]

		fmt.Println(currentBitstring)

		// for byteIndex := 0; byteIndex < 16; byteIndex += 8 {
		// 	binaryByteAsString = currentBitstring[byteIndex : (byteIndex+1)*8]
		// 	byteDecimalValue = bitstring.FromBitstringToUInt8(binaryByteAsString)
		// 	convertedByte = byte(byteDecimalValue)
		// 	outputByteBuffer = append(outputByteBuffer, convertedByte)
		// }
	}

	return outputByteBuffer
}
