package convert

import (
	_ "bytes"
	"encoding/binary"
	_ "encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
)

func HexToLittleEndian(s string) {
	data, _ := hex.DecodeString(s)
	v := binary.LittleEndian.Uint32(data)
	fmt.Println(v)
}

func HexToBigEndian(s string) {
	data, _ := hex.DecodeString(s)
	v := binary.BigEndian.Uint32(data)
	fmt.Println(v)
}

func LittleEndianToHex(n int) {
	var v []byte
	c := byte(n)
	v = append(v, c)
	data := hex.EncodeToString(v)
	fmt.Println(data)
}

func BigEndianToHex(n *big.Int) {
	f := n.Bytes()
	data := hex.EncodeToString(f)
	fmt.Println(data)
}
