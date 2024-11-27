package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v52/github"
	"github.com/xuri/excelize/v2"
	"golang.org/x/oauth2"
)

func main() {
	// GitHub configuration
	organization := "DharitriOne" // Replace with the organization name
	token := "your_personal_access_token" // Replace with your GitHub token

	// Authenticate with GitHub
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Fetch all repositories for the organization
	repos, err := fetchRepositories(ctx, client, organization)
	if err != nil {
		log.Fatalf("Error fetching repositories: %v", err)
	}

	// Create a new Excel file
	excelFile := excelize.NewFile()
	sheetName := "Repositories"
	excelFile.NewSheet(sheetName)
	excelFile.DeleteSheet("Sheet1")

	// Write the header row
	headers := []string{"S.No", "Repository Name", "No of Test Cases", "No of Cases Passed", "No of Cases Failed", "Error Messages", "Reasons For Failures", "Solution", "Current Status", "Notes"}
	for colIndex, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		excelFile.SetCellValue(sheetName, cell, header)
	}

	// Populate the Excel sheet with repository data
	for i, repo := range repos {
		rowIndex := i + 2
		writeExcelRow(excelFile, sheetName, rowIndex, i+1, repo)
	}

	// Save the Excel file
	outputFileName := "repositories_report.xlsx"
	if err := excelFile.SaveAs(outputFileName); err != nil {
		log.Fatalf("Error saving Excel file: %v", err)
	}

	fmt.Printf("Excel report successfully created: %s\n", outputFileName)
}

// fetchRepositories fetches all repositories for the given organization
func fetchRepositories(ctx context.Context, client *github.Client, org string) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opts := &github.RepositoryListByOrgOptions{Type: "all", ListOptions: github.ListOptions{PerPage: 50}}

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opts)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return allRepos, nil
}

// writeExcelRow writes a row of data to the Excel file
func writeExcelRow(excelFile *excelize.File, sheetName string, rowIndex, serialNo int, repo *github.Repository) {
	data := []interface{}{
		serialNo,
		repo.GetName(),
		"Pending", // Placeholder for "No of Test Cases"
		"Pending", // Placeholder for "No of Cases Passed"
		"Pending", // Placeholder for "No of Cases Failed"
		"Pending", // Placeholder for "Error Messages"
		"Pending", // Placeholder for "Reasons For Failures"
		"Pending", // Placeholder for "Solution"
		"Pending", // Placeholder for "Current Status"
		"Pending", // Placeholder for "Notes"
	}

	for colIndex, value := range data {
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex)
		excelFile.SetCellValue(sheetName, cell, value)
	}
}
