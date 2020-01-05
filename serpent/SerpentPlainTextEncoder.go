package serpent

import (
	b64_helper "serpent_aes/helper/base64"
	t "serpent_aes/types"
	"strings"
)

type SerpentPlainTextEncoder interface {
	EncryptText(plainText string, userKey string) string
	DecryptText(cipherText string, userKey string) string
}

func (serpent *Serpent) EncryptText(plainText string, userKey string) string {
	var bitstring t.Bitstring
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)
	var sliceToEncrypt t.BitSlice = bitstring.FromLongStringToBitSlice(plainText)
	var sliceSize int = len(sliceToEncrypt)
	var bitstringToEncrypt t.Bitstring
	var bitstringEncrypted t.Bitstring
	var stringEncrypted string
	var stringBuilder strings.Builder

	userKeyEncoded = serpent.MakeLongKey(userKeyEncoded)

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		bitstringToEncrypt = sliceToEncrypt[sliceIndex]
		bitstringEncrypted = serpent.Encrypt(bitstringToEncrypt, userKeyEncoded)
		stringEncrypted = bitstring.FromBitstringToBase64(bitstringEncrypted)
		stringEncrypted = b64_helper.EncodeString(stringEncrypted)

		stringBuilder.WriteString(stringEncrypted)
	}

	return stringBuilder.String()
}

func (serpent *Serpent) DecryptText(cipherText string, userKey string) string {
	var bitstring t.Bitstring
	var bitslice t.BitSlice
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)

	var sliceToDecrypt t.BitSlice = bitslice.FromLongBase64ToBitSlice(cipherText)
	var sliceSize int = len(sliceToDecrypt)

	var bitstringToDecrypt t.Bitstring
	var bitstringDecrypted t.Bitstring
	var stringDecrypted string
	var stringBuilder strings.Builder

	userKeyEncoded = serpent.MakeLongKey(userKeyEncoded)

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		bitstringToDecrypt = sliceToDecrypt[sliceIndex]
		bitstringDecrypted = serpent.Decrypt(bitstringToDecrypt, userKeyEncoded)
		stringDecrypted = bitstring.String(bitstringDecrypted)

		stringBuilder.WriteString(stringDecrypted)
	}

	return stringBuilder.String()
}
