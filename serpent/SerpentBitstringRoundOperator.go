package serpent

import "fmt"
import t "serpent_aes/types"
import transform "serpent_aes/transformation"

type SerpentBitstringRoundOperator interface {
	ApplyRound(rounds int, inputBitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring
	ApplyInverseRound(rounds int, outputBitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring
}

func (serpent *Serpent) ApplyRound(rounds int, inputBitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring {
	var outputBitstring t.Bitstring
	var statusBoxAppliedBitstring t.Bitstring
	var xoredInputBitstring t.Bitstring

	xoredInputBitstring = xoredInputBitstring.BitSliceXor(t.BitSlice{inputBitstring, subKeys[rounds]})
	statusBoxAppliedBitstring = serpent.ApplyStatusBoxHat(rounds, xoredInputBitstring)

	if 0 <= rounds && rounds <= SERPENT_ROUNDS-2 {
		outputBitstring = transform.LinearTransformation(statusBoxAppliedBitstring)
	} else if rounds == SERPENT_ROUNDS-1 {
		outputBitstring = statusBoxAppliedBitstring.BitSliceXor(t.BitSlice{statusBoxAppliedBitstring, subKeys[SERPENT_ROUNDS]})
	} else {
		fmt.Println("Status box rounds out of range")
	}

	return outputBitstring
}

func (serpent *Serpent) ApplyInverseRound(rounds int, outputBitstring t.Bitstring, subKeys t.BitSlice) t.Bitstring {
	var inputBitstring t.Bitstring
	var xoredBitstring t.Bitstring
	var statusBoxAppliedBitstring t.Bitstring

	if 0 <= rounds && rounds <= SERPENT_ROUNDS-2 {
		statusBoxAppliedBitstring = transform.LinearInverseTransformation(outputBitstring)
	} else if rounds == SERPENT_ROUNDS-1 {
		statusBoxAppliedBitstring = xoredBitstring.BitSliceXor(t.BitSlice{outputBitstring, subKeys[SERPENT_ROUNDS]})
	} else {
		fmt.Println("Status box rounds out of range")
	}

	xoredBitstring = serpent.ApplyInverseStatusBoxHat(rounds, statusBoxAppliedBitstring)
	inputBitstring = xoredBitstring.BitSliceXor(t.BitSlice{xoredBitstring, subKeys[rounds]})

	return inputBitstring
}
