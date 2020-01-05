package test

import "testing"
import t "serpent_aes/types"
import s "serpent_aes/serpent"

func TestSerpentKeyGeneration(test *testing.T) {
	var serpent s.Serpent
	var userKey t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var longKey t.Bitstring = "110011100101001100101001010100101011100010" +
		"0110010100101001010111001010010101011001001001010110101001" +
		"0100101010010101001010100101100000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000"

	//expected subkeys
	var subKey3 t.Bitstring = "010001101101011101010100110100011111110001011001" +
		"00000011100010100011001010110000000101111110001000000011001" +
		"011011001011010111011"

	var subKey13 t.Bitstring = "11010100111111001101000010100101110011000000011" +
		"00010110100110010101000011111111010101100110011010111101101" +
		"0111110010010011101011"

	var subKeyHat3 t.Bitstring = "010011000110011001001100101100011010110000111" +
		"1100101100110001101000110000000101100001011011101101111101" +
		"0001110010101000001111001"

	var subKeyHat15 t.Bitstring = "10110100110111001000100111011110010110001010" +
		"01011101101011110001110101101010000001001101010110100110001" +
		"1000100001011101000110101"

	var expectedKey t.BitSlice
	var expectedKeyHat t.BitSlice

	var userLongKey t.Bitstring = serpent.MakeLongKey(userKey)
	var userLongKeyLength int = len(userLongKey)

	if userLongKeyLength == 256 {
		if userLongKey != longKey {
			test.Errorf("User generated long key doesn't match with user long key (%s) != (%s)", userLongKey, userKey)
			test.Fail()
		}

		expectedKey, expectedKeyHat = serpent.MakeSubKeys(userLongKey)
	}

	if len(expectedKey) != 33 && len(expectedKeyHat) != 33 {
		test.Errorf("Expected keys doens't have a length of 33 slices")
		test.Fail()
	}

	if expectedKey[3] != subKey3 {
		test.Errorf("Sub key in position 3 doesn't match with expected output (%s) != (%s)", expectedKey[3], subKey3)
		test.Fail()
	}

	if expectedKey[13] != subKey13 {
		test.Errorf("Sub key in position 13 doesn't match with expected output (%s) != (%s)", expectedKey[13], subKey13)
		test.Fail()
	}

	if expectedKeyHat[3] != subKeyHat3 {
		test.Errorf("Sub key hat in position 3 doesn't match with expected output (%s) != (%s)", expectedKeyHat[3], subKeyHat3)
		test.Fail()
	}

	if expectedKeyHat[15] != subKeyHat15 {
		test.Errorf("Sub key hat in position 15 doesn't match with expected output (%s) != (%s)", expectedKeyHat[15], subKeyHat15)
		test.Fail()
	}
}
