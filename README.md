# csv-dupfinder

Finds duplicate email addresses in a CSV file, even when the same email is typed with different capital letters or extra spaces.

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
```

`Sara@Shop.com` was matched with `sara@shop.com` (different capital letters), and ` omar@mail.com` was matched with `omar@mail.com` (an extra space at the start). The tool ignores both kinds of difference before comparing.

## How it handles problems

| Situation | Output |
|---|---|
| Duplicates found | Each duplicate email, one per line |
| No duplicates | `No duplicates found` |
| No file name given | `Usage: go run . <file.csv>` |
| File doesn't exist | `Error: open <name>: The system cannot find the file specified.` |
| Empty file | `Error: No data rows found` |
| File has only a header row | `Error: No data rows found` |
| A row with missing columns | `Error: record on line <n>: wrong number of fields` |

## Limitations

- It only checks emails, not whole rows for duplicates.
- Emails must be in the second column of the file.
- The first row must be a header. If a file has no header, the first row of data is skipped.
- It doesn't show how many times each email appears, or on which rows.
- It only reports duplicates. It doesn't remove them or create a cleaned file.
- It doesn't check whether emails are valid (for example, `omar @mail.com`).