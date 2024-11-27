package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unidoc/unioffice/document"
)

func main() {
	// Input TXT file path
	inputFilePath := "input.txt"
	// Output DOCX file path
	outputFilePath := "output.docx"

	// Open the input text file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a new Word document
	doc := document.New()
	defer doc.Close()

	// Read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Add an empty paragraph for blank lines
			doc.AddParagraph()
		} else {
			// Add each line as a new paragraph in the Word document
			doc.AddParagraph().AddRun().AddText(line)
		}
	}

	// Check for errors during file scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Save the Word document
	err = doc.SaveToFile(outputFilePath)
	if err != nil {
		fmt.Printf("Error saving Word document: %v\n", err)
		return
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFilePath, outputFilePath)
}
