package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

func main() {

	hexString := `
2333000000000002333000000000000000000000000000000000000000000031
2333000000000002333000000000000000000000000000000000000000000032
2333000000000002333000000000000000000000000000000000000000000033
2333000000000002333000000000000000000000000000000000000000000012
2333000000000002333000000000000000000000000000000000000000000013
2333000000000002333000000000000000000000000000000000000000000011
2333000000000002333000000000000000000000000000000000000000000000
2333000000000002333000000000000000000000000000000000000000000021
`


	hexString = strings.ReplaceAll(hexString, "\n", "")
	hexString = strings.ReplaceAll(hexString, " ", "")

	
	const chunkSize = 64
	var chunks []string

	for i := 0; i < len(hexString); i += chunkSize {
		end := i + chunkSize
		if end > len(hexString) {
			end = len(hexString)
		}
		chunks = append(chunks, hexString[i:end])
	}


	for i, chunk := range chunks {
		//fmt.Printf("Processing Chunk #%d: %s\n", i+1, chunk)

		
		data, err := hex.DecodeString(chunk)
		if err != nil {
			log.Printf("Failed to decode chunk #%d: %v\n", i+1, err)
			continue
		}

	
		convertedData, err := bech32.ConvertBits(data, 8, 5, true)
		if err != nil {
			log.Printf("Failed to convert to 5-bit groups for chunk #%d: %v\n", i+1, err)
			continue
		}

		
		bech32Address, err := bech32.Encode("moa", convertedData)
		if err != nil {
			log.Printf("Failed to encode Bech32 address for chunk #%d: %v\n", i+1, err)
			continue
		}


		fmt.Println(bech32Address)
	}
}
