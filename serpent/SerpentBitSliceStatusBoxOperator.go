package serpent

import t "serpent_aes/types"

type SerpentBitSliceStatusBoxOperator interface {
	ApplyStatusBoxToBitSlice(statusBox int, slice t.BitSlice) t.BitSlice
	ApplyInverseStatusBoxToBitSlice(statusBox int, slice t.BitSlice) t.BitSlice
}

func (serpent *Serpent) ApplyStatusBoxToBitSlice(statusBox int, slice t.BitSlice) t.BitSlice {
	var sliceTransformed = make(t.BitSlice, 4)
	var _32bit_bitstring0 t.Bitstring
	var _32bit_bitstring1 t.Bitstring
	var _32bit_bitstring2 t.Bitstring
	var _32bit_bitstring3 t.Bitstring
	var quadJoinedBitstring t.Bitstring

	for bitIndex := 0; bitIndex < 32; bitIndex++ {
		_32bit_bitstring0 = t.Bitstring(int(slice[0][bitIndex]))
		_32bit_bitstring1 = t.Bitstring(int(slice[1][bitIndex]))
		_32bit_bitstring2 = t.Bitstring(int(slice[2][bitIndex]))
		_32bit_bitstring3 = t.Bitstring(int(slice[3][bitIndex]))

		quadJoinedBitstring = serpent.ApplyStatusBox(statusBox, _32bit_bitstring0+_32bit_bitstring1+_32bit_bitstring2+_32bit_bitstring3)

		serpent.appendQuadBitstringToOutputSlice(&sliceTransformed, quadJoinedBitstring)
	}

	return sliceTransformed
}

func (serpent *Serpent) ApplyInverseStatusBoxToBitSlice(statusBox int, slice t.BitSlice) t.BitSlice {
	var sliceTransformed = make(t.BitSlice, 4)
	var _32bit_bitstring0 t.Bitstring
	var _32bit_bitstring1 t.Bitstring
	var _32bit_bitstring2 t.Bitstring
	var _32bit_bitstring3 t.Bitstring
	var quadJoinedBitstring t.Bitstring

	for bitIndex := 0; bitIndex < 32; bitIndex++ {
		_32bit_bitstring0 = t.Bitstring(int(slice[0][bitIndex]))
		_32bit_bitstring1 = t.Bitstring(int(slice[1][bitIndex]))
		_32bit_bitstring2 = t.Bitstring(int(slice[2][bitIndex]))
		_32bit_bitstring3 = t.Bitstring(int(slice[3][bitIndex]))

		quadJoinedBitstring = serpent.ApplyInverseStatusBox(statusBox, _32bit_bitstring0+_32bit_bitstring1+_32bit_bitstring2+_32bit_bitstring3)

		serpent.appendQuadBitstringToOutputSlice(&sliceTransformed, quadJoinedBitstring)
	}

	return sliceTransformed
}

func (serpent *Serpent) appendQuadBitstringToOutputSlice(outputSlice *t.BitSlice, quadJoinedBitstring t.Bitstring) {
	for _32bitIndex := 0; _32bitIndex < 4; _32bitIndex++ {
		(*outputSlice)[_32bitIndex] += t.Bitstring(int(quadJoinedBitstring[_32bitIndex]))
	}
}
