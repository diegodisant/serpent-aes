package test

import "testing"
import t "serpent_aes/types"
import s "serpent_aes/serpent"

func TestBitstringEncrypt(test *testing.T) {
	var serpent s.Serpent

	var plainTextKey t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var plainText t.Bitstring = "10011010011010001101110001110100" +
		"10100101010010101001110001010010" +
		"10010100000001111111010010111111" +
		"00000110101001110011001010110010"

	var encryptedExpectedText t.Bitstring = "11111101110001101000110011011111" +
		"01001100001111010101101101101100" +
		"11010100010001001100111010010110" +
		"11010111001011010111011110100101"

	var encryptedText t.Bitstring = serpent.Encrypt(plainText, serpent.MakeLongKey(plainTextKey))

	if encryptedText != encryptedExpectedText {
		test.Errorf("Encrypted texts doesn't match (%s) != (%s)", encryptedExpectedText, encryptedText)
		test.Fail()
	}
}

func TestBitstringDecrypt(test *testing.T) {
	var serpent s.Serpent

	var plainTextKey t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var decryptedExpectedText t.Bitstring = "10011010011010001101110001110100" +
		"10100101010010101001110001010010" +
		"10010100000001111111010010111111" +
		"00000110101001110011001010110010"

	var encryptedText t.Bitstring = "11111101110001101000110011011111" +
		"01001100001111010101101101101100" +
		"11010100010001001100111010010110" +
		"11010111001011010111011110100101"

	var decryptedText t.Bitstring = serpent.Decrypt(encryptedText, serpent.MakeLongKey(plainTextKey))

	if decryptedText != decryptedExpectedText {
		test.Errorf("Decrypted texts doesn't match (%s) != (%s)", decryptedText, decryptedExpectedText)
		test.Fail()
	}
}
