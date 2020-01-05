package types

import (
	"fmt"
	"math"
	"strings"
)

type BitstringTransformer interface {
	FromStringToBitstring(stringToConvert string) (convertedBitstring Bitstring)
	FromLongStringToBitSlice(longString string) BitSlice
	FromLongBitstringToBitSlice(longBitstring Bitstring) BitSlice
	FromBitstringSliceToString(slice BitSlice) string
	FromIntToBitstring(number int64, characterLength int) (bitstringTransformed Bitstring)
	FromBitstringToByteSlice() (sliceTransformed []byte)
	FromByteSliceToBitstring(byteSlice []byte) (outputBitstring Bitstring)
	FromBitstringOctectToHex() (hexstringTransformed HexString)
	FromHexToBitstringOctect(hexstring HexString) (bitstringTransformed Bitstring)
	FromBitstringToHexString() (hexstringTransformed HexString)
	FromBitstringToUInt8(bitstring Bitstring) uint8
	String(bitstringToConvert Bitstring) (asciiString string)
}

func (bitstring *Bitstring) FromStringToBitstring(stringToConvert string) (convertedBitstring Bitstring) {
	var stringLength int = len(stringToConvert)
	var charDecValue int

	for charIndex := 0; charIndex < stringLength; charIndex++ {
		charDecValue = int(stringToConvert[charIndex])
		convertedBitstring += bitstring.FromIntToBitstring(int64(charDecValue), 8)
	}

	if len(convertedBitstring) < 128 {
		convertedBitstring.Pad(128 - len(convertedBitstring))
	}

	return convertedBitstring
}

func (bitstring *Bitstring) FromLongStringToBitSlice(longString string) BitSlice {
	var _128BString string
	var _128BBitstring Bitstring
	var longStringLength int = len(longString)
	var initialSliceIndex int = 0
	var finalSliceIndex int = 0
	var bitstringSliceLength int = int(float64(longStringLength / 16))
	var bitstringOutputBuffer BitSlice
	var residualCharacters int = (longStringLength % 16)
	var hasResidualString bool = false
	var finalSliceSize int = bitstringSliceLength

	if bitstringSliceLength == 0 {
		bitstringSliceLength = 1
		finalSliceSize = 1
	}

	if residualCharacters > 0 && longStringLength > 16 {
		hasResidualString = true
		finalSliceSize++
	}

	bitstringOutputBuffer = make([]Bitstring, finalSliceSize)

	for sliceIndex := 0; sliceIndex < finalSliceSize || hasResidualString; sliceIndex++ {
		initialSliceIndex = 16 * sliceIndex
		finalSliceIndex = 16 * (sliceIndex + 1)

		if sliceIndex == (finalSliceSize-1) && hasResidualString {
			finalSliceIndex = (bitstringSliceLength * 16) + residualCharacters

			hasResidualString = false
		}

		if finalSliceSize == 1 {
			initialSliceIndex = 0
			finalSliceIndex = longStringLength
		}

		_128BString = longString[initialSliceIndex:finalSliceIndex]
		_128BBitstring = bitstring.FromStringToBitstring(_128BString)

		bitstringOutputBuffer[sliceIndex] = _128BBitstring
	}

	return bitstringOutputBuffer
}

func (bitstring *Bitstring) FromLongBitstringToBitSlice(longBitstring Bitstring) BitSlice {
	var longBitstringLength = len(longBitstring)
	var bitSliceLength int = longBitstringLength / 128
	var _128Bitstring Bitstring
	var outputBitSlice BitSlice
	var initialLimit int = 0
	var endLimit int = 0
	var hasResidual bool

	fmt.Println(longBitstring)

	if (longBitstringLength % 128) > 0 {
		bitSliceLength++
		hasResidual = true
	}

	outputBitSlice = make([]Bitstring, bitSliceLength)

	for sliceIndex := 0; sliceIndex < bitSliceLength; sliceIndex++ {
		initialLimit = sliceIndex * 128
		endLimit = (sliceIndex + 1) * 128

		if hasResidual && endLimit > longBitstringLength {
			_128Bitstring = longBitstring[initialLimit:longBitstringLength]
			_128Bitstring.Pad(128 - len(_128Bitstring))
		} else {
			_128Bitstring = longBitstring[initialLimit:endLimit]
		}

		outputBitSlice[sliceIndex] = _128Bitstring
	}

	return outputBitSlice
}

