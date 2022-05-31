package main

import (
	"BlockChainTasks/secondTask/convert"
	"math/big"
)

var Keys [][]byte

func main() {
	var g = 255
	var h = big.NewInt(3456790467467)
	a := "ff00000000000000000000000000000000000000000000000000000000000000"
	convert.HexToLittleEndian(a)
	convert.LittleEndianToHex(g)
	convert.HexToBigEndian(a)
	convert.BigEndianToHex(h)

}
