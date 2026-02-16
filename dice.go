package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"github.com/fatih/color"
)

var (
	wordlistOpt int
)

func init () {

	// Usage: go run main.go -wordlist=2
	flag.IntVar(&wordlistOpt, "wordlist", 1, "1: Large List, 2: Short List 1, 3: Short List 2")

}

func rollTheDice (choice int) int{

	// Parse User Input and create appropriate random die cast.
	var numDice int

	// Standard EFF Large list uses 5 dice. Short lists use 4.
	if choice == 1 {
		
		numDice = 5
	
	} else {
	
		numDice = 4
	}

	final := 0

	for i := 0; i < numDice; i++ {
		// rand.Int returns 0 to max-1. We need 1-6.
		nBig, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			log.Fatal(err)
		}
		
		// Convert big.Int to standard int and add 1
		n := int(nBig.Int64()) + 1
		
		// Concatenate the digit mathematically (e.g., 1 then 2 becomes 12)
		final = final*10 + n
	}

	return final
	
}

func findNewPassphrase (uChoice int, cast int) string{

	var fName string

	// Read in text files
	switch uChoice {
		case 1:
			fName = "eff_large_wordlist.txt"
		case 2:
			fName = "eff_short_wordlist_1.txt"
		case 3:
			fName = "eff_short_wordlist_2_0.txt"
		default:
		fName = "eff_large_wordlist.txt"

	}

	filePath := fmt.Sprintf("wordlists/%s", fName)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error reading file. Exiting...")
	}
	defer file.Close()

	// Read the file by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			line := scanner.Text()

			// Split by whitespace
			parts := strings.Fields(line)

			// Make sure there are two columns
			if len(parts) >= 2 {
				// Convert first column to integer
				id, err := strconv.Atoi(parts[0])
				if err != nil {
					continue
				}
				
				// If we find a match then return the phrase
				if id == cast {
				
					return parts[1]
				}
			}
	}

	if err := scanner.Err(); err != nil {

		log.Fatal(err)

	}

	return "No match found, try again..."
}

func shannonEntropy(s string) float64 {

	if len(s) == 0 {
		return 0
	}

	frequencies := make(map[rune]int)
	for _, char := range s {
		frequencies[char]++
	}

	var entropy float64
	totalChars := float64(len(s))

	for _, count := range frequencies {
		p := float64(count) / totalChars
		entropy -= p * math.Log2(p)
	}

	return entropy
}

func main() {
	
	flag.Parse()
	fmt.Println("Which wordlist would you like to use: ", wordlistOpt)
	
	var passphraseParts []string

	for i := 0; i < 6; i++ {

		cast := rollTheDice(wordlistOpt)
		word := findNewPassphrase(wordlistOpt, cast)
		
		passphraseParts = append(passphraseParts, word)
	}

	fullPassphrase := strings.Join(passphraseParts, ",")	// Todo give user choice of delimiter

	entropy := shannonEntropy(fullPassphrase)

	// Create Sprint color function
	green := color.New(color.FgHiGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("Your new secret passphrase: %s has an entropy of %s bits.\n", green(fullPassphrase), yellow(entropy)) 

}
