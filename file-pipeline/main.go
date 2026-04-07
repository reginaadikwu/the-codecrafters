// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Adikwu Regina
// Squad: The Channels

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Channels]
// ───────────────────────────────────────────
// Input line types:
//   - Normal report lines
//   - Lines in ALL CAPS
//   - Lines in all lowercase
//   - Lines starting with TODO:
//   - Lines starting with CLASSIFIED:
//   - Lines that are only dashes or blanks
//   - Lines with extra leading/trailing spaces
//   - Lines containing numbers and symbols
//
// Transformation rules (in order):
//   1. [Rule 1: Trim all leading and trailing whitespace]
//   2. [Rule 2: Replace TODO: with ✦ ACTION]
//   3. [Rule 3: Replace CLASSIFIED: with [REDACTED]]
//   4. [Rule 4: Convert ALL CAPS lines to Title Case]
//   5. [Rule 5: Convert all lowercase lines to uppercase   ]
//
// Output format:
//   [Header: yes — "SENTINEL FIELD REPORT — PROCESSED"]
//   [Line numbering format: ]
//   [Summary block: yes]
//   - Total lines processed
//   - Total lines removed
//
// Terminal summary fields:
//  [ ✦ Lines read
//   ✦ Lines written
//   ✦ Lines removed
//   ✦ Rules applied
//  ]
// ═══════════════════════════════════════════

package main

import (
	"fmt"
	"os"
	"strings"
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func replaceTodo(word string) string {
	word = strings.ReplaceAll(word, "TODO:", "✦ ACTION:")
	return word
}

func CLASSIFIED(word string) string {
	return strings.ReplaceAll(word, "CLASSIFIED", "[REDACTED]")
}

func CapsToTitle(word string) string {
	return strings.Title(word)
}

func lowerToUpper(word string) string {
	return strings.ToUpper(word)
}

func main() {

	// It protects the program from crashing.
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	// Take's the file names and stores for usage.
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Reading input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Conversion of raw bytes into a string
	text := string(data)

	// Write output file
	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
