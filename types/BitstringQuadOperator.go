package types

import "fmt"

type BitstringQuadOperator interface {
	QuadSplit() BitSlice
	QuadJoin(bitslice BitSlice) Bitstring
}

func (bitstring Bitstring) QuadSplit() BitSlice {
	if len(bitstring) != 128 {
		fmt.Println("To apply QuadSplit operation to one bitstring needs have a 128-bitslice")
	}

	var outputSlice = make(BitSlice, 4)

	for _32BitHop := 0; _32BitHop < 4; _32BitHop++ {
		outputSlice[_32BitHop] = bitstring[(_32BitHop * 32):((_32BitHop + 1) * 32)]
	}

	return outputSlice
}

func (bitstring *Bitstring) QuadJoin(bitslice BitSlice) Bitstring {
	if len(bitslice) != 4 {
		fmt.Println("To apply QuadSplit operation to on bit slice needs to have the size of 4")
	}

	return bitslice[0] + bitslice[1] + bitslice[2] + bitslice[3]
}
