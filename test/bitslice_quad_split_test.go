package test

import "testing"
import t "serpent_aes/types"

func TestQuadSplit(test *testing.T) {
  var bitstring t.Bitstring = "00100100011001010110001101110010"+
                              "01100101011101000111010001100101"+
                              "00100100011101000110010001100001"+
                              "01110100011010110110010101111001";
  var slice t.BitSlice = bitstring.QuadSplit()
  var sliceLength int = len(slice)
  var testedLength int
  var expectedBitstring t.Bitstring

  if sliceLength == 4 {
    for sliceIndex := 0; sliceIndex < sliceLength; sliceIndex++ {
      testedLength = len(slice[sliceIndex])
      expectedBitstring = bitstring[sliceIndex*32 : (sliceIndex+1)*32]

      if slice[sliceIndex] != expectedBitstring {
        test.Errorf("Slices are different s0 (%s) != s1 (%s)", slice[sliceIndex], expectedBitstring)
        test.Fail()
      }

      if testedLength != 32 {
        test.Errorf("Slice length (%d) is not 32-bit size", testedLength)
        test.Fail()
      }
    }
  }
}
