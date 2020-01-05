package serpent_aes

import s "serpent_aes/serpent"

func SerpentEncrypt(plainText string, key string) string {
	var serpent s.Serpent

	return serpent.EncryptText(plainText, key)
}

func SerpentDecrypt(encryptedText string, key string) string {
	var serpent s.Serpent

	return serpent.DecryptText(encryptedText, key)
}
