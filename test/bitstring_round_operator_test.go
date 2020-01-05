package test

import "testing"
import t "serpent_aes/types"
import s "serpent_aes/serpent"

func TestBitstringRoundOperator(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var targetKeyWithRoundsApplied t.Bitstring = "00010011010001101100000010110000" +
		"10001110011011000010011010011101" +
		"01110100011100101010110011100100" +
		"01010001110011101011110010101110"

	_, subKeysHat := serpent.MakeSubKeys(serpent.MakeLongKey(bitstring))

	var outputRoundApplied t.Bitstring = bitstring

	for roundIndex := 0; roundIndex < s.SERPENT_ROUNDS; roundIndex++ {
		outputRoundApplied = serpent.ApplyRound(roundIndex, outputRoundApplied, subKeysHat)
	}

	if outputRoundApplied != targetKeyWithRoundsApplied {
		test.Errorf("Output keys with round applied doesn't match target (%s) != (%s)", outputRoundApplied, targetKeyWithRoundsApplied)
		test.Fail()
	}
}

func TestBitstringInverseRoundOperator(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var targetKeyWithRoundsApplied t.Bitstring = "00010011010001101100000010110000" +
		"10001110011011000010011010011101" +
		"01110100011100101010110011100100" +
		"01010001110011101011110010101110"

	_, subKeysHat := serpent.MakeSubKeys(serpent.MakeLongKey(bitstring))

	var outputRoundApplied t.Bitstring = targetKeyWithRoundsApplied

	for roundIndex := s.SERPENT_ROUNDS - 1; roundIndex >= 0; roundIndex-- {
		outputRoundApplied = serpent.ApplyInverseRound(roundIndex, outputRoundApplied, subKeysHat)
	}

	if outputRoundApplied != bitstring {
		test.Errorf("Output keys with round applied doesn't match target (%s) != (%s)", outputRoundApplied, targetKeyWithRoundsApplied)
		test.Fail()
	}
}
