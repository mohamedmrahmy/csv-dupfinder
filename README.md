# csv-dupfinder
This program finds duplicate emails in CSV files, handling messy capitals and spaces
## How to run
To run main.go
```
go run .
```
To run regression tests available
```
go run . testdata\filename.csv
```
To build .exe
```
go build
```
## Limitations
- It only checks one kind of value (emails), and not duplicate rows in general.
- The email column is fixed at index 1
- Expects the first row to be a header, and if the csv file has no header, then the first customer/row is skipped.
- It doesn't show how many duplicates found
- It only finds duplicates, and it doesn't clear a duplicate email
- It doesn’t check whether emails are valid or not (EX: omar @mail.com)

## Regression Testing Outputs
testdata/customers.csv
```
sara@shop.com
```
testdata/empty.csv
```
Error: No data rows found
```
testdata/header_only.csv
```
Error: No data rows found
```
testdata/messy.csv
```
sara@shop.com
omar@mail.com`
```
testdata/no_dups.csv
```
No duplicates found
```
testdata/short_row.csv
```
Error: record on line 3: wrong number of fields
```