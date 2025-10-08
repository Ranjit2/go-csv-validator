# Go CSV Validator

This is a **CLI tool in Go** to validate CSV files. It supports:

- Dynamic headers (reads columns from CSV)
- Checks for missing required fields
- Validates email format
- Validates 10-digit phone numbers
- Detects duplicate emails
- Detects duplicate UserIDs
- Reports unwanted characters

## Usage

1. **Run the program:**

```bash
cd go-csv-validator
go run main.go
