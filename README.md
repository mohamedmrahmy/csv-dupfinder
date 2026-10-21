# csv-dupfinder

Finds duplicate email addresses in a CSV file, even when the same email is typed with different capital letters or extra spaces, it removes those duplicates and also saves a cleaned copy.

## How to run

Run the program on your CSV file:

```
.\csv-dupfinder.exe yourfile.csv
```

### For developers (requires Go)

Run from source:
```
go run . yourfile.csv
```

Run on one of the included test files:

```
go run . testdata\messy.csv
```

Build the `.exe`:

```
go build
```
## How duplicates are removed
When the same email appears more than once, the cleaned file keeps only the first row with that email and removes the later ones. Emails are compared without regard to capital letters or extra spaces, but the rows that are kept are copied exactly as they appear in your file, unchanged and in their original order. Your original file is never modified; the cleaned copy is saved as a new file ending in _cleaned.csv.

## Example

Input file (`testdata\messy.csv`):

```
name,email,city
Sara,sara@shop.com,Cairo
Sara,Sara@Shop.com,Cairo
Omar,omar@mail.com,Giza
Omar, omar@mail.com,Giza
```

Command:

```
.\csv-dupfinder.exe testdata\messy.csv
```

Output:

```
sara@shop.com
omar@mail.com
Saved testdata\messy_cleaned.csv : removed 2 duplicate row/s
```

`Sara@Shop.com` was matched with `sara@shop.com` (different capital letters), and ` omar@mail.com` was matched with `omar@mail.com` (an extra space at the start). The tool ignores both kinds of difference before comparing.

## How it handles problems

| Situation | Output |
|---|---|
| Duplicates found | Each duplicate email, one per line, Saved ..._cleaned.csv : removed N duplicated row/s |
| No duplicates | `No duplicates found, and no file is created` |
| No file name given | `Usage: go run . <file.csv>` |
| File doesn't exist | `Error: open <name>: The system cannot find the file specified.` |
| Empty file | `Error: No data rows found` |
| File has only a header row | `Error: No data rows found` |
| A row with missing columns | `Error: record on line <n>: wrong number of fields` |

## Limitations

- It only checks emails duplicates, not whole rows for duplicates.
- Emails must be in the second column of the file.
- The first row must be a header. If a file has no header, the first row of data is skipped.
- When a file has NO duplicates, NO cleaned file is created.
- It doesn't show how many times each email appears, or on which rows.
- It doesn't check whether emails are valid (for example, `omar @mail.com`).