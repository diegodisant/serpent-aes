package test

import "testing"
import t "serpent_aes/types"

func TestQuadJoin(test *testing.T) {
  var bitstring t.Bitstring = "00100100011001010110001101110010"+
                              "01100101011101000111010001100101"+
                              "00100100011101000110010001100001"+
                              "01110100011010110110010101111001";
  var slice t.BitSlice = bitstring.QuadSplit()
  var bitstringJoined t.Bitstring = bitstring.QuadJoin(slice)

  if bitstringJoined != bitstring {
    test.Errorf("Bitstrings are not equals after apply quad join operation")
    test.Fail()
  }
}
