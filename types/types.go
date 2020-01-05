package types

import helper_ascii "serpent_aes/helper/ascii"

var BINARY_TO_HEX_TABLE map[Bitstring]HexString = map[Bitstring]HexString{
	"0000": "0",
	"1000": "1",
	"0100": "2",
	"1100": "3",
	"0010": "4",
	"1010": "5",
	"0110": "6",
	"1110": "7",
	"0001": "8",
	"1001": "9",
	"0101": "a",
	"1101": "b",
	"0011": "c",
	"1011": "d",
	"0111": "e",
	"1111": "f",
}

var HEX_TO_BINARY_TABLE map[HexString]Bitstring = map[HexString]Bitstring{
	"0": "0000",
	"1": "1000",
	"2": "0100",
	"3": "1100",
	"4": "0010",
	"5": "1010",
	"6": "0110",
	"7": "1110",
	"8": "0001",
	"9": "1001",
	"a": "0101",
	"b": "1101",
	"c": "0011",
	"d": "1011",
	"e": "0111",
	"f": "1111",
}

var ASCII_PRINTABLE_CHARS map[uint8]string = helper_ascii.GeneratePrintableAsciiString()
