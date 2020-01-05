package serpent

import (
	transform "serpent_aes/transformation"
	t "serpent_aes/types"
)

type SerpentBitstringEncoder interface {
	Encrypt(plainText t.Bitstring, userKey t.Bitstring) t.Bitstring
	Decrypt(cipherText t.Bitstring, userKey t.Bitstring) t.Bitstring
}

func (serpent *Serpent) Encrypt(plainText t.Bitstring, userKey t.Bitstring) t.Bitstring {
	var encryptKey t.BitSlice
	var plainTextPermutated t.Bitstring
	var cipherText t.Bitstring

	_, encryptKey = serpent.MakeSubKeys(userKey)
	plainTextPermutated = transform.MakeInitialPermutation(plainText)

	for round := 0; round < SERPENT_ROUNDS; round++ {
		plainTextPermutated = serpent.ApplyRound(round, plainTextPermutated, encryptKey)
	}

	cipherText = transform.MakeFinalPermutation(plainTextPermutated)

	return cipherText
}

func (serpent *Serpent) Decrypt(cipherText t.Bitstring, userKey t.Bitstring) t.Bitstring {
	var decryptKey t.BitSlice
	var cipherTextPermutated t.Bitstring
	var plainText t.Bitstring

	_, decryptKey = serpent.MakeSubKeys(userKey)
	cipherTextPermutated = transform.MakeFinalInversePermutation(cipherText)

	for round := SERPENT_ROUNDS - 1; round >= 0; round-- {
		cipherTextPermutated = serpent.ApplyInverseRound(round, cipherTextPermutated, decryptKey)
	}

	plainText = transform.MakeInitialInversePermutation(cipherTextPermutated)

	return plainText
}
