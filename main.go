package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	// Define flags
	help := flag.Bool("help", false, "Show help message")
	keySize := flag.Int("keysize", 32, "Size of the key in bytes (default: 32 for 256-bit key)")
	ivSize := flag.Int("ivsize", 16, "Size of the IV in bytes (default: 16 for 128-bit IV)")
	
	// Parse the flags
	flag.Parse()

	// Show help message if --help flag is set
	if *help {
		showHelp()
		return
	}

	// Generate random key
	key, err := generateRandomBytes(*keySize)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	
	// Generate random IV
	iv, err := generateRandomBytes(*ivSize)
	if err != nil {
		fmt.Println("Error generating IV:", err)
		return
	}

	// Print results
	fmt.Println("Base64 Encoded Key:", base64.StdEncoding.EncodeToString(key))
	fmt.Println("Base64 Encoded IV:", base64.StdEncoding.EncodeToString(iv))
}

// showHelp prints the help message for the program
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