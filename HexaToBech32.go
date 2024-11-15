package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/bech32"
)
                 //2333000000000000000000000000000000023330000000000000000000000000
func main() {
	hexString := "23330000000000023330000000000000000000000000000000000000a9fFFFff"

	data, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	
	convertedData, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		log.Fatalf("Failed to convert to 5-bit groups: %v", err)
	}


	bech32Address, err := bech32.Encode("moa", convertedData)
	if err != nil {
		log.Fatalf("Failed to encode Bech32 address: %v", err)
	}

	fmt.Println("Bech32 Address with HRP 'moa':", bech32Address)
}
