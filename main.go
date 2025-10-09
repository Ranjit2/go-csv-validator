package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

// Regex patterns
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var phoneRegex = regexp.MustCompile(`^\d{10}$`)

func main() {
	var filePath string
	fmt.Print("Enter CSV file path: ")
	fmt.Scanln(&filePath)

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read CSV
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(rows) < 2 {
		fmt.Println("CSV has no data")
		return
	}

	header := rows[0]
	emailIndex := indexOf(header, "Email")
	userIDIndex := indexOf(header, "UserID")
	phoneIndex := indexOf(header, "Phone")

	emailMap := make(map[string][]int)
	userIDMap := make(map[string][]int)
	errorsFound := false
	errorMessages := []string{}

	for i, row := range rows[1:] {
		rowNum := i + 2
		rowErrors := []string{}

		// Check missing fields
		for j, val := range row {
			if val == "" {
				rowErrors = append(rowErrors, fmt.Sprintf("%s missing", header[j]))
			}
		}

		// Validate email
		if emailIndex != -1 && emailIndex < len(row) && row[emailIndex] != "" && !emailRegex.MatchString(row[emailIndex]) {
			rowErrors = append(rowErrors, "Invalid email format")
		}

		// Validate phone
		if phoneIndex != -1 && phoneIndex < len(row) && row[phoneIndex] != "" && !phoneRegex.MatchString(row[phoneIndex]) {
			rowErrors = append(rowErrors, "Invalid phone number")
		}

		// Track duplicates
		if emailIndex != -1 && emailIndex < len(row) {
			emailMap[row[emailIndex]] = append(emailMap[row[emailIndex]], rowNum)
		}
		if userIDIndex != -1 && userIDIndex < len(row) {
			userIDMap[row[userIDIndex]] = append(userIDMap[row[userIDIndex]], rowNum)
		}

		if len(rowErrors) > 0 {
			errorsFound = true
			errorMessages = append(errorMessages, fmt.Sprintf("Row %d: %v", rowNum, rowErrors))
		}
	}

	// Check duplicate emails
	for email, rows := range emailMap {
		if len(rows) > 1 {
			errorsFound = true
			errorMessages = append(errorMessages, fmt.Sprintf("Duplicate email '%s' in rows %v", email, rows))
		}
	}

	// Check duplicate UserIDs
	for uid, rows := range userIDMap {
		if len(rows) > 1 {
			errorsFound = true
			errorMessages = append(errorMessages, fmt.Sprintf("Duplicate UserID '%s' in rows %v", uid, rows))
		}
	}

	if errorsFound {
		fmt.Println("\nSuccessfully run: There are errors")
		for _, msg := range errorMessages {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("\nSuccessfully validated: Looks good!")
	}
}

func indexOf(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
