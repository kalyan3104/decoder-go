package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	
	hexString := "233300000000000233300000000000000000000000000000000000000001ffff"


	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	fmt.Println("Byte array:", byteArray)
}
