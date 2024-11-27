package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"time"
)

type AccountInfo struct {
	Address     string `json:"Address"`
	OwnerAddress string `json:"Owner Address"`
}

func readAddressesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var addresses []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addresses = append(addresses, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return addresses, nil
}

func checkAccountOwner(addr string, retries int, delay time.Duration) (AccountInfo, error) {
	url := fmt.Sprintf("https://api.elrond.com/accounts/%s", addr)
	client := &http.Client{}

	for attempt := 0; attempt < retries; attempt++ {
		resp, err := client.Get(url)
		if err != nil {
			return AccountInfo{Address: addr, OwnerAddress: "N/A"}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			var data map[string]interface{}
			body, _ := ioutil.ReadAll(resp.Body)
			err := json.Unmarshal(body, &data)
			if err != nil {
				return AccountInfo{Address: addr, OwnerAddress: "N/A"}, err
			}

			owner, exists := data["ownerAddress"].(string)
			if !exists {
				owner = "N/A"
			}
			return AccountInfo{Address: addr, OwnerAddress: owner}, nil
		} else if resp.StatusCode == 429 {
			// Rate limit error, retry after delay
			fmt.Printf("Rate limited for address %s, retrying in %v seconds...\n", addr, delay.Seconds())
			time.Sleep(delay)
			delay *= 2 // Exponential backoff
		} else {
			return AccountInfo{Address: addr, OwnerAddress: "N/A"}, fmt.Errorf("failed with status %d", resp.StatusCode)
		}
	}

	// After retries, return failure
	return AccountInfo{Address: addr, OwnerAddress: "N/A"}, fmt.Errorf("failed after %d attempts", retries)
}

func processAddresses(addresses []string) ([]AccountInfo, error) {
	var results []AccountInfo
	retries := 3
	delay := 10 * time.Second

	// Use a channel to gather results
	resultCh := make(chan AccountInfo)
	errCh := make(chan error, len(addresses))

	// Process each address concurrently
	for _, addr := range addresses {
		go func(addr string) {
			result, err := checkAccountOwner(addr, retries, delay)
			if err != nil {
				errCh <- err
			} else {
				resultCh <- result
			}
		}(addr)
	}

	// Wait for all results
	for i := 0; i < len(addresses); i++ {
		select {
		case result := <-resultCh:
			results = append(results, result)
		case err := <-errCh:
			log.Printf("Error processing address: %v", err)
		}
	}

	return results, nil
}

func saveResultsToFile(results []AccountInfo) error {
	file, err := os.Create("addresses_with_owners.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	err = writer.Write([]string{"Address", "Owner Address"})
	if err != nil {
		return err
	}

	// Write each result
	for _, result := range results {
		err := writer.Write([]string{result.Address, result.OwnerAddress})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	addresses, err := readAddressesFromFile("addresses.txt")
	if err != nil {
		log.Fatalf("Error reading addresses: %v", err)
	}

	// Process the addresses concurrently
	results, err := processAddresses(addresses)
	if err != nil {
		log.Fatalf("Error processing addresses: %v", err)
	}

	// Save the results to a CSV file
	err = saveResultsToFile(results)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	fmt.Println("CSV file has been created successfully!")
}
