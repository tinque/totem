
# Totem

## Project Description

Totem is a Go tool to extract, merge, and deduplicate contacts from SGDF (Scouts et Guides de France) intranet export files and Gmail contacts. It generates a Gmail-compatible CSV file, making it easy to manage and synchronize scout contacts.

**SGDF** stands for "Scouts et Guides de France", a major French scouting organization. The tool is designed to process their intranet export files.

## How to Build

Requirements: Go ≥ 1.25.
Clone the repository and build with:

```sh
go build -o totem main.go
```

## How to Use

Example usage:

```sh
./totem -intranet "20250916 - exportIndividus.xls" -gmail "contacts.csv" -out "output.csv"
```

- `-intranet`: path to the SGDF intranet export file (required)
- `-gmail`: path to the Gmail contacts CSV file (optional)
- `-out`: path to the output CSV file (optional, default: output.csv)

The generated file is compatible with Gmail import.

## How to Contribute

1. Fork the project
2. Create a branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add a feature'`)
4. Push the branch (`git push origin feature/my-feature`)
5. Open a Pull Request

Contributions are welcome!

## License

Beerware License

```
"THE BEER-WARE LICENSE" (Revision 42):
Quentin Georget wrote this file. As long as you retain this notice you can do whatever you want with this stuff. If we meet some day, and you think this stuff is worth it, you can buy me a beer in return.
```
