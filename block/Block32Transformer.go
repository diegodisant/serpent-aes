package block

type Block32Transformer interface {
	FromUInt32ToBlock32(words []byte) Block32
}

func (block Block32) FromUInt32ToBlock32(words []byte) Block32 {
	var block32OutputTransformed Block32

	for byteIndex, _ := range words {
		block32OutputTransformed[byteIndex] = words[byteIndex]
	}

	return block32OutputTransformed
}
