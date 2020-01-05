package test

import "fmt"
import "testing"
import t "serpent_aes/types"
import blk "serpent_aes/block"

func TestSegmentBlock128Xor(test *testing.T) {
  var targetBitstring t.Bitstring
  var target128Block blk.Block128
  var testBlock1 blk.Block128
  var testBlock2 blk.Block128

  targetBitstring = "0101"
  target128Block = target128Block.FromBitstringToBlock128(targetBitstring)
  testBlock1 = target128Block.FromBitstringToBlock128(t.Bitstring("0100"))
  testBlock2 = target128Block.FromBitstringToBlock128(t.Bitstring("0001"))

  fmt.Println(target128Block)
  testBlock1.Block128SizeXor(testBlock2)

  if testBlock1 != target128Block {
    test.Errorf("Xor operation between two bitstrings (%s) XOR (%s) with expected result (%s) doesn't match with (%s)", "0100", "0001", targetBitstring, testBlock1)
    test.Fail()
  }
}
