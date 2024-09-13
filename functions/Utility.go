package functions

import (
	"bufio"
	"fmt"
	"os"
)

// AssignArt reads a file containing ASCII art and assigns each character's art to a map.
func AssignArt(banner string) map[rune][]string {
	// Opening the file specified by the 'banner' argument.
	file, err := os.Open(banner)
	if err != nil {
		fmt.Printf("error opening the file: %s ", banner)
		os.Exit(1)
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	fileLines := bufio.NewScanner(file)

	// Initialize a map to store ASCII art for each character.
	asciiArtMap := make(map[rune][]string)
	var char rune = ' '
	var charArt []string

	// Iterate through each line in the file.
	for fileLines.Scan() {
		if len(fileLines.Text()) != 0 {
			charArt = append(charArt, fileLines.Text())
			// Assign the ASCII art to the character in the map when it reaches 8 lines.
			if len(charArt) == 8 {
				asciiArtMap[char] = charArt
				char++
				charArt = nil
			}
		}
	}
	return asciiArtMap
}

// writing an array of words into a file with the help of AsciiPrintOutArr
func SaveTextToFile(fileName string, textArr []string, asciiArtMap map[rune][]string) {
	// Opening the output file
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file '%s'", fileName)
		return
	}
	defer file.Close()

	// putting the output into a string
	output := AsciiPrintOutArr(textArr, asciiArtMap)

	// Writing the new content into the output file
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Printf("Error writing to file '%s'", fileName)
		return
	}
}
