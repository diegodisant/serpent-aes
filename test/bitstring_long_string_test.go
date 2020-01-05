package test

import "testing"
import t "serpent_aes/types"

func TestBitstringLongStringText(test *testing.T) {
	var bitstring t.Bitstring
	var testString string = "datadatadatadatadatadatadatadata" + "datadatadatadatadatadatadatadata"
	var outputSlice t.BitSlice = t.BitSlice{
		"00100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110100001100010111010000110",
		"00100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110100001100010111010000110",
		"00100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110100001100010111010000110",
		"00100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110100001100010111010000110",
	}
	var bitstringExpected t.Bitstring
	var bitstringSliceEncoded t.BitSlice = bitstring.FromLongStringToBitSlice(testString)
	var currentBitstring t.Bitstring

	for sliceIndex := 0; sliceIndex < 4; sliceIndex++ {
		bitstringExpected = outputSlice[sliceIndex]
		currentBitstring = bitstringSliceEncoded[sliceIndex]

		if bitstringExpected != currentBitstring {
			test.Errorf("Bitstring gotten from slices are different %s\n!=\n%s", bitstringExpected, currentBitstring)
			test.Fail()
		}
	}
}
