package test

import "testing"
import t "serpent_aes/types"

func TestRotateLeft(test *testing.T) {
  var bitstring t.Bitstring = "00100100011001010110001101110010"+
                              "01100101011101000111010001100101"+
                              "00100100011101000110010001100001"+
                              "01110100011010110110010101111001";

  var targetBitstring t.Bitstring = "10101111001001001000110010101100"+
                                    "01101110010011001010111010001110"+
                                    "10001100101001001000111010001100"+
                                    "10001100001011101000110101101100"
  var testedBitstring = bitstring.RotateLeft(11)


  if testedBitstring != targetBitstring {
    test.Errorf("Bitstring tested (%s) with left rotate operator to 11th position doesn't match with target bitstring (%s)", testedBitstring, bitstring)
    test.Fail()
  }

}
