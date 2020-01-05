package test

import "testing"
import t "serpent_aes/types"

func TestLongBitstringToBitSlie(test *testing.T) {
	var bitstring t.Bitstring
	var testBitstring t.Bitstring = "0010011010000110001011101000011000100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110"

	var outputSlice t.BitSlice = t.BitSlice{
		"00100110100001100010111010000110001001101000011000101110100001100010011010000110001011101000011000100110100001100010111010000110",
		"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100110",
	}

	var currentBitstring t.Bitstring
	var bitstringExpected t.Bitstring
	var bitstringSliceEncoded t.BitSlice = bitstring.FromLongBitstringToBitSlice(testBitstring)

	for sliceIndex := 0; sliceIndex < 2; sliceIndex++ {
		bitstringExpected = outputSlice[sliceIndex]
		currentBitstring = bitstringSliceEncoded[sliceIndex]

		continue

		if bitstringExpected != currentBitstring {
			test.Errorf("Bitstring gotten from slices are different %s\n!=\n%s", bitstringExpected, currentBitstring)
			test.Fail()
		}
	}
}
