package test

import "testing"
import t "serpent_aes/types"
import transform "serpent_aes/transformation"

func TestBitstringSerpentPermutation(test *testing.T) {
	var bitstringInitialPermutation t.Bitstring
	var bitstringFinalPermutation t.Bitstring
	var bitstringInitialInversePermutation t.Bitstring
	var bitstringFinalInversePermutation t.Bitstring
	var bitstring t.Bitstring

	bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	bitstringInitialPermutation = transform.MakeInitialPermutation(bitstring)
	bitstringFinalPermutation = transform.MakeFinalPermutation(bitstringInitialPermutation)
	bitstringFinalInversePermutation = transform.MakeFinalInversePermutation(bitstringFinalPermutation)
	bitstringInitialInversePermutation = transform.MakeInitialInversePermutation(bitstringFinalInversePermutation)

	if bitstringInitialInversePermutation != bitstring {
		test.Errorf("Permutation is not working bitstrings are not equal (%s) != (%s)", bitstring, bitstringInitialInversePermutation)
		test.Fail()
	}

	var targetFinalPermutation t.Bitstring = "11000100111101000100010110110010" +
		"11100010000010110011101001001001" +
		"01011001100001011001000100100110" +
		"00110110101100110110011010011001"

	bitstringFinalPermutation = transform.MakeFinalPermutation(bitstring)

	if bitstringFinalPermutation != targetFinalPermutation {
		test.Errorf("Final permutation bitstrings doesn't match (%s) != (%s)", bitstringFinalPermutation, targetFinalPermutation)
		test.Fail()
	}

	// @todo check why this assertion is not working

	bitstringFinalInversePermutation = transform.MakeFinalInversePermutation(targetFinalPermutation)

	if bitstringFinalInversePermutation != bitstring {
		test.Errorf("Final inverse permutation is not equals than original bitstring.\n(%s) != (%s)", bitstringFinalInversePermutation, bitstring)
		test.Fail()
	}
}
