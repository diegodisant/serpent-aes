package ascii

func GeneratePrintableAsciiString() map[uint8]string {
	var printableCharsMap map[uint8]string = make(map[uint8]string)

	for charDecValue := 32; charDecValue < 127; charDecValue++ {
		printableCharsMap[uint8(charDecValue)] = string(charDecValue)
	}

	return printableCharsMap
}
