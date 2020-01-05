package test

import "testing"
import t "serpent_aes/types"
import s "serpent_aes/serpent"

func TestBitSliceRoundOperator(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var expectedBitstring t.Bitstring = "11111000001100111001000110000100" +
		"01111000010100101110011001100101" +
		"00001101001001000001100011111101" +
		"10111101101110010110111010111010"

	subKeys, _ := serpent.MakeSubKeys(serpent.MakeLongKey(bitstring))

	var bitstringPart = serpent.ApplyBitSliceRound(2, bitstring, subKeys)

	if bitstringPart != expectedBitstring {
		test.Errorf("Round operator doesn't match with expected bitstring (%s) != (%s)", expectedBitstring, bitstringPart)
		test.Fail()
	}

	var originalBitstring t.Bitstring = serpent.ApplyBitSliceInverseRound(2, bitstringPart, subKeys)

	if originalBitstring != bitstring {
		test.Errorf("Round inverse operator doesn't match with original bitstring (%s) != (%s)", bitstring, originalBitstring)
		test.Fail()
	}
}
