package block

import helper_ascii "serpent_aes/helper/ascii"

var ASCII_PRINTABLE_CHARS map[uint8]string = helper_ascii.GeneratePrintableAsciiString()

var CHAR_SHIFTS = []uint{
	1 << 1: 1,
	1 << 2: 2,
	1 << 3: 3,
	1 << 4: 4,
	1 << 5: 5,
}
