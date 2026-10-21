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
		fmt.Println("No duplicates found, and no file is created")
		return
	}
	for _, email := range checkdups {
		fmt.Println(email)
	}

	remdups := removeDuplicates(rows)

	outName := strings.TrimSuffix(os.Args[1], ".csv") + "_cleaned.csv"
	out, err := os.Create(outName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer out.Close()

	writer := csv.NewWriter(out)
	err = writer.WriteAll(remdups)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Saved", outName, ":", "removed", len(rows)-len(remdups), "duplicate row/s")

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

func removeDuplicates(rows [][]string) [][]string {
	seen := map[string]bool{}

	kept := [][]string{rows[0]}

	for _, row := range rows[1:] {
		email := strings.ToLower(strings.TrimSpace(row[1]))
		if seen[email] {
			continue
		}
		seen[email] = true
		kept = append(kept, row)
	}
	return kept
}
