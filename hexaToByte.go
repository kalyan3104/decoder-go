package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	
	hexString := "000000000100000000000000000"


	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	fmt.Println("Byte array:", byteArray)
}
