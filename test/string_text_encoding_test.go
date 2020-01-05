package test

import (
	s "serpent_aes/serpent"
	t "serpent_aes/types"
	"testing"

	helper_b64 "serpent_aes/helper/base64"
)

func TestEncryptString(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring
	var targetCipherText string = "TURFeE1EQXdNVEF3TURFeE1URXdNVEV4TVRBeE1ERXdNREF4TVRBeE1ERXdNREF3TURFeE1URXdNREF3TURBd01EQXdNVEF3TURFd01URXhNREF3TVRFd01URXdNVEV4TVRBd01ERXhNREV4TURFd01EQXdNREF4TVRBd01ERXhNREV3TURFd01UQXdNVEV3TVRBd01ERXdNREF3TVRFd01URXhNREV3TVRFPQ=="
	var plainText string = "datadatadatadata" //16 byte perfect message
	var userKey string = "$3cr3tk3y"

	// make a bitstring representation of components for encryption
	var dataEncoded t.Bitstring = bitstring.FromStringToBitstring(plainText)
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)

	// encrypt text
	var encryptedText t.Bitstring = serpent.Encrypt(dataEncoded, serpent.MakeLongKey(userKeyEncoded))

	//create a base 64 from bitstring representation
	var encryptedTextBase64String string = bitstring.FromBitstringToBase64(encryptedText)
	var doubleEncryptedBase64String string = helper_b64.EncodeString(encryptedTextBase64String)

	if targetCipherText != doubleEncryptedBase64String {
		test.Errorf("Encrypted texts doesn't match (%s) != (%s)", targetCipherText, doubleEncryptedBase64String)
		test.Fail()
	}
}

func TestDecryptString(test *testing.T) {
	var serpent s.Serpent
	var bitstring t.Bitstring
	var targetDecryptedText string = "datadatadatadata"
	var cipherText string = "TURFeE1EQXdNVEF3TURFeE1URXdNVEV4TVRBeE1ERXdNREF4TVRBeE1ERXdNREF3TURFeE1URXdNREF3TURBd01EQXdNVEF3TURFd01URXhNREF3TVRFd01URXdNVEV4TVRBd01ERXhNREV4TURFd01EQXdNREF4TVRBd01ERXhNREV3TURFd01UQXdNVEV3TVRBd01ERXdNREF3TVRFd01URXhNREV3TVRFPQ=="
	var userKey string = "$3cr3tk3y"

	// make a bitstring representation of encrypted data
	var base64DataDecoded string = helper_b64.DecodeString(cipherText)
	var dataEncoded t.Bitstring = bitstring.FromBase64ToBitstring(base64DataDecoded)
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)

	var decryptedText t.Bitstring = serpent.Decrypt(dataEncoded, serpent.MakeLongKey(userKeyEncoded))
	var decryptedTextString string = bitstring.String(decryptedText)

	if targetDecryptedText != decryptedTextString {
		test.Errorf("Decrypted texts doesn't match (%s) != (%s)", targetDecryptedText, decryptedTextString)
		test.Fail()
	}
}
