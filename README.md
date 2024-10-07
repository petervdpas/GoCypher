
# GoCypher - A Random Key and IV Generator

This Go program generates a random cryptographic key and initialization vector (IV) that can be used in encryption algorithms. The key and IV are output in Base64-encoded format. The sizes of the key and IV are configurable via command-line flags.

## Features

- Generates a random cryptographic key.
- Generates a random initialization vector (IV).
- Outputs Base64-encoded values for easy integration into various systems.
- Customizable key and IV sizes.

## Requirements

This project uses Go modules. Make sure you are using Go 1.22 or later, as defined in the `go.mod` file.

The `go.mod` file for this project looks like:

```go
module gocypher

go 1.22.0
```

## Installation

First, clone the repository to your local machine:

```bash
git clone https://github.com/petervdpas/gocypher.git
cd gocypher
```

Next, make sure you have Go 1.22 or later installed. Run the following command to download dependencies (if any):

```bash
go mod tidy
```

### Compiling the Program

To compile the program into an executable (`.exe`), use the following command:

```bash
go build -ldflags="-s -w" -o randombyte.exe main.go
```

This command minimizes the size of the executable by stripping symbol tables and debug information.

### Running the Program

You can run the executable with the following command:

```bash
./myprogram.exe
```

By default, the program will generate a 256-bit key (32 bytes) and a 128-bit IV (16 bytes).

### Command-Line Options

You can use the following flags to customize the size of the key and IV:

- `--keysize`: Size of the key in bytes (default is 32 for a 256-bit key).
- `--ivsize`: Size of the IV in bytes (default is 16 for a 128-bit IV).
- `--help`: Displays help information.

#### Example Usage

```bash
./myprogram.exe --keysize 24 --ivsize 12
```

This example will generate a 192-bit key (24 bytes) and a 96-bit IV (12 bytes).

### Output Example

```bash
Base64 Encoded Key: YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYQ==
Base64 Encoded IV: YmJiYmJiYmJiYmJiYmJiYmJi
```

### Help

To display the help message:

```bash
./myprogram.exe --help
```

This will display:

```bash
This program generates random keys and IVs for encryption.
You can specify the size of the key and IV using the following flags:

Options:
  --help        Show this help message
  --keysize     Size of the key in bytes (default: 32 for 256-bit key)
  --ivsize      Size of the IV in bytes (default: 16 for 128-bit IV)

Example usage:
  go run main.go --keysize 32 --ivsize 16
```

## Security Considerations

This program uses the Go `crypto/rand` package to securely generate random bytes for both the key and IV. If the system's random number generator fails, an error is returned, and the program will not proceed.

## License

This project is licensed under the MIT License.
