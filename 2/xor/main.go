package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	// xor --cipher || --decipher --secret "secret"
	flag.Parse()

	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided! Exiting now...")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plainText := getUserInput("Enter your text to cipher: ")
		fmt.Println(plainText)

		cipheredText, err := cipherer.Cipher(plainText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("Enter your text to decipher: ")
		fmt.Println(cipheredText)

		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Fprintln(os.Stderr, "Invalid mode. Use 'cipher' or 'decipher'")
	}
}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		res, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading the entered text! Please try again...")
			continue
		}

		return strings.TrimRight(res, "\n")
	}
}
