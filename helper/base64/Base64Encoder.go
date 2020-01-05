package base64

import b64 "encoding/base64"

func EncodeString(normalString string) string {
	return b64.StdEncoding.EncodeToString([]byte(normalString))
}

func DecodeString(stringEncoded string) string {
	var decodedBuffer []byte

	decodedBuffer, _ = b64.StdEncoding.DecodeString(stringEncoded)

	return string(decodedBuffer)
}
