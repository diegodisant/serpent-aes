package serpent

import t "serpent_aes/types"

const SERPENT_PHI int64 = 0x9e3779b9
const SERPENT_ROUNDS int = 32

var BitstringStatusBox []map[t.Bitstring]t.Bitstring
var BitstringStatusBoxInverse []map[t.Bitstring]t.Bitstring

type Serpent struct{}

func init() {
	var bitstringDictionary map[t.Bitstring]t.Bitstring
	var bitstringInverseDictionary map[t.Bitstring]t.Bitstring
	var statusBoxLength int

	for _, statusBox := range StatusBoxDecimalTable {
		statusBoxLength = len(statusBox)
		bitstringDictionary = make(map[t.Bitstring]t.Bitstring, statusBoxLength)
		bitstringInverseDictionary = make(map[t.Bitstring]t.Bitstring, statusBoxLength)

		makeDictionaries(statusBox, &bitstringDictionary, &bitstringInverseDictionary)

		BitstringStatusBox = append(BitstringStatusBox, bitstringDictionary)
		BitstringStatusBoxInverse = append(BitstringStatusBoxInverse, bitstringInverseDictionary)
	}
}

func makeDictionaries(statusBox t.StateBox, dictionary *map[t.Bitstring]t.Bitstring, inverseDictionary *map[t.Bitstring]t.Bitstring) {
	var bitstring t.Bitstring
	var statusBoxBitstringIndex t.Bitstring
	var statusBoxBitStringValue t.Bitstring

	for boxIndex, boxValue := range statusBox {
		statusBoxBitstringIndex = bitstring.FromIntToBitstring(int64(boxIndex), 4)
		statusBoxBitStringValue = bitstring.FromIntToBitstring(int64(boxValue), 4)

		(*dictionary)[statusBoxBitstringIndex] = statusBoxBitStringValue
		(*inverseDictionary)[statusBoxBitStringValue] = statusBoxBitstringIndex
	}
}
