package serpent

import t "serpent_aes/types"

type SerpentParallelStatusBoxHatOperator interface {
	ApplyStatusBoxHat(statusBox int, input t.Bitstring) t.Bitstring
	ApplyInverseStatusBoxHat(statusBox int, output t.Bitstring) t.Bitstring
}

func (serpent *Serpent) ApplyStatusBoxHat(statusBox int, input t.Bitstring) t.Bitstring {
	var bitstring t.Bitstring

	for octectIndex := 0; octectIndex < 32; octectIndex++ {
		bitstring += serpent.ApplyStatusBox(statusBox, input[4*octectIndex:4*(octectIndex+1)])
	}

	return bitstring
}

func (serpent *Serpent) ApplyInverseStatusBoxHat(statusBox int, output t.Bitstring) t.Bitstring {
	var bitstring t.Bitstring

	for octectIndex := 0; octectIndex < 32; octectIndex++ {
		bitstring += serpent.ApplyInverseStatusBox(statusBox, output[4*octectIndex:4*(octectIndex+1)])
	}

	return bitstring
}
