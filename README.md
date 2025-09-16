
# Totem

## Project Description

Totem is a Go tool to extract, merge, and deduplicate contacts from SGDF (Scouts et Guides de France) intranet export files and Gmail contacts. It generates a Gmail-compatible CSV file, making it easy to manage and synchronize scout contacts.

**SGDF** stands for "Scouts et Guides de France", a major French scouting organization. The tool is designed to process their intranet export files.


## How to Build

Requirements: Go ≥ 1.25.

### Build for your platform (manual)
Clone the repository and build with:
```sh
go build -o totem main.go
```

### Build for all platforms (recommended)
To build binaries for macOS, Linux, and Windows (both amd64 and arm64), use the provided script:

```sh
bash build_all.sh
```

This will create binaries in the `build/` directory, named like:
```
build/totem-v1.0-darwin-amd64
build/totem-v1.0-linux-amd64
build/totem-v1.0-windows-amd64.exe
...etc
```
The version is automatically detected from the latest git tag (e.g., `v1.0`). If no tag is present, the version will be `dev`.

The `build/` directory is ignored by git and safe for binary outputs.

## How to Use

Example usage:

```sh
./totem -intranet "20250916 - exportIndividus.xls" -gmail "contacts.csv" -out "output.csv"
```

- `-intranet`: path to the SGDF intranet export file (required)
- `-gmail`: path to the Gmail contacts CSV file (optional)
- `-out`: path to the output CSV file (optional, default: output.csv)


## Download & Use Pre-built Binaries

You can download ready-to-use binaries from the [GitHub Releases page](https://github.com/tinque/totem/releases).

1. Go to the release page and download the binary matching your platform:
	- macOS (Intel): `totem-v1.0-darwin-amd64`
	- macOS (Apple Silicon): `totem-v1.0-darwin-arm64`
	- Linux (Intel): `totem-v1.0-linux-amd64`
	- Linux (ARM): `totem-v1.0-linux-arm64`
	- Windows: `totem-v1.0-windows-amd64.exe` or `totem-v1.0-windows-arm64.exe`
2. Place the binary in your desired folder and make it executable (if needed):
	- On macOS/Linux: `chmod +x totem-v1.0-*`
3. Run the binary as shown below:
	```sh
	./totem-v1.0-darwin-amd64 -intranet "20250916 - exportIndividus.xls" -gmail "contacts.csv" -out "output.csv"
	```

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
