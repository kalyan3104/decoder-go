package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/bech32"
)

func main() {

	bech32Address := "moa188anxz35atlef7cucszypmvx88lhz4m7a7t7lhcwt6sfphpsqlksr33h66"

	hrp, data, err := bech32.Decode(bech32Address)
	if err != nil {
		log.Fatalf("Failed to decode Bech32 address: %v", err)
	}
	fmt.Println("HRP:", hrp) // Prints the HRP, e.g., "moa"

	decodedBytes, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		log.Fatalf("Failed to convert 5-bit groups to bytes: %v", err)
	}

	hexString := hex.EncodeToString(decodedBytes)
	fmt.Println("Hexadecimal:", hexString)
}
