package serpent

import t "serpent_aes/types"

type SerpentBinaryEncoder interface {
	EncryptBinary(binaryContent []byte, userKey string) []byte
	DecryptBinary(binaryEncryptedContent []byte, userKey string) []byte
}

func (serpent *Serpent) EncryptBinary(binaryContent []byte, userKey string) []byte {
	var bitstring t.Bitstring
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)
	var longBitstringToEncrypt t.Bitstring = bitstring.FromBinaryByteSliceToBitstring(binaryContent)
	var sliceToEncrypt t.BitSlice = bitstring.FromLongBitstringToBitSlice(longBitstringToEncrypt)
	var sliceSize int = len(sliceToEncrypt)
	var bitstringToEncrypt t.Bitstring
	var bitstringEncrypted t.Bitstring
	var sliceEncrypted t.BitSlice

	userKeyEncoded = serpent.MakeLongKey(userKeyEncoded)
	sliceEncrypted = make([]t.Bitstring, sliceSize)

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		bitstringToEncrypt = sliceToEncrypt[sliceIndex]
		bitstringEncrypted = serpent.Encrypt(bitstringToEncrypt, userKeyEncoded)
		sliceEncrypted = append(sliceEncrypted, bitstringEncrypted)
	}

	return bitstring.FromBitSliceToByteSlice(sliceEncrypted)
}

func (serpent *Serpent) DecryptBinary(binaryEncryptedContent []byte, userKey string) []byte {
	var bitstring t.Bitstring
	var userKeyEncoded t.Bitstring = bitstring.FromStringToBitstring(userKey)
	var longBitstringToDecrypt t.Bitstring = bitstring.FromBinaryByteSliceToBitstring(binaryEncryptedContent)
	var sliceToDecrypt t.BitSlice = bitstring.FromLongBitstringToBitSlice(longBitstringToDecrypt)
	var sliceSize int = len(sliceToDecrypt)
	var bitstringToDecrypt t.Bitstring
	var bitstringDecrypted t.Bitstring
	var sliceDecrypted t.BitSlice

	userKeyEncoded = serpent.MakeLongKey(userKeyEncoded)
	sliceDecrypted = make([]t.Bitstring, sliceSize)

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		bitstringToDecrypt = sliceToDecrypt[sliceIndex]
		bitstringDecrypted = serpent.Decrypt(bitstringToDecrypt, userKeyEncoded)
		sliceDecrypted = append(sliceDecrypted, bitstringDecrypted)
	}

	return bitstring.FromBitSliceToByteSlice(sliceDecrypted)
}
