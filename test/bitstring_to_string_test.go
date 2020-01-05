package test

import "testing"
import t "serpent_aes/types"

func TestBitstringToStringCasting(test *testing.T) {
	var bitstring t.Bitstring
	var targetBitstring t.Bitstring = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100110100001100010111010000110"
	var plainText string = "data"
	var decodedBitstring string = bitstring.String(targetBitstring)

	if decodedBitstring != plainText {
		test.Errorf("Plain text doesn't match with enconded-string bitstring (%s) != (%s)", plainText, decodedBitstring)
		test.Fail()
	}
}
