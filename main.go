package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file.csv>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(rows) < 2 {
		fmt.Println("Error: No data rows found")
		return
	}

	emails := []string{}
	for _, row := range rows[1:] {
		emails = append(emails, row[1])
	}

	checkdups := findDuplicates(emails)
	if len(checkdups) == 0 {
		fmt.Println("No duplicates found")
		return
	}
	for _, email := range checkdups {
		fmt.Println(email)
	}
}

func findDuplicates(emails []string) []string {
	counts := map[string]int{}
	for _, email := range emails {
		email = strings.ToLower(strings.TrimSpace(email))
		counts[email]++
	}

	dups := []string{}
	for email, n := range counts {
		if n > 1 {
			dups = append(dups, email)
		}
	}
	return dups
}
