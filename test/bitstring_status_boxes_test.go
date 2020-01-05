package test

import "testing"
import s "serpent_aes/serpent"
import t "serpent_aes/types"

func TestStatusBoxesBitstring(test *testing.T) {

	var normalStatusBitString t.Bitstring = s.BitstringStatusBox[4]["1110"]

	if normalStatusBitString != "0110" {
		test.Errorf("Normal status box bitstring result wasn't expected (%s) != (0110)", normalStatusBitString)
		test.Fail()
	}

	var inverseStatusBitstring t.Bitstring = s.BitstringStatusBoxInverse[4]["0110"]

	if inverseStatusBitstring != "1110" {
		test.Errorf("Inverse status box bitstring result wasn't expected (%s) != (1110)", inverseStatusBitstring)
		test.Fail()
	}
}
