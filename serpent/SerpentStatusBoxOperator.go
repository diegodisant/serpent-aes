package serpent

import t "serpent_aes/types"

type SerpentStatusBoxOperator interface {
  ApplyStatusBox(statusBox int, input t.Bitstring) t.Bitstring
  ApplyInverseStatusBox(statusBox int, output t.Bitstring) t.Bitstring
}

func (serpent *Serpent) ApplyStatusBox(statusBox int, input t.Bitstring) t.Bitstring {
  return BitstringStatusBox[statusBox % 8][input]
}

func (serpent *Serpent) ApplyInverseStatusBox(statusBox int, output t.Bitstring) t.Bitstring {
  return BitstringStatusBoxInverse[statusBox % 8][output]
}
