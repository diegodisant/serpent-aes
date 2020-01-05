package test

import "testing"
import t "serpent_aes/types"

func TestStringToBitstringCasting(test *testing.T) {
	var bitstring t.Bitstring
	var targetBitstring t.Bitstring = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100110100001100010111010000110"
	var plainText string = "data"
	var encodedBitstring t.Bitstring = bitstring.FromStringToBitstring(plainText)

	if targetBitstring != encodedBitstring {
		test.Errorf("Target bitstring doesn't match with encoded bitstring (%s) != (%s)", targetBitstring, encodedBitstring)
		test.Fail()
	}
}
