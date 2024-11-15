package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	data := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 255, 255}

	
	hexString := hex.EncodeToString(data)
	fmt.Println("Hexadecimal:", hexString)

	utf8String := string(data)
	fmt.Println("UTF-8 String:", utf8String)


	var intValue uint64
	for _, b := range data {
		intValue = (intValue << 8) | uint64(b)
	}
	fmt.Println("Integer value:", intValue)
}
