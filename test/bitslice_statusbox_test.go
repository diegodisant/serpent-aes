package test

import "testing"
import s "serpent_aes/serpent"
import t "serpent_aes/types"

func TestBitSliceStatusBoxOperator(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring = "00100100011001010110001101110010" +
		"01100101011101000111010001100101" +
		"00100100011101000110010001100001" +
		"01110100011010110110010101111001"

	var testedBitslice t.BitSlice = serpent.ApplyStatusBoxToBitSlice(2, bitstring.QuadSplit())
	var testedInverseBitslice t.BitSlice = serpent.ApplyInverseStatusBoxToBitSlice(2, testedBitslice)
	var bitstringJoined t.Bitstring = bitstring.QuadJoin(testedInverseBitslice)

	if bitstringJoined != bitstring {
		test.Errorf("Bit slice operator on 32-bit bitstring pieces is not decoding status boxes correctly, please debug")
	}
}
