package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Initialize the program with flags
	keySize, ivSize := parseFlags()

	// Generate random key and IV
	key, err := generateRandomBytes(keySize)
	if err != nil {
		log.Fatalf("Error generating key: %v", err)
	}

	iv, err := generateRandomBytes(ivSize)
	if err != nil {
		log.Fatalf("Error generating IV: %v", err)
	}

	// Print the Base64 encoded results
	fmt.Println("Base64 Encoded Key:", base64.StdEncoding.EncodeToString(key))
	fmt.Println("Base64 Encoded IV:", base64.StdEncoding.EncodeToString(iv))
}

// parseFlags parses the command-line flags and returns the keySize and ivSize
func parseFlags() (int, int) {
	help := flag.Bool("help", false, "Show help message")
	keySize := flag.Int("keysize", 32, "Size of the key in bytes (default: 32 for 256-bit key)")
	ivSize := flag.Int("ivsize", 16, "Size of the IV in bytes (default: 16 for 128-bit IV)")
	flag.Usage = showHelp // Use custom help message

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Validate key and IV sizes
	if *keySize <= 0 {
		log.Fatalf("Invalid key size: %d. Key size must be greater than 0.", *keySize)
	}
	if *ivSize <= 0 {
		log.Fatalf("Invalid IV size: %d. IV size must be greater than 0.", *ivSize)
	}

	return *keySize, *ivSize
}

// showHelp prints the help message
func showHelp() {
	fmt.Println("This program generates random keys and IVs for encryption.")
	fmt.Println("You can specify the size of the key and IV using the following flags:")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help        Show this help message")
	fmt.Println("  --keysize     Size of the key in bytes (default: 32 for 256-bit key)")
	fmt.Println("  --ivsize      Size of the IV in bytes (default: 16 for 128-bit IV)")
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("  go run main.go --keysize 32 --ivsize 16")
}

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random number generator fails to function correctly.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
