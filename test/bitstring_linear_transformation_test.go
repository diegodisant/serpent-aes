package test

import "testing"
import t "serpent_aes/types"
import transform "serpent_aes/transformation"

func TestLinearTransformation(test *testing.T){
  var bitstring t.Bitstring = "00100100011001010110001101110010"+
                              "01100101011101000111010001100101"+
                              "00100100011101000110010001100001"+
                              "01110100011010110110010101111001"

  var targetBitstring t.Bitstring = "11010010101001101000011001100110"+
                                    "11010100101010101000101110110101"+
                                    "00001010100011110001011000110110"+
                                    "01011111100111101101010000001011"

  var testedBitstring = transform.LinearTransformation(bitstring)

  if testedBitstring != targetBitstring {
    test.Errorf("failed to apply linear transformation on bitstring (%s) doesn't match with (%s)", bitstring, testedBitstring)
  }
}
