package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func readTxtFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func createExcelFile(data []string, excelFilePath string) error {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet or use the default sheet
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName) // Corrected: Handling both return values
	if err != nil {
		return err
	}

	// Write headers
	f.SetCellValue(sheetName, "A1", "Address")
	f.SetCellValue(sheetName, "B1", "Owner Address")

	// Loop through the data and write each line to a new row
	for rowIndex, line := range data {
		// In the text file, it is expected that each line is "address,owner_address"
		columns := strings.Split(line, ",")
		if len(columns) == 2 {
			// Write the address to column A (rowIndex + 2 for header row)
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex+2), columns[0])
			// Write the owner address to column B
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex+2), columns[1])
		}
	}

	// Set the active sheet of the workbook
	f.SetActiveSheet(index)

	// Save the file to the specified path
	err = f.SaveAs(excelFilePath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Specify the path to your .txt file and the output Excel file
	txtFilePath := "data.txt"   // Replace with your actual .txt file
	excelFilePath := "output.xlsx" // Output Excel file path

	// Read data from .txt file
	data, err := readTxtFile(txtFilePath)
	if err != nil {
		fmt.Println("Error reading .txt file:", err)
		return
	}

	// Create Excel file and write the data
	err = createExcelFile(data, excelFilePath)
	if err != nil {
		fmt.Println("Error creating Excel file:", err)
		return
	}

	fmt.Println("Excel file created successfully:", excelFilePath)
}
