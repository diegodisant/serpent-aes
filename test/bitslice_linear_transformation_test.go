package test

import "testing"
import t "serpent_aes/types"
import transform "serpent_aes/transformation"

func TestLinearTransformationBitSlice(test *testing.T) {
	var bitstring t.Bitstring = "11001110010100110010100101010010" +
		"10111000100110010100101001010111" +
		"00101001010101100100100101011010" +
		"10010100101010010101001010100101"

	var bitslice t.BitSlice = t.BitSlice{
		"00000001010101011010100101000110",
		"11011011100100101111100011110010",
		"10110001001110011100011111111100",
		"10111011101100011010001010101011",
	}

	var bitsliceQuadSplited t.BitSlice = bitstring.QuadSplit()
	var bitsliceLinearTransformed t.BitSlice = transform.LinearTransformationBitSlice(bitsliceQuadSplited)

	for sliceIndex := 0; sliceIndex < 4; sliceIndex++ {
		if bitsliceQuadSplited[sliceIndex] != bitslice[sliceIndex] {
			test.Errorf("Bit slice (%d) with 32-bit segment (%s) is not equals than (%s)", sliceIndex, bitslice[sliceIndex], bitsliceQuadSplited[sliceIndex])
			test.Fail()
		}
	}

	var bitsliceLinearInverseTransformed t.BitSlice = transform.LinearInverseTransformationBitSlice(bitsliceLinearTransformed)

	for sliceIndex := 0; sliceIndex < 4; sliceIndex++ {
		if bitsliceLinearInverseTransformed[sliceIndex] != bitsliceQuadSplited[sliceIndex] {
			test.Errorf("Linear inverse transformation of bitslice [%d] (%s) is not equals than (%s)", sliceIndex, bitsliceLinearInverseTransformed[sliceIndex], bitsliceQuadSplited[sliceIndex])
			test.Fail()
		}
	}
}
