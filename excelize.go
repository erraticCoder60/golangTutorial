package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Create a new Excel file
	file := excelize.NewFile()

	// Create a new sheet
	sheetName := "Sheet1"
	a, _ := file.NewSheet(sheetName)
	fmt.Println(a)

	// Set value of a cell
	file.SetCellValue(sheetName, "A1", "Hello")
	file.SetCellValue(sheetName, "B1", "World!")
	file.SetCellValue(sheetName, "A2", "Hello")
	file.SetCellValue(sheetName, "B2", "Erratic Coder!")

	// Set the column headers
	for col := 1; col <= 100; col++ {
		// columnName := string('A' - 1 + col)
		cell := ToAlphaString(col) + "1"
		file.SetCellValue(sheetName, cell, "Column "+strconv.Itoa(col))
	}

	// Set the data for a million rows
	for row := 2; row <= 1000; row++ {
		for col := 1; col <= 100; col++ {
			columnName := ToAlphaString(col)
			cell := columnName + strconv.Itoa(row)
			file.SetCellValue(sheetName, cell, "Data "+strconv.Itoa(row)+"-"+strconv.Itoa(col))
		}
	}

	// Get all the rows in a specific sheet
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error getting rows from Excel file:", err)
		return
	}

	// Iterate over each row and print the data
	for _, row := range rows {
		for _, cell := range row {
			fmt.Printf("%s\t", cell)
		}
		fmt.Println()
	}

	// Save the Excel file
	if err := file.SaveAs("example.xlsx"); err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	fmt.Println("Excel file created successfully")
}

// ToAlphaString converts column number to Excel-style column letter
func ToAlphaString(column int) string {
	result := ""
	for {
		if column <= 0 {
			break
		}
		column--
		result = string('A'+(column%26)) + result
		column /= 26
	}
	return result
}
