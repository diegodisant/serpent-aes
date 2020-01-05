package types

import "fmt"

type BitstringXorOperator interface {
	BinaryXor(secondBitstring Bitstring) Bitstring
	BitSliceXor(slice BitSlice) Bitstring
}

func (bitstring Bitstring) BinaryXor(secondBitstring Bitstring) Bitstring {
	if len(bitstring) != len(secondBitstring) {
		fmt.Println("Cannot make binary XOR with two bitstrings of different lengths")
	}

	var outputBitstringWithXorApplied Bitstring = ""

	for bitIndex, bit := range bitstring {
		if uint8(bit) == secondBitstring[bitIndex] {
			outputBitstringWithXorApplied = outputBitstringWithXorApplied + "0"
		} else {
			outputBitstringWithXorApplied = outputBitstringWithXorApplied + "1"
		}
	}

	return outputBitstringWithXorApplied
}

func (bitstring Bitstring) BitSliceXor(slice BitSlice) Bitstring {
	if len(slice) == 0 {
		fmt.Println("At least one bit string is needed to make xor operation")
	}

	var outputBitstringWithXorBitSliceApplied Bitstring = slice[0]

	for _, bitstringInSlice := range slice[1:] {
		outputBitstringWithXorBitSliceApplied = outputBitstringWithXorBitSliceApplied.BinaryXor(bitstringInSlice)
	}

	return outputBitstringWithXorBitSliceApplied
}
