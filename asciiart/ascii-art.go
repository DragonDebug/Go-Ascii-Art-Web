package asciiart

import (
	art "asciiartweb/functions"
	"fmt"
	"strings"
)

func AsciiArt(args ...string) string {
	// Declare the necessary variables
	var text, banner, option, fileName string

	// Handling cases based on the number of arguments
	if len(args) == 3 {
		option = args[0]
		text = args[1]
		banner = args[2]

		// Verify option
		var err error
		fileName, err = art.VarifyOption(option)
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nX: go run . --output=<fileName.txt> something standard")
			return "" // Return on error
		}

		// Verify banner
		banner, err = art.VarifyBanner(banner)
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nX: go run . --output=<fileName.txt> something standard")
			return "" // Return on error
		}

	} else if len(args) == 2 {
		text = args[0]
		banner = args[1]

		// Verify banner
		var err error
		banner, err = art.VarifyBanner(banner)
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nX: go run . --output=<fileName.txt> something standard")
			return "" // Return on error
		}

	} else if len(args) == 1 {
		text = args[0]
		banner = "standard.txt" // Default banner
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nX: go run . --output=<fileName.txt> something standard")
		return ""
	}

	// Return nothing if the text value is empty
	if text == "" {
		fmt.Println("No input text provided")
		return ""
	}

	// Create the map for the ASCII art
	asciiArtMap := art.AssignArt(banner)

	// Convert Windows line endings to Unix line endings
	text = strings.ReplaceAll(text, "\r\n", "\\n")

	// Check if the input consists of only ASCII characters
	if !art.IsAsciiInput(text, asciiArtMap) {
		fmt.Println("Only ASCII characters are accepted as input")
		return ""
	}

	// Split the input wherever a new line (\n) is detected
	textArr := strings.Split(text, "\\n")

	// Print or save output depending on the number of arguments
	if len(args) == 3 {
		art.SaveTextToFile(fileName, textArr, asciiArtMap)
		fmt.Printf("Your ASCII art has been written into the text box: '%s'\n", fileName)
	}

	// Return the ASCII art as a string
	return art.AsciiPrintOutArr(textArr, asciiArtMap)
}
