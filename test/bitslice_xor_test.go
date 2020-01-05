package test

import "testing"
import t "serpent_aes/types"

func TestBitSliceXor(test *testing.T) {
  var targetBitstring t.Bitstring
  var outputBitstring t.Bitstring

  targetBitstring = "0000"
  outputBitstring = targetBitstring.BitSliceXor(t.BitSlice{"0000", "0000"})

  if outputBitstring != targetBitstring {
    test.Errorf("Xor of two 0 (0000) bit strings are not equal (%s) != (%s)", targetBitstring, outputBitstring)
    test.Fail()
  }

  targetBitstring = "1001"
  outputBitstring = targetBitstring.BitSliceXor(t.BitSlice{"1000", "0001"})

  if outputBitstring != targetBitstring {
    test.Errorf("Xor of two bitstrings (%s) XOR (%s) with result (%s) is not (%s)", "1000", "0001", targetBitstring, outputBitstring)
    test.Fail()
  }

  targetBitstring = "00100"
  outputBitstring = targetBitstring.BitSliceXor(t.BitSlice{"00001", "00100", "00001"})

  if outputBitstring != targetBitstring {
    test.Errorf("Xor of multiple bitstrings (%s) XOR (%s) XOR (%s) with expected result (%s) doesn't match with (%s)", "00001", "00100", "00001", outputBitstring, targetBitstring)
    test.Fail()
  }
}
