package test

import "testing"
import t "serpent_aes/types"
import s "serpent_aes/serpent"

func TestBitstringStatusBoxOperator(test *testing.T) {
	var serpent s.Serpent
	var targetBitstring t.Bitstring = "0111"
	var resultBitstring t.Bitstring = serpent.ApplyStatusBox(2, "0101")

	if targetBitstring != resultBitstring {
		test.Errorf("Output bitstring from status box doesn't match with target bitstring (%s) != (%s)", resultBitstring, targetBitstring)
		test.Fail()
	}
}
