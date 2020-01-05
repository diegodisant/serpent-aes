package serpent

import transform "serpent_aes/transformation"
import t "serpent_aes/types"

type SerpentBitSliceRoundOperator interface {
	ApplyBitSliceRound(rounds int, bitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring
	ApplyBitSliceInverseRound(rounds int, bitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring
}

func (serpent *Serpent) ApplyBitSliceRound(rounds int, bitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring {
	var xoredBitstring t.Bitstring
	var outputBitstring t.Bitstring
	var slice t.BitSlice

	xoredBitstring = xoredBitstring.BitSliceXor(t.BitSlice{bitstring, subKeys[rounds]})
	slice = serpent.ApplyStatusBoxToBitSlice(rounds, xoredBitstring.QuadSplit())

	if rounds == SERPENT_ROUNDS-1 {
		outputBitstring = xoredBitstring.BitSliceXor(t.BitSlice{xoredBitstring.QuadJoin(slice), subKeys[SERPENT_ROUNDS]})
	} else {
		outputBitstring = xoredBitstring.QuadJoin(transform.LinearTransformationBitSlice(slice))
	}

	return outputBitstring
}

func (serpent *Serpent) ApplyBitSliceInverseRound(rounds int, bitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring {
	var xoredBitSlice t.BitSlice
	var bitstringInput t.Bitstring
	var temporaryBitstring t.Bitstring
	var slice t.BitSlice

	if rounds == SERPENT_ROUNDS-1 {
		temporaryBitstring = temporaryBitstring.BitSliceXor(t.BitSlice{bitstring, subKeys[SERPENT_ROUNDS]})
		slice = temporaryBitstring.QuadSplit()
	} else {
		slice = transform.LinearInverseTransformationBitSlice(bitstring.QuadSplit())
	}

	xoredBitSlice = serpent.ApplyInverseStatusBoxToBitSlice(rounds, slice)

	bitstringInput = bitstringInput.BitSliceXor(t.BitSlice{bitstringInput.QuadJoin(xoredBitSlice), subKeys[rounds]})

	return bitstringInput
}
