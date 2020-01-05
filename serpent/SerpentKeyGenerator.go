package serpent

import "fmt"
import t "serpent_aes/types"
import transform "serpent_aes/transformation"

type SerpentKeyGenerator interface {
	MakeSubKeys(userKey t.Bitstring) (t.BitSlice, t.BitSlice)
	MakeLongKey(key t.Bitstring) t.Bitstring
}

func (serpent *Serpent) MakeSubKeys(userKey t.Bitstring) (t.BitSlice, t.BitSlice) {
	var keyTo32BitWords = make(t.BitMap, 132)
	var temporaryBitstring t.Bitstring
	var temporarySlice t.BitSlice
	var subKeys = make(t.BitSlice, 132)
	var selectedStatusBox int = 0
	var inputBitstring t.Bitstring
	var outputBitstring t.Bitstring
	var keySlice t.BitSlice = t.BitSlice{}
	var keyWithPermutationSlice t.BitSlice = t.BitSlice{}

	for _32BitHop := -8; _32BitHop < 0; _32BitHop++ {
		keyTo32BitWords[_32BitHop] = userKey[(_32BitHop+8)*32 : (_32BitHop+9)*32]
	}

	for tableIndex := 0; tableIndex < 132; tableIndex++ {
		temporarySlice = t.BitSlice{
			keyTo32BitWords[tableIndex-8],
			keyTo32BitWords[tableIndex-5],
			keyTo32BitWords[tableIndex-3],
			keyTo32BitWords[tableIndex-1],
			temporaryBitstring.FromIntToBitstring(SERPENT_PHI, 32),
			temporaryBitstring.FromIntToBitstring(int64(tableIndex), 32),
		}

		temporaryBitstring = temporaryBitstring.BitSliceXor(temporarySlice)

		keyTo32BitWords[tableIndex] = temporaryBitstring.RotateLeft(11)
	}

	for round := 0; round < SERPENT_ROUNDS+1; round++ {
		selectedStatusBox = (SERPENT_ROUNDS + 3 - round) % SERPENT_ROUNDS

		subKeys[0+4*round] = ""
		subKeys[1+4*round] = ""
		subKeys[2+4*round] = ""
		subKeys[3+4*round] = ""

		for byteIndex := 0; byteIndex < 32; byteIndex++ {
			inputBitstring = t.Bitstring(keyTo32BitWords[0+4*round][byteIndex]) +
				t.Bitstring(keyTo32BitWords[1+4*round][byteIndex]) +
				t.Bitstring(keyTo32BitWords[2+4*round][byteIndex]) +
				t.Bitstring(keyTo32BitWords[3+4*round][byteIndex])

			outputBitstring = serpent.ApplyStatusBox(selectedStatusBox, inputBitstring)

			for _32BitHop := 0; _32BitHop < 4; _32BitHop++ {
				subKeys[_32BitHop+4*round] = subKeys[_32BitHop+4*round] + t.Bitstring(outputBitstring[_32BitHop])
			}
		}
	}

	for keyIndex := 0; keyIndex < 33; keyIndex++ {
		keySlice = append(keySlice, subKeys[4*keyIndex]+subKeys[4*keyIndex+1]+subKeys[4*keyIndex+2]+subKeys[4*keyIndex+3])
	}

	for keyIndex := 0; keyIndex < 33; keyIndex++ {
		keyWithPermutationSlice = append(keyWithPermutationSlice, transform.MakeInitialPermutation(keySlice[keyIndex]))
	}

	return keySlice, keyWithPermutationSlice
}

func (serpent *Serpent) MakeLongKey(key t.Bitstring) t.Bitstring {
	var keyLength int = len(key)

	if keyLength%32 != 0 || keyLength < 64 || keyLength > 256 {
		fmt.Printf("Invalid key length (%d bits) \n", keyLength)
	} else if keyLength == 256 {
		return key
	}

	for keyBitIndex := 0; keyBitIndex < 256-keyLength; keyBitIndex++ {
		if keyBitIndex == 0 {
			key += "1"
		} else {
			key += "0"
		}
	}

	return key
}
