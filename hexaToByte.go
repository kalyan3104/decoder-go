package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	
	hexString := "00000000000000000f0f1234567890abcdef"


	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	fmt.Println("Byte array:", byteArray)
}
