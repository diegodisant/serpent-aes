package serpent

import t "serpent_aes/types"

type SerpentBitSliceEncoder interface {
	EncryptBitSlice(plainText t.Bitstring, userKey t.Bitstring) t.Bitstring
	DecryptBitSlice(cipherText t.Bitstring, userKey t.Bitstring) t.Bitstring
}

func (serpent *Serpent) EncryptBitSlice(plainText t.Bitstring, userKey t.Bitstring) t.Bitstring {
	var encryptKey t.BitSlice
	var cipherText t.Bitstring

	encryptKey, _ = serpent.MakeSubKeys(userKey)
	cipherText = plainText

	for round := 0; round < SERPENT_ROUNDS; round++ {
		cipherText = serpent.ApplyBitSliceRound(round, cipherText, encryptKey)
	}

	return cipherText
}

func (serpent *Serpent) DecryptBitSlice(cipherText t.Bitstring, userKey t.Bitstring) t.Bitstring {
	var decryptKey t.BitSlice
	var plainText t.Bitstring

	decryptKey, _ = serpent.MakeSubKeys(userKey)
	plainText = cipherText

	for round := SERPENT_ROUNDS - 1; round >= 0; round-- {
		plainText = serpent.ApplyBitSliceInverseRound(round, plainText, decryptKey)
	}

	return plainText
}
