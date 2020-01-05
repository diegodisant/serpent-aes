package test

import "testing"
import s "serpent_aes/serpent"
import t "serpent_aes/types"

func TestParallelStatusBoxHat(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring = "00100100011001010110001101110010" +
		"01100101011101000111010001100101" +
		"00100100011101000110010001100001" +
		"01110100011010110110010101111001"

	var testedBitstring t.Bitstring = serpent.ApplyStatusBoxHat(3, bitstring)
	var originalBitstring t.Bitstring = serpent.ApplyInverseStatusBoxHat(3, testedBitstring)

	if originalBitstring != bitstring {
		test.Errorf("bitstrings doesn't match with application of parallel copy (%s) != (%s)", originalBitstring, testedBitstring)
		test.Fail()
	}
}