func (bitstring *Bitstring) FromBitstringSliceToString(slice BitSlice) string {
	var stringBuilder strings.Builder
	var sliceLength int = len(slice)
	var currentBitstring Bitstring
	var stringDecoded string

	for sliceIndex := 0; sliceIndex < sliceLength; sliceIndex++ {
		currentBitstring = slice[sliceIndex]
		stringDecoded = bitstring.String(currentBitstring)
		stringBuilder.WriteString(stringDecoded)
	}

	return stringBuilder.String()
}

func (bitstring *Bitstring) FromIntToBitstring(number int64, characterLength int) (bitstringTransformed Bitstring) {
	if characterLength < 1 {
		fmt.Println("One bitstring must have at least one character")
	}

	if number < 0 {
		fmt.Println("One bitstring can't have a negative representation")
	}

	for number > 0 {
		if number&1 == 1 {
			bitstringTransformed += "1"
		} else {
			bitstringTransformed += "0"
		}

		number = number >> 1
	}

	for len(bitstringTransformed) < characterLength {
		bitstringTransformed += "0"
	}

	return bitstringTransformed
}

func (bitstring Bitstring) FromBitstringToByteSlice() (sliceTransformed []byte) {
	for _, bit := range bitstring {
		if string(bit) == "0" {
			sliceTransformed = append(sliceTransformed, '0')
		} else {
			sliceTransformed = append(sliceTransformed, '1')
		}
	}

	return sliceTransformed
}

func (bitstring *Bitstring) FromByteSliceToBitstring(byteSlice []byte) (outputBitstring Bitstring) {
	for _, byte := range byteSlice {
		if byte == 49 {
			outputBitstring += "1"
		} else {
			outputBitstring += "0"
		}
	}

	return outputBitstring
}

func (bitstring Bitstring) FromBitstringOctectToHex() (hexstringTransformed HexString) {
	if len(bitstring) > 4 {
		fmt.Println("Bitstring with length of more than four characters cannot be converted into hex string")
	}

	return BINARY_TO_HEX_TABLE[bitstring]
}

func (bitstring Bitstring) FromBitstringToHexString() (hexstringTransformed HexString) {
	var bitstringLength = len(bitstring)
	var bitslice BitSlice
	var transformedBlock = make([]byte, 4)

	for octectIndex := 0; octectIndex < bitstringLength; bitstringLength += 4 {
		transformedBlock[0] = bitstring[octectIndex]
		transformedBlock[1] = bitstring[octectIndex+1]
		transformedBlock[2] = bitstring[octectIndex+2]
		transformedBlock[3] = bitstring[octectIndex+3]
		bitslice = append(bitslice, Bitstring(transformedBlock))
	}

	for _, bitstringFromSlice := range bitslice {
		hexstringTransformed += bitstringFromSlice.FromBitstringOctectToHex()
	}

	return hexstringTransformed
}

func (bitstring *Bitstring) FromHexToBitstringOctect(hexstring HexString) Bitstring {
	if len(hexstring) > 1 {
		fmt.Println("An hexstring with length of more than one character cannot be converted into 4-bit bitstring")
	}

	return HEX_TO_BINARY_TABLE[hexstring]
}

func (bitstring *Bitstring) FromBitstringToUInt8(character Bitstring) uint8 {
	var decimalCharValue uint8 = 0
	var bit uint8

	for bitIndex := 0; bitIndex < 8; bitIndex++ {
		if character[bitIndex] == 49 {
			bit = 1
		} else {
			bit = 0
		}
		decimalCharValue += uint8(math.Pow(2, float64(bitIndex))) * bit
	}

	return decimalCharValue
}

func (bitstring *Bitstring) String(bitstringToConvert Bitstring) (asciiString string) {
	var bitstringChar Bitstring
	var charDecimalValue uint8

	if len(bitstringToConvert) < 128 {
		bitstringToConvert.Pad(128 - len(bitstringToConvert))
	}

	for bitIndex := 0; bitIndex < 16; bitIndex++ {
		bitstringChar = bitstringToConvert[bitIndex*8 : (bitIndex+1)*8]

		if bitstringChar == BITSTIRNG_EMPTY_BLOCK {
			continue
		}

		charDecimalValue = bitstring.FromBitstringToUInt8(bitstringChar)
		asciiString += string(charDecimalValue)
	}

	return string(asciiString)
}
