package functions

import (
	"fmt"
)

// Printing the ASCII art for a word
func AsciiPrint(word string, asciiArtMap map[rune][]string) {
	if word == "" {
		return
	}
	for i := 0; i <= 7; i++ {
		for _, char := range word {
			fmt.Print(asciiArtMap[char][i])
		}
		fmt.Println()
	}
}

// Printing the ASCII art from an array of strings using AsciiPrint
func AsciiPrintArr(textArr []string, asciiArtMap map[rune][]string) {
	for _, text := range textArr {
		if text == "" {
			fmt.Println()
		} else {
			AsciiPrint(text, asciiArtMap)
		}
	}
}

// Converts a word into ASCII art and put it into a string
func AsciiPrintOut(word string, asciiArtMap map[rune][]string) string {
	if word == "" {
		return ""
	}
	var output string
	for i := 0; i <= 7; i++ {
		for _, char := range word {
			output += fmt.Sprint(asciiArtMap[char][i])
		}
		output += fmt.Sprintln()
	}
	return output
}

// Converts an array of words into ASCII art and put it into a string
func AsciiPrintOutArr(textArr []string, asciiArtMap map[rune][]string) string {
	var output string
	for _, text := range textArr {
		if text == "" {
			output += fmt.Sprintln()
		} else {
			output += AsciiPrintOut(text, asciiArtMap)
		}
	}
	return output
}
